package rpcService

import "casino_common/proto/ddproto"

//Greeter
var GreeterPool Pool
func GetGreeter() ddproto.GreeterClient {
	return ddproto.NewGreeterClient(GreeterPool.Get())
}
