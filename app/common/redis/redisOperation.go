package redis

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

type Operation struct {
	ctx context.Context
}

func NewOperation() *Operation {
	return &Operation{ctx: context.Background()}
}

func (o *Operation) Set(key string, value any, attrs ...*OperationAttr) *Result[ResultType] {
	expr := OperationAttrs(attrs).Find(ATTR_EXPIRE).UnwrapOr(time.Second * 0)
	nx := OperationAttrs(attrs).Find(ATTR_NX).UnwrapOr(nil)
	if nx != nil {
		return NewResult[ResultType](myRedis.SetNX(o.ctx, key, value, expr.(time.Duration)).Result())
	}
	xx := OperationAttrs(attrs).Find(ATTR_XX).UnwrapOr(nil)
	if xx != nil {
		return NewResult[ResultType](myRedis.SetXX(o.ctx, key, value, expr.(time.Duration)).Result())
	}
	return NewResult[ResultType](myRedis.Set(o.ctx, key, value, expr.(time.Duration)).Result())
}

func (o *Operation) MSet(values ...any) *Result[ResultType] {
	return NewResult[ResultType](myRedis.MSet(o.ctx, values...).Result())
}

func (o *Operation) HSet(key string, values ...any) *Result[ResultType] {
	return NewResult[ResultType](myRedis.HSet(o.ctx, key, values...).Result())
}

func (o *Operation) HMSet(key string, values ...any) *Result[ResultType] {
	return NewResult[ResultType](myRedis.HMSet(o.ctx, key, values...).Result())
}

func (o *Operation) Get(key string) *Result[string] {
	return NewResult(myRedis.Get(o.ctx, key).Result())
}

func (o *Operation) MGet(keys ...string) *Result[[]any] {
	return NewResult(myRedis.MGet(o.ctx, keys...).Result())
}

func (o *Operation) HGet(key string, field string) *Result[string] {
	return NewResult(myRedis.HGet(o.ctx, key, field).Result())
}

func (o *Operation) HGetAll(key string) *Result[ResultType] {
	result, err := myRedis.HGetAll(o.ctx, key).Result()
	if len(result) == 0 {
		err = errors.New("empty error")
	}
	return NewResult[ResultType](result, err)
}

func (o *Operation) Expire(key string, ttl time.Duration) *Result[ResultType] {
	return NewResult[ResultType](myRedis.Expire(o.ctx, key, ttl).Result())
}

func (o *Operation) Del(key string) *Result[ResultType] {
	return NewResult[ResultType](myRedis.Del(o.ctx, key).Result())
}
