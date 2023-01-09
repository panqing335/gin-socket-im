package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func StructToMap(content any) map[string]any {
	var name map[string]any
	marshalContent, _ := json.Marshal(content)

	d := json.NewDecoder(bytes.NewReader(marshalContent))
	d.UseNumber()
	if err := d.Decode(&name); err != nil {
		fmt.Println("StructToMap:", err.Error())
	} else {
		for k, v := range name {
			name[k] = v
		}
	}

	return name
}
