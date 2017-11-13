// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc_niuniu.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of niuniu_client_desk from niuniu_play.proto

// Ignoring public import of niuniu_client_user from niuniu_play.proto

// Ignoring public import of niu_create_desk_req from niuniu_play.proto

// Ignoring public import of niu_enter_desk_req from niuniu_play.proto

// Ignoring public import of niu_enter_desk_ack from niuniu_play.proto

// Ignoring public import of niu_enter_desk_bc from niuniu_play.proto

// Ignoring public import of niu_switch_ready_req from niuniu_play.proto

// Ignoring public import of niu_switch_ready_ack from niuniu_play.proto

// Ignoring public import of niu_switch_ready_bc from niuniu_play.proto

// Ignoring public import of niu_start_game_ot from niuniu_play.proto

// Ignoring public import of niu_qiangzhuang_ot from niuniu_play.proto

// Ignoring public import of niu_qiangzhuang_req from niuniu_play.proto

// Ignoring public import of niu_qiangzhuang_ack from niuniu_play.proto

// Ignoring public import of niu_qiangzhuang_res_item from niuniu_play.proto

// Ignoring public import of niu_qiangzhuang_res_bc from niuniu_play.proto

// Ignoring public import of niu_jiabei_ot from niuniu_play.proto

// Ignoring public import of niu_jiabei_req from niuniu_play.proto

// Ignoring public import of niu_jiabei_ack from niuniu_play.proto

// Ignoring public import of niu_jiabei_bc from niuniu_play.proto

// Ignoring public import of niu_xuanpai_ot from niuniu_play.proto

// Ignoring public import of niu_xuanpai_req from niuniu_play.proto

// Ignoring public import of niu_xuanpai_ack from niuniu_play.proto

// Ignoring public import of niu_xuanpai_bc from niuniu_play.proto

// Ignoring public import of niu_bipai_result_item from niuniu_play.proto

// Ignoring public import of niu_bipai_result_bc from niuniu_play.proto

// Ignoring public import of niu_game_end from niuniu_play.proto

// Ignoring public import of niu_desk_dissolve_done_bc from niuniu_play.proto

// Ignoring public import of niu_owner_dissolve_req from niuniu_play.proto

// Ignoring public import of niu_owner_dissolve_ack from niuniu_play.proto

// Ignoring public import of niu_offline_bc from niuniu_play.proto

// Ignoring public import of niu_coin_room_list_req from niuniu_play.proto

// Ignoring public import of niu_coin_room_list_ack from niuniu_play.proto

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NiuniuRpc service

type NiuniuRpcClient interface {
	// 创建房间
	CreateRoom(ctx context.Context, in *NiuCreateDeskReq, opts ...grpc.CallOption) (*NiuEnterDeskAck, error)
}

type niuniuRpcClient struct {
	cc *grpc.ClientConn
}

func NewNiuniuRpcClient(cc *grpc.ClientConn) NiuniuRpcClient {
	return &niuniuRpcClient{cc}
}

func (c *niuniuRpcClient) CreateRoom(ctx context.Context, in *NiuCreateDeskReq, opts ...grpc.CallOption) (*NiuEnterDeskAck, error) {
	out := new(NiuEnterDeskAck)
	err := grpc.Invoke(ctx, "/ddproto.NiuniuRpc/CreateRoom", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NiuniuRpc service

type NiuniuRpcServer interface {
	// 创建房间
	CreateRoom(context.Context, *NiuCreateDeskReq) (*NiuEnterDeskAck, error)
}

func RegisterNiuniuRpcServer(s *grpc.Server, srv NiuniuRpcServer) {
	s.RegisterService(&_NiuniuRpc_serviceDesc, srv)
}

func _NiuniuRpc_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NiuCreateDeskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NiuniuRpcServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ddproto.NiuniuRpc/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NiuniuRpcServer).CreateRoom(ctx, req.(*NiuCreateDeskReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _NiuniuRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ddproto.NiuniuRpc",
	HandlerType: (*NiuniuRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _NiuniuRpc_CreateRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_niuniu.proto",
}

func init() { proto.RegisterFile("rpc_niuniu.proto", fileDescriptor51) }

var fileDescriptor51 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2a, 0x48, 0x8e,
	0xcf, 0xcb, 0x2c, 0xcd, 0xcb, 0x2c, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49,
	0x01, 0x33, 0xa4, 0x04, 0x21, 0xc2, 0xf1, 0x05, 0x39, 0x89, 0x95, 0x10, 0x39, 0xa3, 0x30, 0x2e,
	0x4e, 0x3f, 0xb0, 0x60, 0x50, 0x41, 0xb2, 0x90, 0x27, 0x17, 0x97, 0x73, 0x51, 0x6a, 0x62, 0x49,
	0x6a, 0x50, 0x7e, 0x7e, 0xae, 0x90, 0x8c, 0x1e, 0x54, 0x9f, 0x1e, 0x48, 0x4f, 0x32, 0x58, 0x22,
	0x3e, 0x25, 0xb5, 0x38, 0x3b, 0xbe, 0x28, 0xb5, 0x50, 0x4a, 0x1a, 0x45, 0x36, 0x35, 0xaf, 0x24,
	0xb5, 0x08, 0x22, 0x99, 0x98, 0x9c, 0xad, 0xc4, 0xe0, 0xc4, 0xe4, 0xc1, 0x1c, 0xc0, 0x00, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x2b, 0xf7, 0xeb, 0x6a, 0x8c, 0x00, 0x00, 0x00,
}
