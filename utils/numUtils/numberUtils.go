package numUtils

import (
	"fmt"
	"strconv"
)

/**
int类型转字符串
*/
func Int2String(i int32) (string, error) {
	str := fmt.Sprintf("%d", i)
	return str, nil
}

func Int2String2(i int32) string {
	str := fmt.Sprintf("%d", i)
	return str
}

func Int642String(i int64) (string, error) {
	str := fmt.Sprintf("%d", i)
	return str, nil
}

/**
uint类型转字符串u
*/
func Uint2String(i uint32) (string, error) {
	str := fmt.Sprintf("%d", i)
	return str, nil
}

/**
字符转转int类型
*/
func String2Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func String2Uint32(s string) uint32 {
	return uint32(String2Int(s))
}

/**
字符转转float64类型
*/
func String2Float64(s string) float64 {
	i, _ := strconv.ParseFloat(s, len(s))
	return i
}

/**
float64,防止丢精度.
*/
func Float64Format(num float64) float64 {
	res, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", num), 64)
	return res
}
