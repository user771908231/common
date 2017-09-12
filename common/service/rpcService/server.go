package rpcService

import (
	"casino_common/proto/ddproto"
	"google.golang.org/grpc"
	"net"
)

//监听rpc服务
func LisenAndServeGreeter(addr string, greeter ddproto.GreeterServer) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	ddproto.RegisterGreeterServer(s, greeter)
	go s.Serve(lis)
	return s, nil
}
