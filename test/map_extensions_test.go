package test

import (
	"fmt"
	"github.com/oneliang/util-golang/common"
	"testing"
)

type Bean struct {
	Id   int
	Name string
}

func TestMapExtensions(t *testing.T) {
	inputMap := make(map[int]*Bean)
	inputMap[0] = &Bean{
		Id:   0,
		Name: "0",
	}
	inputMap[1] = &Bean{
		Id:   1,
		Name: "1",
	}
	inputMap[2] = &Bean{
		Id:   2,
		Name: "2",
	}
	otherMap := make(map[int]*Bean)
	otherMap[0] = &Bean{
		Id:   0,
		Name: "0",
	}
	otherMap[1] = &Bean{
		Id:   1,
		Name: "1",
	}
	otherMap[2] = &Bean{
		Id:   2,
		Name: "2",
	}
	list := common.MapDiffersDefault[int, *Bean](inputMap, otherMap)
	fmt.Println(list)
	for _, item := range list {
		fmt.Println(item)
	}
}
