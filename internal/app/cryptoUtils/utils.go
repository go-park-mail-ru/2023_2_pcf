package cryptoUtils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenToken(length int) (str string, err error) {
	b := make([]byte, length)
	_, err = rand.Read(b)
	if err != nil {
		return str, err
	}
	str = base64.URLEncoding.EncodeToString(b)[:length]
	return str, err
}
