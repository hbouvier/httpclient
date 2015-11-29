package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BasicAuthentication struct {
	Name   string
	Secret string
}

type Client struct {
	url         string
	credentials *BasicAuthentication
	headers     map[string]string
}

func New(url string, credentials *BasicAuthentication, headers map[string]string) Client {
	if headers == nil {
		headers = map[string]string{}
	}
	Client := Client{url: url, credentials: credentials, headers: headers}
	return Client
}

func (this *Client) Get(url string, response interface{}) error {
	return this.request("GET", url, nil, response)
}

func (this *Client) Post(url string, body interface{}, response interface{}) error {
	return this.request("POST", url, body, response)
}

func (this *Client) Put(url string, body interface{}, response interface{}) error {
	return this.request("PUT", url, body, response)
}

func (this *Client) Delete(url string, body interface{}, response interface{}) error {
	return this.request("DELETE", url, body, response)
}

///////////////////////////////////////////////////////////////////////////////

func (this *Client) request(method string, url string, body interface{}, response interface{}) error {
	jsonBuffer, marshallingError := toJsonByteArray(body)
	if marshallingError != nil {
		return NewMarshallingError(fmt.Sprintf("%s %s >> %v >> serialization %v", method, url, body, marshallingError), marshallingError)
	}
	rawContent, httpResponse, networkError := this.send(method, url, jsonBuffer)
	if networkError != nil {
		return NewNetworkError(fmt.Sprintf("%s %s >> network %v", method, url, networkError), networkError)
	}

	switch {
	case httpResponse.StatusCode >= 200 && httpResponse.StatusCode <= 299:
		if unmarshallingError := parseJsonByteArray(rawContent, response); unmarshallingError != nil {
			return NewUnmarshallingError(fmt.Sprintf("%s %s >> %v >> deserialization %v >> %v", method, url, body, rawContent, unmarshallingError), unmarshallingError)
		}
		return OK
	case httpResponse.StatusCode == 401:
		authenticationResponse := ""
		if unmarshallingError := parseJsonByteArray(rawContent, &authenticationResponse); unmarshallingError != nil {
			return NewAuthenticationError(fmt.Sprintf("%s %s >> Unauthorized (HTTP/401) response deserialization error %v", method, url, unmarshallingError))
		}
		return NewAuthenticationError(fmt.Sprintf("%s %s >> Unauthorized (HTTP/401) >> %s", method, url, authenticationResponse))
	case httpResponse.StatusCode == 404:
		notFoundResponse := ""
		if unmarshallingError := parseJsonByteArray(rawContent, &notFoundResponse); unmarshallingError != nil {
			return NewNotFoundError(fmt.Sprintf("%s %s >> Not found (HTTP/404) response deserialization error %v", method, url, unmarshallingError))
		}
		return NewNotFoundError(fmt.Sprintf("%s %s >> Not found (HTTP/404) >> %s", method, url, notFoundResponse))
	case httpResponse.StatusCode >= 500 && httpResponse.StatusCode <= 599:
		serverFailureResponse := ""
		if unmarshallingError := parseJsonByteArray(rawContent, &serverFailureResponse); unmarshallingError != nil {
			return NewServerFailureError(fmt.Sprintf("%s %s >> Not found (HTTP/%d) response deserialization error %v", method, url, httpResponse.StatusCode, unmarshallingError), httpResponse.StatusCode)
		}
		return NewServerFailureError(fmt.Sprintf("%s %s >> Not found (HTTP/%d) >> %s", method, url, httpResponse.StatusCode, serverFailureResponse), httpResponse.StatusCode)
	}
	unexpectedResponse := ""
	if unmarshallingError := parseJsonByteArray(rawContent, &unexpectedResponse); unmarshallingError != nil {
		return NewNotOkError(fmt.Sprintf("%s %s >> Unexpected HTTP/%d status code response deserialization error %v", method, url, httpResponse.StatusCode, unmarshallingError), httpResponse.StatusCode)
	}
	return NewNotOkError(fmt.Sprintf("%s %s >> Unexpected HTTP/%d status code >> %s", method, url, httpResponse.StatusCode, unexpectedResponse), httpResponse.StatusCode)
}
func toJsonByteArray(object interface{}) ([]byte, error) {
	var jsonBuffer []byte

	switch body := object.(type) {
	case string:
		jsonBuffer = []byte(body)
	case []byte:
		jsonBuffer = body
	default:
		var marshallingError error
		jsonBuffer, marshallingError = json.Marshal(object)
		if marshallingError != nil {
			return nil, marshallingError
		}
	}
	return jsonBuffer, nil
}
func (this *Client) send(method string, url string, body []byte) ([]byte, *http.Response, error) {
	return sendImpl(method, this.url+url, this.headers, this.credentials, body)
}

var sendImpl = func(method string, url string, headers map[string]string, credentials *BasicAuthentication, body []byte) ([]byte, *http.Response, error) {
	httpClient := &http.Client{}
	request, requestError := http.NewRequest(method, url, bytes.NewReader(body))
	if requestError != nil {
		return nil, nil, requestError
	}
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	if credentials != nil {
		request.SetBasicAuth(credentials.Name, credentials.Secret)
	}
	response, responseError := httpClient.Do(request)
	if responseError != nil {
		return nil, nil, responseError
	}

	defer response.Body.Close()
	content, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		return nil, response, readError
	}
	return content, response, nil
}

func parseJsonByteArray(rawContent []byte, result interface{}) error {
	switch body := result.(type) {
	case *string:
		(*body) = string(rawContent)
	case *[]byte:
		(*body) = rawContent
	default:
		if unmarshallingError := json.Unmarshal(rawContent, &result); unmarshallingError != nil {
			return unmarshallingError
		}
	}
	return nil
}
