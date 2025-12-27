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

func ExecuteFunctionWithRecover(executeFunction func(params ...any) error, recoverCallback func(recover any), params ...any) error {
	defer func() {
		if r := recover(); r != nil {
			recoverCallback(r)
			return
		}
	}()
	return executeFunction(params)
}
