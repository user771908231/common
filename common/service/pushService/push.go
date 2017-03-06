package pushService

import (
	"net"
	"casino_common/proto/ddproto"
	"gopkg.in/fatih/pool.v2"
	"casino_hall/conf"
)

func init() {
	factory := func() (net.Conn, error) {
		return net.Dial("tcp", conf.Server.TCPAddr)
	}
	var err error = nil
	PoolStack, err = pool.NewChannelPool(5, 30, factory)
	if err != nil {
		panic("pool init fail.")
	}
}

var PoolStack pool.Pool

//向大厅服务器推送消息
func PushUserData(userId uint32) error {
	conn,err := PoolStack.Get()
	if err != nil {
		return err
	}
	data := AssembleData(ddproto.HallEnumProtoId_HALL_PID_PUSH_REQ, &ddproto.Push{
		Id:&userId,
	})
	_,err = conn.Write(data)
	defer conn.Close()
	return err
}
