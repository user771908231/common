// Code generated by protoc-gen-go.
// source: hot_update.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of common_enum_game from common_enum.proto

type VersionInfo struct {
	FileId           *int32 `protobuf:"varint,1,opt,name=fileId" json:"fileId,omitempty"`
	FileVer          *int32 `protobuf:"varint,2,opt,name=fileVer" json:"fileVer,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *VersionInfo) Reset()                    { *m = VersionInfo{} }
func (m *VersionInfo) String() string            { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()               {}
func (*VersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *VersionInfo) GetFileId() int32 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *VersionInfo) GetFileVer() int32 {
	if m != nil && m.FileVer != nil {
		return *m.FileVer
	}
	return 0
}

type AssetInfo struct {
	FileId           *int32          `protobuf:"varint,1,opt,name=fileId" json:"fileId,omitempty"`
	FileVer          *int32          `protobuf:"varint,2,opt,name=fileVer" json:"fileVer,omitempty"`
	FilePath         *string         `protobuf:"bytes,3,opt,name=filePath" json:"filePath,omitempty"`
	FileSize         *int32          `protobuf:"varint,4,opt,name=fileSize" json:"fileSize,omitempty"`
	Md5              *string         `protobuf:"bytes,5,opt,name=md5" json:"md5,omitempty"`
	IsCompress       *bool           `protobuf:"varint,6,opt,name=isCompress" json:"isCompress,omitempty"`
	IsCode           *bool           `protobuf:"varint,7,opt,name=isCode" json:"isCode,omitempty"`
	GameId           *CommonEnumGame `protobuf:"varint,8,opt,name=gameId,enum=ddproto.CommonEnumGame" json:"gameId,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *AssetInfo) Reset()                    { *m = AssetInfo{} }
func (m *AssetInfo) String() string            { return proto.CompactTextString(m) }
func (*AssetInfo) ProtoMessage()               {}
func (*AssetInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

func (m *AssetInfo) GetFileId() int32 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *AssetInfo) GetFileVer() int32 {
	if m != nil && m.FileVer != nil {
		return *m.FileVer
	}
	return 0
}

func (m *AssetInfo) GetFilePath() string {
	if m != nil && m.FilePath != nil {
		return *m.FilePath
	}
	return ""
}

func (m *AssetInfo) GetFileSize() int32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *AssetInfo) GetMd5() string {
	if m != nil && m.Md5 != nil {
		return *m.Md5
	}
	return ""
}

func (m *AssetInfo) GetIsCompress() bool {
	if m != nil && m.IsCompress != nil {
		return *m.IsCompress
	}
	return false
}

func (m *AssetInfo) GetIsCode() bool {
	if m != nil && m.IsCode != nil {
		return *m.IsCode
	}
	return false
}

func (m *AssetInfo) GetGameId() CommonEnumGame {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return CommonEnumGame_GID_SRC
}

// 获取资源文件的最新版本信息
type HotupdateReqVersionInfo struct {
	Header              *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId         *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	DownloadGameId      []int32      `protobuf:"varint,3,rep,name=downloadGameId" json:"downloadGameId,omitempty"`
	ClientAssetsVersion *int32       `protobuf:"varint,4,opt,name=clientAssetsVersion" json:"clientAssetsVersion,omitempty"`
	XXX_unrecognized    []byte       `json:"-"`
}

func (m *HotupdateReqVersionInfo) Reset()                    { *m = HotupdateReqVersionInfo{} }
func (m *HotupdateReqVersionInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqVersionInfo) ProtoMessage()               {}
func (*HotupdateReqVersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

func (m *HotupdateReqVersionInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqVersionInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetDownloadGameId() []int32 {
	if m != nil {
		return m.DownloadGameId
	}
	return nil
}

func (m *HotupdateReqVersionInfo) GetClientAssetsVersion() int32 {
	if m != nil && m.ClientAssetsVersion != nil {
		return *m.ClientAssetsVersion
	}
	return 0
}

type HotupdateAckVersionInfo struct {
	Header               *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Version              []*VersionInfo `protobuf:"bytes,2,rep,name=version" json:"version,omitempty"`
	LastestAssetsVersion *int32         `protobuf:"varint,3,opt,name=lastestAssetsVersion" json:"lastestAssetsVersion,omitempty"`
	ForceUpdate          *bool          `protobuf:"varint,4,opt,name=forceUpdate" json:"forceUpdate,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *HotupdateAckVersionInfo) Reset()                    { *m = HotupdateAckVersionInfo{} }
func (m *HotupdateAckVersionInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckVersionInfo) ProtoMessage()               {}
func (*HotupdateAckVersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

func (m *HotupdateAckVersionInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckVersionInfo) GetVersion() []*VersionInfo {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *HotupdateAckVersionInfo) GetLastestAssetsVersion() int32 {
	if m != nil && m.LastestAssetsVersion != nil {
		return *m.LastestAssetsVersion
	}
	return 0
}

func (m *HotupdateAckVersionInfo) GetForceUpdate() bool {
	if m != nil && m.ForceUpdate != nil {
		return *m.ForceUpdate
	}
	return false
}

// 获取资源文件的详细信息
type HotupdateReqAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId      *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	FileIds          []int32      `protobuf:"varint,3,rep,name=fileIds" json:"fileIds,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateReqAssetsInfo) Reset()                    { *m = HotupdateReqAssetsInfo{} }
func (m *HotupdateReqAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqAssetsInfo) ProtoMessage()               {}
func (*HotupdateReqAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{4} }

func (m *HotupdateReqAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqAssetsInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqAssetsInfo) GetFileIds() []int32 {
	if m != nil {
		return m.FileIds
	}
	return nil
}

type HotupdateAckAssetsInfo struct {
	Header               *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AssetHost            *string      `protobuf:"bytes,2,opt,name=assetHost" json:"assetHost,omitempty"`
	Assets               []*AssetInfo `protobuf:"bytes,3,rep,name=assets" json:"assets,omitempty"`
	LastestAssetsVersion *int32       `protobuf:"varint,4,opt,name=lastestAssetsVersion" json:"lastestAssetsVersion,omitempty"`
	XXX_unrecognized     []byte       `json:"-"`
}

func (m *HotupdateAckAssetsInfo) Reset()                    { *m = HotupdateAckAssetsInfo{} }
func (m *HotupdateAckAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckAssetsInfo) ProtoMessage()               {}
func (*HotupdateAckAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{5} }

func (m *HotupdateAckAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckAssetsInfo) GetAssetHost() string {
	if m != nil && m.AssetHost != nil {
		return *m.AssetHost
	}
	return ""
}

func (m *HotupdateAckAssetsInfo) GetAssets() []*AssetInfo {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *HotupdateAckAssetsInfo) GetLastestAssetsVersion() int32 {
	if m != nil && m.LastestAssetsVersion != nil {
		return *m.LastestAssetsVersion
	}
	return 0
}

// 获取某个独立游戏的下载信息
type HotupdateReqGameAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId      *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	GameId           *int32       `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateReqGameAssetsInfo) Reset()                    { *m = HotupdateReqGameAssetsInfo{} }
func (m *HotupdateReqGameAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqGameAssetsInfo) ProtoMessage()               {}
func (*HotupdateReqGameAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{6} }

func (m *HotupdateReqGameAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqGameAssetsInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqGameAssetsInfo) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

type HotupdateAckGameAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AssetHost        *string      `protobuf:"bytes,2,opt,name=assetHost" json:"assetHost,omitempty"`
	Assets           []*AssetInfo `protobuf:"bytes,3,rep,name=assets" json:"assets,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateAckGameAssetsInfo) Reset()                    { *m = HotupdateAckGameAssetsInfo{} }
func (m *HotupdateAckGameAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckGameAssetsInfo) ProtoMessage()               {}
func (*HotupdateAckGameAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{7} }

func (m *HotupdateAckGameAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckGameAssetsInfo) GetAssetHost() string {
	if m != nil && m.AssetHost != nil {
		return *m.AssetHost
	}
	return ""
}

func (m *HotupdateAckGameAssetsInfo) GetAssets() []*AssetInfo {
	if m != nil {
		return m.Assets
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionInfo)(nil), "ddproto.VersionInfo")
	proto.RegisterType((*AssetInfo)(nil), "ddproto.AssetInfo")
	proto.RegisterType((*HotupdateReqVersionInfo)(nil), "ddproto.hotupdate_req_versionInfo")
	proto.RegisterType((*HotupdateAckVersionInfo)(nil), "ddproto.hotupdate_ack_versionInfo")
	proto.RegisterType((*HotupdateReqAssetsInfo)(nil), "ddproto.hotupdate_req_assetsInfo")
	proto.RegisterType((*HotupdateAckAssetsInfo)(nil), "ddproto.hotupdate_ack_assetsInfo")
	proto.RegisterType((*HotupdateReqGameAssetsInfo)(nil), "ddproto.hotupdate_req_gameAssetsInfo")
	proto.RegisterType((*HotupdateAckGameAssetsInfo)(nil), "ddproto.hotupdate_ack_gameAssetsInfo")
}

var fileDescriptor16 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0xc9, 0x84, 0xc9, 0xcf, 0x8d, 0x14, 0x66, 0xdc, 0x11, 0xf2, 0x94, 0x2e, 0x2a, 0x0b,
	0x24, 0xd8, 0x64, 0x51, 0x89, 0x07, 0x28, 0x2c, 0x68, 0x77, 0x91, 0x10, 0x6c, 0x23, 0x2b, 0x76,
	0x68, 0x44, 0x12, 0x87, 0xd8, 0x05, 0x09, 0x1e, 0x81, 0x35, 0x2b, 0x5e, 0x81, 0x87, 0xc4, 0x3f,
	0xa1, 0x24, 0x48, 0x48, 0x50, 0xd8, 0x54, 0xf5, 0xf1, 0xbd, 0x39, 0xe7, 0x7e, 0xbe, 0x70, 0x75,
	0x10, 0xaa, 0x38, 0xf6, 0x8c, 0x2a, 0x9e, 0xf5, 0x83, 0x50, 0x02, 0x85, 0x8c, 0xd9, 0x3f, 0xcb,
	0x45, 0x29, 0xda, 0x56, 0x74, 0x45, 0xd9, 0xd4, 0xbc, 0x53, 0xee, 0x76, 0x79, 0x3d, 0x8a, 0xbc,
	0x3b, 0xb6, 0x4e, 0x22, 0x19, 0x24, 0xaf, 0xf9, 0x20, 0x6b, 0xd1, 0xed, 0xbb, 0x4a, 0xa0, 0x14,
	0x82, 0xaa, 0x6e, 0xf8, 0x9e, 0x61, 0x6f, 0xed, 0x3d, 0xbe, 0x44, 0xf7, 0x20, 0x34, 0x67, 0x5d,
	0x82, 0x2f, 0x8c, 0x40, 0xbe, 0x79, 0x10, 0x6f, 0xa5, 0xe4, 0xea, 0x8f, 0xca, 0xd1, 0x15, 0x44,
	0x46, 0xc8, 0xa9, 0x3a, 0x60, 0x5f, 0x2b, 0xf1, 0x0f, 0xe5, 0x65, 0xfd, 0x91, 0xe3, 0xbb, 0xb6,
	0x26, 0x01, 0xbf, 0x65, 0x4f, 0xf1, 0xa5, 0xbd, 0x46, 0x00, 0xb5, 0x7c, 0x2e, 0xda, 0x7e, 0xe0,
	0x52, 0xe2, 0x40, 0x6b, 0x91, 0x71, 0x31, 0x1a, 0xe3, 0x38, 0xb4, 0xe7, 0x27, 0x10, 0xbc, 0xa1,
	0xad, 0x71, 0x8d, 0xf4, 0x39, 0xdd, 0xdc, 0x66, 0xe3, 0xd4, 0xd9, 0x64, 0xbe, 0xc2, 0x94, 0x90,
	0xcf, 0x1e, 0xdc, 0x6a, 0x48, 0x8e, 0x51, 0x31, 0xf0, 0x77, 0xc5, 0xfb, 0xc9, 0xb4, 0x0f, 0x21,
	0x38, 0x70, 0xca, 0x74, 0x5a, 0x13, 0x3f, 0xd9, 0xdc, 0x9c, 0x3e, 0x94, 0x9b, 0xdf, 0x9d, 0xbd,
	0x43, 0x0b, 0x48, 0x1c, 0xc5, 0x6d, 0xdf, 0x6b, 0x4f, 0x37, 0xd8, 0x7d, 0x48, 0x99, 0xf8, 0xd0,
	0x35, 0x82, 0xb2, 0x17, 0x2e, 0x8b, 0xbf, 0xf6, 0xb5, 0xfe, 0x00, 0x16, 0x63, 0xb1, 0x81, 0x24,
	0x47, 0xb6, 0x6e, 0x52, 0xf2, 0x75, 0x96, 0x86, 0x96, 0x6f, 0xcf, 0x48, 0xf3, 0x08, 0xc2, 0xb1,
	0x49, 0x27, 0xf1, 0x67, 0x65, 0xd3, 0x87, 0x5c, 0xc1, 0x4d, 0x43, 0xa5, 0xe2, 0xf2, 0x97, 0x20,
	0xbe, 0x4d, 0xaf, 0x47, 0xaa, 0xc4, 0x50, 0xf2, 0x57, 0x36, 0x89, 0x4d, 0x17, 0x91, 0x0a, 0xf0,
	0x1c, 0x15, 0xb5, 0x9d, 0xff, 0x4a, 0x6a, 0xdc, 0x89, 0x3d, 0x93, 0x0e, 0x11, 0xf9, 0xe2, 0x4d,
	0x8d, 0x0c, 0x85, 0xbf, 0x36, 0xba, 0x86, 0xd8, 0xf6, 0xec, 0x84, 0x54, 0xd6, 0x26, 0x46, 0x04,
	0x02, 0xf7, 0x19, 0xeb, 0x92, 0x6c, 0xd0, 0xa9, 0xf1, 0xe7, 0xba, 0xfe, 0x0e, 0x8a, 0x7b, 0x9d,
	0x1a, 0x56, 0xf3, 0xf9, 0xcd, 0x06, 0x6d, 0xff, 0x0b, 0x83, 0xf4, 0xb4, 0xb1, 0x96, 0x3f, 0xf9,
	0x34, 0xb5, 0x32, 0x04, 0xce, 0xb2, 0x3a, 0x8f, 0xc2, 0xb3, 0x8b, 0x9d, 0x9f, 0xdf, 0xc9, 0xbd,
	0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xf6, 0x58, 0xa2, 0x3e, 0x04, 0x00, 0x00,
}
