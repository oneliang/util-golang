package common

import (
	"encoding/binary"
)

var Primitive = &primitive{}

type primitive struct {
}

// IntToByteArray int to byte array
func (this *primitive) IntToByteArray(i int) []byte {
	byteArray := make([]byte, 4)
	binary.BigEndian.PutUint32(byteArray, uint32(i))
	return byteArray
}

// ByteArrayToInt byte array to int
func (this *primitive) ByteArrayToInt(byteArray []byte) int {
	return int(binary.BigEndian.Uint32(byteArray))
}
