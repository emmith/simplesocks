package ciper

import (
	"testing"
)

func TestRandPassword(t *testing.T) {
	password := RandPassword()
	t.Log(password)
	bsPassword, err := ParsePassword(password)
	if err != nil {
		t.Error(err)
	}
	t.Log(bsPassword.String() == password)
}
