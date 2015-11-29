package httpclient

import (
	"errors"
	"fmt"
	"testing"
)

func TestAuthenticationError(t *testing.T) {
	msg := "not authorized"
	if err := NewAuthenticationError(msg); err != nil {
		if !IsAuthenticationError(err) {
			t.Fatalf("Expected error to be an Autentication (%#v)", err)
		}
		if fmt.Sprintf("%s", err) != msg {
			t.Fatalf("Expected error message to be '%s' but was '%s' >> (%#v)", msg, fmt.Sprintf("%s", err), err)
		}
	} else {
		t.Fatalf("Expected an Autentication error")
	}

	if IsAuthenticationError(errors.New("a generic error")) {
		t.Fatalf("Expected IsAutenticationError to return false on a generic error")
	}
}
