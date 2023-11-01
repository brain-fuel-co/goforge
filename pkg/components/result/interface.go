package result

type Result[T any] struct {
	right T
	left  error
}

type MonadicFunction[T any, U any] func(T) *Result[U]

func (result *Result[T]) Right() T {
	return result.right
}

func (result *Result[T]) Left() error {
	return result.left
}

func Return[T any](a T) *Result[T] {
	return &Result[T]{right: a}
}

func ErrorResult[T any](e error) *Result[T] {
	return &Result[T]{left: e}
}

func Bind[T any, U any](a *Result[T], f MonadicFunction[T, U]) *Result[U] {
	if a.Left() != nil {
		return &Result[U]{left: a.Left()}
	} else {
		return f(a.Right())
	}
}
