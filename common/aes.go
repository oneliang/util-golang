package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func AESEncrypt(key []byte, plainString string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	plainBytes := []byte(plainString)
	// 对于CBC模式，需要使用PKCS#7填充plain text到block size的整数倍
	plainBytes = padding(plainBytes, aes.BlockSize)
	cipherString := make([]byte, aes.BlockSize+len(plainBytes))
	iv := cipherString[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherString[aes.BlockSize:], plainBytes)
	return base64.StdEncoding.EncodeToString(cipherString), nil
}

func AESDecrypt(key []byte, encryptString string) (string, error) {
	cipherString, err := base64.StdEncoding.DecodeString(encryptString)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(cipherString) < aes.BlockSize {
		return "", err
	}
	iv := cipherString[:aes.BlockSize]
	cipherString = cipherString[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherString, cipherString)
	// 删除PKCS#7填充
	plaintext := unPadding(cipherString)
	return string(plaintext), nil
}
