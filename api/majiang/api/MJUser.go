package api

import "github.com/golang/protobuf/proto"

type PDKUser interface {
	WriteMsg(p proto.Message) error //发送信息
}
