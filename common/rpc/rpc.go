package rpc

import (
	"casino_common/common/Error"
	"casino_common/common/log"
	"net/rpc"
)

//"127.0.0.1:1234"
func Dial(address string, call string, args interface{}, reply interface{}) {
	go func() {
		Error.ErrorRecovery("调用rpc")
		dail(address, call, args, reply)
	}()
}

func dail(address string, call string, args interface{}, reply interface{}) {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.E("链接rpc服务器:%v失败:%v", address, err)
	}
	err = client.Call(call, args, reply)
	if err != nil {
		log.E("链接rpc服务器:%v失败:%v", address, err)
	}
	log.T("链接rpc服务器:%v结果:%v", address, reply)
}
