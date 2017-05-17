package rpc

//apc 需要满足的功能
type RpcApi interface {
}

//默认的空实现
type RpcCore struct {
}

//默认的空实现
func NewRpcCore() *RpcCore {
	return &RpcCore{}
}
