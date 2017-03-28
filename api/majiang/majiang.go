package majiang

type MJ_FLOWER int32

var (
	MJ_FLOWER_T MJ_FLOWER = 1
	MJ_FLOWER_S MJ_FLOWER = 2
	MJ_FLOWER_W MJ_FLOWER = 3
)

type PengPai struct {
	OutUserId uint32
}

type GangPai struct {
}

type ChiPai struct {
}

//麻将牌的结构
type MJPAI struct {
	Index  int32
	Value  int32
	Flower MJ_FLOWER //花色
}
