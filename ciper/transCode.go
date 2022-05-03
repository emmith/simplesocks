package ciper

type Cipher struct {
	encodepassword *password
	decodepassword *password
}

// 加密数据
func (cipher *Cipher) Encode(data []byte) {
	for i, v := range data {
		data[i] = cipher.encodepassword[v]
	}
}

// 解密数据
func (cipher *Cipher) Decode(data []byte) {
	for i, v := range data {
		data[i] = cipher.decodepassword[v]
	}
}

func NewCipher(pwd *password) *Cipher {
	decodepassword := &password{}
	for i, v := range pwd {
		decodepassword[v] = byte(i)
	}

	return &Cipher{
		encodepassword: pwd,
		decodepassword: decodepassword,
	}
}
