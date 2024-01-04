package common

import (
	"fmt"
	"strings"
)

func GenerateZeroString(number int, length int) string {
	return fmt.Sprintf("%0*d", length, number)
}

func StringIsBlank(content string) bool {
	return len(strings.TrimSpace(content)) == 0
}

func StringIsEmpty(content string) bool {
	return len(content) == 0
}
