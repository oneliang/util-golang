package test

import (
	"fmt"
	common "github.com/oneliang/util-golang/common"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println(common.GenerateZeroString(11, 20))
	blankString := "    "
	emptyString := ""
	fmt.Printf("is blank:%t\n", common.StringIsBlank(blankString))
	fmt.Printf("is empty:%t\n", common.StringIsEmpty(emptyString))

}
