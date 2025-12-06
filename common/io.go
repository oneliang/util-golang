package common

import (
	"io"
)

func ReadReadCloser(readCloser io.ReadCloser) ([]byte, int, error) {
	defer func() { _ = readCloser.Close() }()
	var result []byte
	buffer := make([]byte, 1024) // buffer
	length := 0
	for {
		n, err := readCloser.Read(buffer)
		if n > 0 {
			result = append(result, buffer[:n]...)
			length += n
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return result, 0, err
		}
	}
	return result, length, nil
}
