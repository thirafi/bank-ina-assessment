package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomString() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err.Error())
	}
	state := base64.StdEncoding.EncodeToString(b)
	return state
}
