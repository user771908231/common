package majiang

import (
	"casino_majiang/msg/protogo"
	"strings"
	"casino_common/utils/numUtils"
	"casino_majiang/msg/funcsInit"
)

var clienMap map[int]int32

func init() {

	clienMap = make(map[int]int32, 108)
	clienMap[0] = 19
	clienMap[1] = 19
	clienMap[2] = 19
	clienMap[3] = 19

	clienMap[4] = 20
	clienMap[5] = 20
	clienMap[6] = 20
	clienMap[7] = 20

	clienMap[8] = 21
	clienMap[9] = 21
	clienMap[10] = 21
	clienMap[11] = 21

	clienMap[12] = 22
	clienMap[13] = 22
	clienMap[14] = 22
	clienMap[15] = 22

	clienMap[16] = 23
	clienMap[17] = 23
	clienMap[18] = 23
	clienMap[19] = 23

	clienMap[20] = 24
	clienMap[21] = 24
	clienMap[22] = 24
	clienMap[23] = 24

	clienMap[24] = 25
	clienMap[25] = 25
	clienMap[26] = 25
	clienMap[27] = 25

	clienMap[28] = 26
	clienMap[29] = 26
	clienMap[30] = 26
	clienMap[31] = 26

	clienMap[32] = 27
	clienMap[33] = 27
	clienMap[34] = 27
	clienMap[35] = 27

	clienMap[36] = 10
	clienMap[37] = 10
	clienMap[38] = 10
	clienMap[39] = 10

	clienMap[40] = 11
	clienMap[41] = 11
	clienMap[42] = 11
	clienMap[43] = 11

	clienMap[44] = 12
	clienMap[45] = 12
	clienMap[46] = 12
	clienMap[47] = 12

	clienMap[48] = 13
	clienMap[49] = 13
	clienMap[50] = 13
	clienMap[51] = 13

	clienMap[52] = 14
	clienMap[53] = 14
	clienMap[54] = 14
	clienMap[55] = 14

	clienMap[56] = 15
	clienMap[57] = 15
	clienMap[58] = 15
	clienMap[59] = 15

	clienMap[60] = 16
	clienMap[61] = 16
	clienMap[62] = 16
	clienMap[63] = 16

	clienMap[64] = 17
	clienMap[65] = 17
	clienMap[66] = 17
	clienMap[67] = 17

	clienMap[68] = 18
	clienMap[69] = 18
	clienMap[70] = 18
	clienMap[71] = 18

	clienMap[72] = 1
	clienMap[73] = 1
	clienMap[74] = 1
	clienMap[75] = 1

	clienMap[76] = 2
	clienMap[77] = 2
	clienMap[78] = 2
	clienMap[79] = 2

	clienMap[80] = 3
	clienMap[81] = 3
	clienMap[82] = 3
	clienMap[83] = 3

	clienMap[84] = 4
	clienMap[85] = 4
	clienMap[86] = 4
	clienMap[87] = 4

	clienMap[88] = 5
	clienMap[89] = 5
	clienMap[90] = 5
	clienMap[91] = 5

	clienMap[92] = 6
	clienMap[93] = 6
	clienMap[94] = 6
	clienMap[95] = 6

	clienMap[96] = 7
	clienMap[97] = 7
	clienMap[98] = 7
	clienMap[99] = 7

	clienMap[100] = 8
	clienMap[101] = 8
	clienMap[102] = 8
	clienMap[103] = 8

	clienMap[104] = 9
	clienMap[105] = 9
	clienMap[106] = 9
	clienMap[107] = 9
}

//番数 顶番5

type PengPai struct {
	OutUserId uint32
	InUserId  uint32
	Pais      []*MJPAI
}

type GangPai struct {
	GangType   int32
	Pais       []*MJPAI
	Pai        *MJPAI
	GetUserId  uint32
	SendUserId uint32
}

type ChiPai struct {
}

type HuPai struct {
	Pai *MJPAI
}

type MJ_FLOWER int32

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
	Des    string
}

func (p *MJPAI) GetIndex() int32 {
	return p.Index
}

func (p *MJPAI) InitByDes() error {
	//拆分描述
	sarry := strings.Split(p.Des, "_")
	flowerStr := sarry[0]

	//初始化花色
	if flowerStr == "T" {
		p.Flower = MJ_FLOWER_T
	} else if flowerStr == "S" {
		p.Flower = MJ_FLOWER_S
	} else {
		p.Flower = MJ_FLOWER_W
	}
	//初始化大小
	p.Value = int32(numUtils.String2Int(sarry[1]))
	return nil

}

//这里得到牌的计数index
func (p *MJPAI) GetCountIndex() int32 {
	return p.Value - 1 + (int32(p.Flower - 1))*9 //todo 此方法是否合适...
}

//todo
func (p *MJPAI) GetCardInfo() *mjproto.CardInfo {
	if p == nil {
		return &mjproto.CardInfo{}
	}
	cardInfo := newProto.NewCardInfo()
	*cardInfo.Id = p.Index
	*cardInfo.Type = int32(p.Flower)
	*cardInfo.Value = p.GetClientId()
	return cardInfo
}

//todo
func (p *MJPAI) GetClientId() int32 {
	return clienMap[int(p.Index)]
}

func (p *MJPAI) LogDes() string {
	if p == nil {
		return "空牌"
	}
	valueStr, _ := numUtils.Int2String(p.Value)
	idStr, _ := numUtils.Int2String(p.Index)
	return idStr + "-" + valueStr + GetFlow(int32(p.Flower))
}

func GetFlow(f int32) string {
	if f == 1 {
		return "万"
	} else if f == 2 {
		return "条"
	} else if f == 3 {
		return "筒"
	} else {
		return "白"
	}

}

//-----------------------------------------------------------------排序--------------------------------------

type MjPAIList []*MJPAI

func (list MjPAIList) Len() int {
	return len(list)

}

func (list MjPAIList) Less(i, j int) bool {
	if list[i].Flower < list[j].Flower {
		return true
	} else if list[i].Flower == list[j].Flower {
		if list[i].Value < list[j].Value {
			return true
		} else {
			return false
		}

	} else {
		return false
	}

}
func (list MjPAIList) Swap(i, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}
