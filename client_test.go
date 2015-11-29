package httpclient

import (
	"net/http"
	"testing"
)

type StringAndInt struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func TestValideJsonResponse(t *testing.T) {
	sendImpl = func(method string, url string, headers map[string]string, credentials *BasicAuthentication, body []byte) ([]byte, *http.Response, error) {
		return []byte("{\"message\":\"Hello World\",\"code\":42}"), &http.Response{StatusCode: 200}, nil
	}
	response := StringAndInt{}
	client := New("http://localhost", nil, nil)
	if err := client.Get("/index.json", &response); err != nil {
		t.Fatalf("Expected to recieved a valide json object (%#v)", err)
	}
}

func TestToJsonByteArray(t *testing.T) {
	message := StringAndInt{"Hello World", 42}
	jsonByteArray, err := toJsonByteArray(message)
	if err != nil {
		t.Fatalf("Expected to serialize object to json byte array (%#v)", err)
	}

	json := string(jsonByteArray)
	if json != "{\"message\":\"Hello World\",\"code\":42}" {
		t.Fatalf("Expected json byte array to be '\"{\"message\":\"Hello World\",\"code\":42}\"' but was (%s)", json)
	}
}

func TestParseJsonByteArray(t *testing.T) {
	response := StringAndInt{}
	err := parseJsonByteArray([]byte("{\"message\":\"Hello World\",\"code\":42}"), &response)
	if err != nil {
		t.Fatalf("Expected to parse json byte array (%#v)", err)
	}
	if response.Message != "Hello World" {
		t.Fatalf("Expected message to be 'Hello World' but was (%s)", response.Message)
	}
	if response.Code != 42 {
		t.Fatalf("Expected code to be the answer to the god, the universe and everything (e.g. 42) but was %d", response.Code)
	}
}
