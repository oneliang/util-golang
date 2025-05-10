package base

type Pair[F any, S any] struct {
	First  F
	Second S
}

func NewPair[F any, S any](first F, second S) *Pair[F, S] {
	return &Pair[F, S]{
		First:  first,
		Second: second,
	}
}

type Tuple[F any, S any, T any] struct {
	First  F
	Second S
	Third  T
}

func NewTuple[F any, S any, T any](first F, second S, third T) *Tuple[F, S, T] {
	return &Tuple[F, S, T]{
		First:  first,
		Second: second,
		Third:  third,
	}
}

type KeyValue[K any, V any] struct {
	Key   K
	Value V
}

func NewKeyValue[K any, V any](key K, value V) *KeyValue[K, V] {
	return &KeyValue[K, V]{
		Key:   key,
		Value: value,
	}
}
