package common

import (
	"bytes"
	"runtime"
	"strconv"
)

func GetGoroutineId() int64 {
	buffer := make([]byte, 64)
	buffer = buffer[:runtime.Stack(buffer, false)]
	buffer = bytes.TrimPrefix(buffer, []byte("goroutine "))
	buffer = buffer[:bytes.IndexByte(buffer, ' ')]
	goroutineId, err := strconv.ParseInt(string(buffer), 10, 64)
	if err != nil {
		return 0
	}
	return goroutineId
}
