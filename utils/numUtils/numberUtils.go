package numUtils

import (
	"fmt"
	"strconv"
)


/**
	int类型转字符串
 */
func Int2String(i int32) (string,error){
	str := fmt.Sprintf("%d", i)
	return str,nil
}

func Int642String(i int64) (string,error){
	str := fmt.Sprintf("%d", i)
	return str,nil
}



/**
	uint类型转字符串u
 */
func Uint2String(i uint32) (string,error){
	str := fmt.Sprintf("%d", i)
	return str,nil
}

/**
 字符转转int类型
 */
func String2Int(s string) (int){
	i,_ := strconv.Atoi(s)
	return i
}
