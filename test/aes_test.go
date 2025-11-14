package test

import (
	"encoding/base64"
	"fmt"
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

func TestAesGcm(t *testing.T) {
	// 原始数据
	plaintext := []byte("Hello, World! This is a secret message.")

	// 生成密钥（16字节=128位，24字节=192位，32字节=256位）
	key := []byte("abcdef1234567890abcdef1234567890") // 32字节 = 256位

	fmt.Printf("原始数据: %s\n", plaintext)
	fmt.Printf("密钥: %s\n", base64.StdEncoding.EncodeToString(key))

	// 使用GCM模式加密
	fmt.Println("\n=== GCM模式加密 ===")
	encryptedGCM, err := common.AesEncryptGCM(plaintext, key)
	if err != nil {
		log.Fatal("GCM加密失败:", err)
	}
	fmt.Printf("加密结果(Base64): %s\n", base64.StdEncoding.EncodeToString(encryptedGCM))

	// GCM解密
	decryptedGCM, err := common.AesDecryptGCM(encryptedGCM, key)
	if err != nil {
		log.Fatal("GCM解密失败:", err)
	}
	fmt.Printf("解密结果: %s\n", decryptedGCM)
	// 验证加解密是否正确
	fmt.Printf("\nGCM模式加解密验证: %v\n", string(decryptedGCM) == string(plaintext))
}
