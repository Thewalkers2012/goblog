package types

import "strconv"

// Int64ToString 将 int64 型转换为 string 类型
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}
