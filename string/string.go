package string

import "fmt"

func GenerateZeroString(number int, length int) string {
	return fmt.Sprintf("%0*d", length, number)
}
