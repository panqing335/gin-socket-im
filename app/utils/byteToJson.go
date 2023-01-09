package util

import "encoding/json"

type BTJType interface {
	any | []any | bool | []string | map[string]any | []map[string]any
}

func ByteToJson[T BTJType](jsonData *[]byte) *T {
	var v T
	json.Unmarshal(*jsonData, &v)
	data := v
	return &data
}
