package majiang

import "casino_majiang/msg/protogo"

type MJ_FLOWER int32

type PengPai struct {
	OutUserId uint32
	Pais      []*MJPAI
}

type GangPai struct {
	GangType int32
	Pai      *MJPAI
}

type ChiPai struct {
}

type HuPai struct {
	Pai *MJPAI
}

var (
	//万条同，分别是123
	MJ_FLOWER_W MJ_FLOWER = 1
	MJ_FLOWER_S MJ_FLOWER = 2
	MJ_FLOWER_T MJ_FLOWER = 3
)

//麻将牌的结构
type MJPAI struct {
	Index  int32
	Value  int32
	Flower MJ_FLOWER //花色
}

//这里得到牌的计数index
func (p *MJPAI) GetCountIndex() int32 {
	return p.Value - 1 + (int32(p.Flower - 1))*9 //todo 此方法是否合适...
}

//todo
func (p *MJPAI) GetCardInfo() *mjproto.CardInfo {
	return nil
}

//todo
func (p *MJPAI) GetClientId() int32 {
	return 0
}

//todo
func (p *MJPAI) LogDes() string {
	return ""
}
