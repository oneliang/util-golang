package test

import (
	"github.com/oneliang/util-golang/common"
	"log"
	"testing"
)

func TestAES(t *testing.T) {
	key := []byte("1234567890123456") // 16字节长度
	plainText := "Hello, World!"

	// 加密
	cipherText, err := common.AESEncrypt(key, plainText)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encrypted:", cipherText)

	// 解密
	decrypted, err := common.AESDecrypt(key, cipherText)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Decrypted:", decrypted)
}
