package ciper

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
	"time"
)

const passwordLength = 256

type password [passwordLength]byte

// 以base64编码的形式转换成字符串
func (password *password) String() string {
	return base64.StdEncoding.EncodeToString(password[:])
}

//  更新随机种子，防止生成一样的密码
func init() {
	rand.Seed(time.Now().Unix())
}

//  生成一个256个字节的随机密码
func RandPassword() string {
	// 随机生成一个由0-255组成的byte数组
	intArr := rand.Perm(passwordLength)
	password := &password{}

	for i, v := range intArr {
		if i == v {
			// 确保每一个位置的下标和对应的值不相等
			return RandPassword()
		}
		password[i] = byte(v)
	}
	return password.String()
}

func ParsePassword(passwordString string) (*password, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(passwordString))
	if err != nil || len(bs) != passwordLength {
		return nil, errors.New("password is corrected")
	}

	password := password{}
	copy(password[:], bs)
	bs = nil
	return &password, nil
}
