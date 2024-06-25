package common

import (
	"fmt"
	"log"
	"testing"
)

type Data struct {
	A int
	B string
	C string
}

func TestReflect(t *testing.T) {
	data := &Data{
		A: 1,
		B: "b",
		C: "c",
	}
	err := CopyDataFromMap(data, map[string]any{
		"A": 2,
		"B": "b_replace",
		"C": "c_replace",
	}, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%+v", data))
}
