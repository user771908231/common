package redisUtils

import (
	"github.com/golang/protobuf/proto"
	"casino_common/utils/redis"
	"casino_common/common/log"
	"strings"
	"casino_common/utils/numUtils"
)

//在需要的地方,需要自己关闭连接
func GetConn() data.Data {
	conn := data.Data{}
	conn.Open()
	return conn
}

/**
	获取redis中的一个对象,如果redis中没有值,那么返回nil
	如果有值,那么返回一个proto.message,需要取值的时候,需要强制转换
	eg.
		proto := GetObj(key,&bbproto.type{})
		if proto == nil{
			//没有找到值
		}else{
			result := proto.(*bbproto.type)
		}
 */

func GetObj(key string, p proto.Message) proto.Message {
	conn := GetConn()
	defer conn.Close()
	return conn.GetObjv2(key, p)
}

//得到int64的值
func GetInt64(key string) int64 {
	conn := GetConn()
	defer conn.Close()
	result, err := conn.GetInt64(key)
	if err != nil {
		return -1
	} else {
		return int64(result)
	}
}

func SetInt64(key string, value int64) {
	conn := GetConn()
	defer conn.Close()
	conn.SetInt(key, value)
}

/**
	保存一个对象到redis
 */
func SetObj(key string, p proto.Message) error {
	conn := GetConn()
	defer conn.Close()
	return conn.SetObj(key, p)
}

func Del(key string) {
	conn := GetConn()
	defer conn.Close()
	conn.Del(key)
}

func ZADD(key string, member string, score int64) {
	conn := GetConn()
	defer conn.Close()
	conn.ZAdd(key, member, score)
}

func ZREVRANK(key string, member string) int64 {
	conn := GetConn()
	defer conn.Close()
	values, err := conn.ZREVRANK(key, member)
	if values == nil || err != nil {
		return -1
	}
	return values.(int64)
}

//	LREM key count value 移除列表元素
func LREM(key string, value string) error {
	conn := GetConn()
	defer conn.Close()
	err := conn.LREM(key, value)
	if err != nil {
		log.E("redis lrem的时候出错 err[%v]", err)
	}
	return err
}

//对指定的key 加上 i的值
func INCRBY(key string, i int64) int64 {
	conn := GetConn()
	defer conn.Close()
	value, err := conn.INCRBY(key, i)
	if err != nil {
		log.E("redis INCRBY的时候出错 err[%v]", err)
		return 0
	} else {
		return value.(int64)
	}
}

//对指定的key 减去i 的值
func DECRBY(key string, i int64) int64 {
	conn := GetConn()
	defer conn.Close()
	value, err := conn.DECRBY(key, i)
	if err != nil {
		log.E("redis DECRBY的时候出错 err[%v]", err)
		return 0
	} else {
		return value.(int64)
	}
}

//浅醉加上userID 的redis-key
func K(pre string, userId uint32) string {
	userIdStr, _ := numUtils.Uint2String(userId)
	result := strings.Join([]string{pre, userIdStr}, "_")
	return result
}
