package auth

import (
	"fmt"
	"testing"
)

func TestSessionStorage_AddAndGetSession(t *testing.T) {
	s := Session{UserId: -1}
	sH := SessionStorage{Sessions: make(map[string]int)}
	sH.AddSession(s)
	userIdGot, err := sH.GetUserId(s.Token)
	if err != nil {
		t.Errorf("Error while getting a userId: %v\n", err)
	}
	if s.UserId != userIdGot {
		t.Errorf("User id expected: %v, user id got: %v\n", s.UserId, userIdGot)
	}
}

func TestSessionStorage_Contains(t *testing.T) {
	s := Session{UserId: -1}
	s2 := Session{UserId: -2}
	s.SetToken()
	s2.SetToken()

	sH := SessionStorage{Sessions: make(map[string]int)}
	sH.AddSession(s)
	if !sH.Contains(s.Token) {
		fmt.Errorf("Expected result: contains, real: no")
	}
	if sH.Contains(s2.Token) {
		fmt.Errorf("Expected result: doesnt contain, real: yes")
	}
}
