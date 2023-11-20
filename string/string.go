package string

import (
	"fmt"
	"strings"
)

func GenerateZeroString(number int, length int) string {
	return fmt.Sprintf("%0*d", length, number)
}

func IsBlank(content string) bool {
	return len(strings.TrimSpace(content)) == 0
}

func IsEmpty(content string) bool {
	return len(content) == 0
}
