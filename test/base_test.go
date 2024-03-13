package test

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T) {
	fmt.Print("test")
	a := &A{}
	a.Also(func(it *A) {
		it.Name = ""
	})
	a.Let(func(it *A) *A {
		it.Name = ""
		return it
	})
}

type A struct {
	Name string
}

func (a *A) Also(block func(it *A)) *A {
	block(a)
	return a
}

func (a *A) Let(block func(it *A) *A) *A {
	return block(a)
}
