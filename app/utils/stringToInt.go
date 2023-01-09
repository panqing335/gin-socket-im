package util

import (
	"strconv"
)

func StringToInt(str string, size int) any {
	i, err := strconv.ParseInt(str, 10, size)
	if err != nil {
		panic("string转换int失败")
	}
	return i
}
