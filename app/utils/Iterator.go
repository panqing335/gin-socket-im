package util

type IteratorType interface {
	map[string]string | map[string]any | []map[string]string | []map[string]any | any
}

type Iterator[T IteratorType] struct {
	data  []T
	index int
}

func NewIterator[T IteratorType](data []T) *Iterator[T] {
	return &Iterator[T]{data: data}
}

func (i *Iterator[T]) HasNext() bool {
	if i.data == nil || len(i.data) == 0 {
		return false
	}
	return i.index < len(i.data)
}

func (i *Iterator[T]) Next() (ret T) {
	ret = i.data[i.index]
	i.index = i.index + 1
	return ret
}
