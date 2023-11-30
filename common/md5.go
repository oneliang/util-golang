package common

import (
	"crypto"
	"encoding/hex"
	"io"
	"os"
)

// FileMD5String file md5 string
func FileMD5String(file *os.File) string {
	return hex.EncodeToString(FileMD5(file))
}

// FileMD5 file md5 []byte
func FileMD5(file *os.File) []byte {
	md5Instance := crypto.MD5.New()

	if _, err := io.Copy(md5Instance, file); err != nil {
		return []byte{}
	}

	return md5Instance.Sum(nil)
}

// StringMD5String string md5 string
func StringMD5String(str string) string {
	return hex.EncodeToString(StringMD5(str))
}

// StringMD5 string md5
func StringMD5(str string) []byte {
	md5Instance := crypto.MD5.New()
	md5Instance.Write([]byte(str))
	return md5Instance.Sum(nil)
}
