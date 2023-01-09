package redis

type Empty struct{}

func NewEmpty() *Empty {
	return &Empty{}
}

type ResultType interface {
	any | []interface{} | Empty | bool | []string | map[string]any | func() string | []map[string]any | map[string]string
}

type Result[T ResultType] struct {
	Result T
	Err    error
}

func NewResult[T ResultType](result T, err error) *Result[T] {
	return &Result[T]{Result: result, Err: err}
}

func (r *Result[T]) Unwrap() T {
	if r.Err != nil {
		panic(r.Err)
	}

	return r.Result
}

func (r *Result[T]) UnwrapOr(str T) T {
	if r.Err != nil {
		return str
	}

	return r.Result
}

func (r *Result[T]) UnwrapOrElse(f func() T) T {
	if r.Err != nil {
		return f()
	}

	return r.Result
}
