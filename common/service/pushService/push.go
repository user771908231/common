package pushService

import (
	"net"
	"casino_common/proto/ddproto"
	"gopkg.in/fatih/pool.v2"
	"errors"
)

func init() {
	//初始化连接池
	//PoolInit()
}

//大厅地址
var HallTcpAddr string

//初始化连接池
func PoolInit(hall_addr string) error {
	//保存配置到变量里
	HallTcpAddr = hall_addr
	//此处注释以便单机运行
	factory := func() (net.Conn, error) {
		return net.Dial("tcp", HallTcpAddr)
	}
	var err error = nil
	PoolStack, err = pool.NewChannelPool(5, 30, factory)
	if err != nil {
		return errors.New("pool init err.")
	}
	return nil
}

var PoolStack pool.Pool

//向大厅推送消息
func Push(data []byte) error {
	//如果遇到错误，则尝试重新初始化连接池，保证与大厅断线后可以第一时间重连。
	if PoolStack == nil {
		if PoolInit(HallTcpAddr) == nil {
			go Push(data)
			return nil
		}
		return errors.New("push err: PoolStack nil.")
	}

	conn,err := PoolStack.Get()
	defer conn.Close()

	if err != nil {
		if PoolInit(HallTcpAddr) == nil {
			go Push(data)
			return nil
		}
		return errors.New("push err: Get err")
	}

	_,err = conn.Write(data)
	if err != nil {
		if PoolInit(HallTcpAddr) == nil {
			go Push(data)
			return nil
		}
		return errors.New("push err: Write err.")
	}

	return nil
}

//向大厅服务器推送消息
func PushUserData(userId uint32) error {
	data := AssembleData(ddproto.HallEnumProtoId_HALL_PID_PUSH_REQ, &ddproto.Push{
		Id:&userId,
	})
	return Push(data)
}
