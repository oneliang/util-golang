package test

import (
	"fmt"
	"github.com/oneliang/util-golang/string"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println(string.GenerateZeroString(11, 20))
}
