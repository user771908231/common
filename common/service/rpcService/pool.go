package rpcService

import (
	"google.golang.org/grpc"
	"sync"
	"casino_common/common/log"
	"errors"
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
	if p.HasInit() {
		const err_msg = "已经初始化过了！请勿重复初始化。"
		log.E(err_msg)
		return errors.New(err_msg)
	}
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

//是否已经初始化
func (p *Pool) HasInit() bool {
	if p==nil || p.address == "" || p.initConn == 0 || len(p.pools) == 0 {
		return false
	}
	return true
}

//从连接池中取一个出来
func (p *Pool) Get() *grpc.ClientConn {
	if !p.HasInit() {
		log.E("Rpc Pool未初始化！ %v", p)
		return nil
	}
	p.mu.Lock()
	next_index := p.pool_index%len(p.pools)
	p.pool_index++
	if p.pool_index > 10000 {
		p.pool_index = 0
	}
	p.mu.Unlock()
	conn := p.pools[next_index]
	conn_state := conn.GetState()
	//断线重连
	if conn==nil || ( conn_state != grpc.Connecting && conn_state != grpc.Ready ) {
		log.T("rpc conn:[%v]-[%v] need reconnect.", p.address, conn.GetState())
		conn.Close()
		conn, _ = grpc.Dial(p.address, grpc.WithInsecure())
		p.pools[next_index] = conn
	}
	return conn
}

