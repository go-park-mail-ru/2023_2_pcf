package auth

import "AdHub/pkg/cryptoUtils"

const (
	tokenLen = 32
)

type Session struct {
	Token  string `json:"token"`
	UserId int    `json:"-"`
}

func (s *Session) SetToken() (err error) {
	s.Token, err = cryptoUtils.GenToken(tokenLen)
	return err
}