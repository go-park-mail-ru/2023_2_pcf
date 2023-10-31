package cryptoUtils

import (
	"testing"
)

func TestGenToken(t *testing.T) {
	length := 32
	token, err := GenToken(length)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

	if len(token) != length {
		t.Errorf("Expected token of length %d, but got length %d", length, len(token))
	}
}
