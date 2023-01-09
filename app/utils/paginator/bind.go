package paginator

import "temp/app/model"

type ModelStructure interface {
	model.User | []map[string]any | map[string]any
}

type PaginatorCollection[T ModelStructure] struct {
	CurrentPage int   `json:"currentPage"`
	Items       *[]T  `json:"items"`
	PageSize    int   `json:"pageSize"`
	Total       int64 `json:"total"`
}

func NewPaginatorCollection[T ModelStructure](currentPage int, items *[]T, pageSize int, total int64) *PaginatorCollection[T] {
	return &PaginatorCollection[T]{CurrentPage: currentPage, Items: items, PageSize: pageSize, Total: total}
}
