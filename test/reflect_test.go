package test

import (
	"fmt"
	"github.com/oneliang/util-golang/common"
	"log"
	"testing"
)

type ReflectBean struct {
	A int
	B string
	C string
}

func TestReflect(t *testing.T) {
	reflectBean := &ReflectBean{
		A: 1,
		B: "b",
		C: "c",
	}
	err := common.CopyDataFromMap(reflectBean, map[string]any{
		"A": 2,
		"B": "b_replace",
		"C": "c_replace",
	}, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%+v", reflectBean))
}
