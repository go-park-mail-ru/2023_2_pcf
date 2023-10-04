package auth

import "testing"

func TestSession_SetToken(t *testing.T) {
	s := Session{}
	s.SetToken()
	if len(s.Token) != tokenLen {
		t.Errorf("Expected token len: %v, real len: %v\n", tokenLen, len(s.Token))
	}
}
