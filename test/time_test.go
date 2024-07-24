package test

import (
	"fmt"
	"github.com/oneliang/util-golang/common"
	"testing"
)

func TestTime(t *testing.T) {

	fmt.Println(common.GetCurrentMonthDatesOffset(-5, 1))
}
