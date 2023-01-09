package redis

import (
	"fmt"
	"time"
)

const (
	ATTR_EXPIRE = "expr"
	ATTR_NX     = "nx"
	ATTR_XX     = "xx"
)

type OperationAttr struct {
	Name  string
	Value any
}

type OperationAttrs []*OperationAttr

func (o OperationAttrs) Find(name string) *Result[any] {
	for _, attr := range o {
		if attr.Name == name {
			return NewResult[any](attr.Value, nil)
		}
	}
	return NewResult[any](nil, fmt.Errorf("OperationAttrs found error: %s", name))
}

func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{Name: ATTR_EXPIRE, Value: t}
}

func WithNx() *OperationAttr {
	return &OperationAttr{Name: ATTR_NX, Value: Empty{}}
}

func WithXX() *OperationAttr {
	return &OperationAttr{
		Name:  ATTR_XX,
		Value: Empty{},
	}
}
