package httpclient

import (
	"testing"
)

type StringAndInt struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
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
