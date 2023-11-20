package base

import "fmt"

const (
	Name = "base"
)

func PrintName() {
	fmt.Println("base/base.go")
}

type Also[T interface{}] interface {
	Also(func(*T)) *T
}

type Let[T interface{}, R interface{}] interface {
	Let(func(*T) *R) *R
}

func AlsoFunc[T interface{}](receiver *T, block func(*T)) *T {
	block(receiver)
	return receiver
}

func LetFunc[T interface{}, R interface{}](receive *T, block func(*T) *R) *R {
	return block(receive)
}
