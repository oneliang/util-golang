package common

import "bytes"

// padding 使用PKCS#7标准填充数据
func padding(buf []byte, blockSize int) []byte {
	paddingData := blockSize - (len(buf) % blockSize)
	paddingText := bytes.Repeat([]byte{byte(paddingData)}, paddingData)
	return append(buf, paddingText...)
}

// unPadding 删除PKCS#7填充的字节
func unPadding(buf []byte) []byte {
	length := len(buf)
	if length == 0 {
		return buf
	}
	unPaddingData := int(buf[length-1])
	return buf[:length-unPaddingData]
}
