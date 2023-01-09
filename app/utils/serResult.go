package util

import (
	"fmt"
	"gorm.io/gorm/logger"
	"temp/app/constants/errorCode"
)

type ResultType interface {
	map[string]any | []map[string]any | string | int | any | map[string]string | []map[string]string | []any
}

type Result[T ResultType] struct {
	data T
	Err  error
}

func NewResult[T ResultType](data T, err error) *Result[T] {
	return &Result[T]{Err: err, data: data}
}

func (r *Result[T]) Unwrap() T {
	if r.Err != nil {
		Logger().Error(r.Err.Error())
		Fail(errorCode.BAD_REQUEST, r.Err.Error())
	}
	return r.data
}

func (r *Result[T]) UnwrapOr(code int, str string) T {
	if r.Err != nil {
		Logger().Error(r.Err.Error())
		Fail(code, str)
	}
	return r.data
}

func (r *Result[T]) UnwrapOrNil() (t T) {
	if r.Err != nil {
		fmt.Printf("error: %v\n", r.Err)
		if r.Err == logger.ErrRecordNotFound {
			return
		} else {
			Logger().Error(r.Err.Error())
			Fail(500, "server error")
		}
	}
	return r.data
}
