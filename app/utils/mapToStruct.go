package util

import "encoding/json"

type MTJType interface {
	any | []any | bool | []string | map[string]any | []map[string]any
}

func MapToStruct[T BTJType](req any) *T {
	jsonData, _ := json.Marshal(&req)

	var v T
	json.Unmarshal(jsonData, &v)
	data := v
	return &data
}
