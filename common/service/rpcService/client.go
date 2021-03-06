package rpcService

import (
	"casino_common/proto/ddproto"
	"casino_common/proto/ddproto/mjproto"
)

//Greeter
var GreeterPool Pool

func GetGreeter() ddproto.GreeterClient {
	return ddproto.NewGreeterClient(GreeterPool.Get())
}

//转转红中麻将
var ZzHzPool Pool

func GetZzHz() mjproto.ZzHzClient {
	return mjproto.NewZzHzClient(ZzHzPool.Get())
}

//跑得快
var PdkPool Pool

func GetPdk() ddproto.PdkRpcClient {
	return ddproto.NewPdkRpcClient(PdkPool.Get())
}

//大厅
var HallPool Pool

func GetHall() ddproto.HallRpcClient {
	return ddproto.NewHallRpcClient(HallPool.Get())
}

//松江河麻将
var SjhMjPool Pool

func GetSjhMj() mjproto.SjhMjRpcClient {
	return mjproto.NewSjhMjRpcClient(SjhMjPool.Get())
}

//牛牛
var NiuniuPool Pool

func GetNiuniu() ddproto.NiuniuRpcClient {
	return ddproto.NewNiuniuRpcClient(NiuniuPool.Get())
}

//转转红中麻将
var SiChuanPool Pool

func GetSiChuan() mjproto.SiChuanRpcClient {
	return mjproto.NewSiChuanRpcClient(SiChuanPool.Get())
}
