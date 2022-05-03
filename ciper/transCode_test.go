package ciper

import (
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	password := RandPassword()
	t.Log(password)
	msg := []byte("i am tcp client.ha ha ha ha")
	t.Log(msg)
	pwd, _ := ParsePassword(password)
	cipher := NewCipher(pwd)
	cipher.Encode(msg)
	t.Log(string(msg))
	cipher.Decode(msg)
	t.Log(string(msg))
}
