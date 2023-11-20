package test

import (
	"fmt"
	stringUtil "github.com/oneliang/util-golang/string"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println(stringUtil.GenerateZeroString(11, 20))
	blankString := "    "
	emptyString := ""
	fmt.Printf("is blank:%t\n", stringUtil.IsBlank(blankString))
	fmt.Printf("is empty:%t\n", stringUtil.IsEmpty(emptyString))

}
