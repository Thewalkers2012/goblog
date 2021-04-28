package types

import (
	"goblog/pkg/logger"
	"strconv"
)

// Int64ToString 将 int64 型转换为 string 类型
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// StringToInt 将字符串转化为 int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}
