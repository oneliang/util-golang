package common

type Pair[F interface{}, S interface{}] struct {
	First  F
	Second S
}

func NewPair[F interface{}, S interface{}](first F, second S) *Pair[F, S] {
	return &Pair[F, S]{
		First:  first,
		Second: second,
	}
}

type Tuple[F interface{}, S interface{}, T interface{}] struct {
	First  F
	Second S
	Third  T
}

func NewTuple[F interface{}, S interface{}, T interface{}](first F, second S, third T) *Tuple[F, S, T] {
	return &Tuple[F, S, T]{
		First:  first,
		Second: second,
		Third:  third,
	}
}
