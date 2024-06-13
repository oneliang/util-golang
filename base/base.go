package base

type Also[T any] interface {
	Also(func(*T)) *T
}

type Let[T any, R any] interface {
	Let(func(*T) *R) *R
}

func AlsoFunc[T any](receiver *T, block func(*T)) *T {
	block(receiver)
	return receiver
}

func LetFunc[T any, R any](receive *T, block func(*T) *R) *R {
	return block(receive)
}
