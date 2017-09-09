package rpcService

import (
	"google.golang.org/grpc"
	"sync"
)

type Pool struct {
	address    string
	initConn   int
	pools      []*grpc.ClientConn
	pool_index int
	mu         sync.Mutex
}

//初始化连接池
func (p *Pool) Init(addr string, init_conn int) error {
	p.address = addr
	p.initConn = init_conn
	for i:=0; i< p.initConn; i++ {
		conn, err := grpc.Dial(p.address, grpc.WithInsecure())
		if err != nil {
			return err
		}
		p.pools = append(p.pools, conn)
	}
	return nil
}

//从连接池中取一个出来
func (p *Pool) Get() *grpc.ClientConn {
	p.mu.Lock()
	next_index := p.pool_index%len(p.pools)
	p.pool_index++
	p.mu.Unlock()
	conn := p.pools[next_index]
	conn_state := conn.GetState()
	//断线重连
	if conn==nil || ( conn_state != grpc.Connecting && conn_state != grpc.Ready ) {
		conn, _ = grpc.Dial(p.address, grpc.WithInsecure())
		p.pools[next_index] = conn
	}
	return conn
}

