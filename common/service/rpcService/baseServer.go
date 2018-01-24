package rpcService

import (
	"casino_common/common/log"
	"casino_common/proto/ddproto"
	"golang.org/x/net/context"
	"io"
)

type BaseServer struct {
}

//普通远程调用
func (s BaseServer) SayHello(ctx context.Context, in *ddproto.HelloRequest) (*ddproto.HelloReply, error) {

	return &ddproto.HelloReply{Message: "hello"}, nil
}

//双向流通性-服务端
func (s BaseServer) SayHelloStream(conn ddproto.Greeter_SayHelloStreamServer) error {
	for {
		req, err := conn.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		log.T("recv: ", req.Name)
		res_msg := "收到:" + req.Name
		conn.Send(&ddproto.HelloReply{Message: res_msg})
		log.T("send: ", res_msg)
	}

	return nil
}
