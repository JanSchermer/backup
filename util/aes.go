package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type AES struct {
	key []byte
}

func GetAES(key string) AES {
	return AES{[]byte(key[:32])}
}

func (a *AES) Encrypt(plainText []byte) []byte {
	ciph, err := aes.NewCipher(a.key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	return cipherText
}

func (a *AES) Decrypt(cipherText []byte) []byte {
	ciph, err := aes.NewCipher(a.key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
		panic(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err)
	}
	return plainText
}
