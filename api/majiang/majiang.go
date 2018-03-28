package majiang

import (
	"casino_common/proto/ddproto/mjproto"
	"casino_common/proto/funcsInit"
	"casino_common/utils/numUtils"
	"strings"
)

var TOTALPAICOUNT int = 148 //4:万条筒 东南西北 中发白 1:春夏秋冬 梅兰菊竹 4:中
var MjpaiMap map[int]string
var clienMap map[int]int32

func init() {
	MjpaiMap = make(map[int]string, TOTALPAICOUNT)
	MjpaiMap[0] = "T_1"
	MjpaiMap[1] = "T_1"
	MjpaiMap[2] = "T_1"
	MjpaiMap[3] = "T_1"
	MjpaiMap[4] = "T_2"
	MjpaiMap[5] = "T_2"
	MjpaiMap[6] = "T_2"
	MjpaiMap[7] = "T_2"
	MjpaiMap[8] = "T_3"
	MjpaiMap[9] = "T_3"
	MjpaiMap[10] = "T_3"
	MjpaiMap[11] = "T_3"
	MjpaiMap[12] = "T_4"
	MjpaiMap[13] = "T_4"
	MjpaiMap[14] = "T_4"
	MjpaiMap[15] = "T_4"
	MjpaiMap[16] = "T_5"
	MjpaiMap[17] = "T_5"
	MjpaiMap[18] = "T_5"
	MjpaiMap[19] = "T_5"
	MjpaiMap[20] = "T_6"
	MjpaiMap[21] = "T_6"
	MjpaiMap[22] = "T_6"
	MjpaiMap[23] = "T_6"
	MjpaiMap[24] = "T_7"
	MjpaiMap[25] = "T_7"
	MjpaiMap[26] = "T_7"
	MjpaiMap[27] = "T_7"
	MjpaiMap[28] = "T_8"
	MjpaiMap[29] = "T_8"
	MjpaiMap[30] = "T_8"
	MjpaiMap[31] = "T_8"
	MjpaiMap[32] = "T_9"
	MjpaiMap[33] = "T_9"
	MjpaiMap[34] = "T_9"
	MjpaiMap[35] = "T_9"
	MjpaiMap[36] = "S_1"
	MjpaiMap[37] = "S_1"
	MjpaiMap[38] = "S_1"
	MjpaiMap[39] = "S_1"
	MjpaiMap[40] = "S_2"
	MjpaiMap[41] = "S_2"
	MjpaiMap[42] = "S_2"
	MjpaiMap[43] = "S_2"
	MjpaiMap[44] = "S_3"
	MjpaiMap[45] = "S_3"
	MjpaiMap[46] = "S_3"
	MjpaiMap[47] = "S_3"
	MjpaiMap[48] = "S_4"
	MjpaiMap[49] = "S_4"
	MjpaiMap[50] = "S_4"
	MjpaiMap[51] = "S_4"
	MjpaiMap[52] = "S_5"
	MjpaiMap[53] = "S_5"
	MjpaiMap[54] = "S_5"
	MjpaiMap[55] = "S_5"
	MjpaiMap[56] = "S_6"
	MjpaiMap[57] = "S_6"
	MjpaiMap[58] = "S_6"
	MjpaiMap[59] = "S_6"
	MjpaiMap[60] = "S_7"
	MjpaiMap[61] = "S_7"
	MjpaiMap[62] = "S_7"
	MjpaiMap[63] = "S_7"
	MjpaiMap[64] = "S_8"
	MjpaiMap[65] = "S_8"
	MjpaiMap[66] = "S_8"
	MjpaiMap[67] = "S_8"
	MjpaiMap[68] = "S_9"
	MjpaiMap[69] = "S_9"
	MjpaiMap[70] = "S_9"
	MjpaiMap[71] = "S_9"
	MjpaiMap[72] = "W_1"
	MjpaiMap[73] = "W_1"
	MjpaiMap[74] = "W_1"
	MjpaiMap[75] = "W_1"
	MjpaiMap[76] = "W_2"
	MjpaiMap[77] = "W_2"
	MjpaiMap[78] = "W_2"
	MjpaiMap[79] = "W_2"
	MjpaiMap[80] = "W_3"
	MjpaiMap[81] = "W_3"
	MjpaiMap[82] = "W_3"
	MjpaiMap[83] = "W_3"
	MjpaiMap[84] = "W_4"
	MjpaiMap[85] = "W_4"
	MjpaiMap[86] = "W_4"
	MjpaiMap[87] = "W_4"
	MjpaiMap[88] = "W_5"
	MjpaiMap[89] = "W_5"
	MjpaiMap[90] = "W_5"
	MjpaiMap[91] = "W_5"
	MjpaiMap[92] = "W_6"
	MjpaiMap[93] = "W_6"
	MjpaiMap[94] = "W_6"
	MjpaiMap[95] = "W_6"
	MjpaiMap[96] = "W_7"
	MjpaiMap[97] = "W_7"
	MjpaiMap[98] = "W_7"
	MjpaiMap[99] = "W_7"
	MjpaiMap[100] = "W_8"
	MjpaiMap[101] = "W_8"
	MjpaiMap[102] = "W_8"
	MjpaiMap[103] = "W_8"
	MjpaiMap[104] = "W_9"
	MjpaiMap[105] = "W_9"
	MjpaiMap[106] = "W_9"
	MjpaiMap[107] = "W_9"

	//东
	MjpaiMap[108] = "FENG_1_EAST"
	MjpaiMap[109] = "FENG_1_EAST"
	MjpaiMap[110] = "FENG_1_EAST"
	MjpaiMap[111] = "FENG_1_EAST"

	//南
	MjpaiMap[112] = "FENG_2_SOUTH"
	MjpaiMap[113] = "FENG_2_SOUTH"
	MjpaiMap[114] = "FENG_2_SOUTH"
	MjpaiMap[115] = "FENG_2_SOUTH"

	//西
	MjpaiMap[116] = "FENG_3_WEST"
	MjpaiMap[117] = "FENG_3_WEST"
	MjpaiMap[118] = "FENG_3_WEST"
	MjpaiMap[119] = "FENG_3_WEST"

	//北
	MjpaiMap[120] = "FENG_4_NORTH"
	MjpaiMap[121] = "FENG_4_NORTH"
	MjpaiMap[122] = "FENG_4_NORTH"
	MjpaiMap[123] = "FENG_4_NORTH"

	//中
	MjpaiMap[124] = "FENG_5_Z"
	MjpaiMap[125] = "FENG_5_Z"
	MjpaiMap[126] = "FENG_5_Z"
	MjpaiMap[127] = "FENG_5_Z"

	//白
	MjpaiMap[128] = "FENG_6_B"
	MjpaiMap[129] = "FENG_6_B"
	MjpaiMap[130] = "FENG_6_B"
	MjpaiMap[131] = "FENG_6_B"

	//发
	MjpaiMap[132] = "FENG_7_F"
	MjpaiMap[133] = "FENG_7_F"
	MjpaiMap[134] = "FENG_7_F"
	MjpaiMap[135] = "FENG_7_F"

	MjpaiMap[136] = "HUA_1_CHUN" //春夏秋冬
	MjpaiMap[137] = "HUA_2_XIA"
	MjpaiMap[138] = "HUA_3_QIU"
	MjpaiMap[139] = "HUA_4_DONG"

	MjpaiMap[140] = "HUA_5_MEI" //梅兰菊竹
	MjpaiMap[141] = "HUA_6_LAN"
	MjpaiMap[142] = "HUA_7_JU"
	MjpaiMap[143] = "HUA_8_ZHU"

	//中
	MjpaiMap[144] = "FENG_5_Z"
	MjpaiMap[145] = "FENG_5_Z"
	MjpaiMap[146] = "FENG_5_Z"
	MjpaiMap[147] = "FENG_5_Z"

	clienMap = make(map[int]int32, TOTALPAICOUNT)
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

	//东
	clienMap[108] = 28
	clienMap[109] = 28
	clienMap[110] = 28
	clienMap[111] = 28

	//南
	clienMap[112] = 29
	clienMap[113] = 29
	clienMap[114] = 29
	clienMap[115] = 29

	//西
	clienMap[116] = 30
	clienMap[117] = 30
	clienMap[118] = 30
	clienMap[119] = 30

	//北
	clienMap[120] = 31
	clienMap[121] = 31
	clienMap[122] = 31
	clienMap[123] = 31

	//中
	clienMap[124] = 32
	clienMap[125] = 32
	clienMap[126] = 32
	clienMap[127] = 32

	//白
	clienMap[128] = 33
	clienMap[129] = 33
	clienMap[130] = 33
	clienMap[131] = 33

	//发
	clienMap[132] = 34
	clienMap[133] = 34
	clienMap[134] = 34
	clienMap[135] = 34

	clienMap[136] = 35 //春
	clienMap[137] = 36 //夏
	clienMap[138] = 37 //秋
	clienMap[139] = 38 //冬

	clienMap[140] = 39 //梅
	clienMap[141] = 40 //兰
	clienMap[142] = 41 //菊
	clienMap[143] = 42 //竹

	//中
	clienMap[144] = 32
	clienMap[145] = 32
	clienMap[146] = 32
	clienMap[147] = 32
}

//番数 顶番5

type PengPai struct {
	OutUserId uint32
	InUserId  uint32
	Pais      []*MJPAI
}

type FlyPai struct {
	OutUserId uint32
	InUserId  uint32
	Pais      []*MJPAI //三张牌 含听用
	Pai       *MJPAI   //飞的是什么牌
}

type GangPai struct {
	GangType   int32
	Pais       []*MJPAI
	Pai        *MJPAI
	GetUserId  uint32
	SendUserId uint32
}

type ChiPai struct {
	OutUserId uint32
	InUserId  uint32
	Pais      []*MJPAI
}

type HuPai struct {
	Pai *MJPAI
}

type MJ_FLOWER int32

var (
	MJ_FLOWER_ERROR MJ_FLOWER = 0

	//万条同，分别是123
	MJ_FLOWER_W MJ_FLOWER = 1
	MJ_FLOWER_S MJ_FLOWER = 2
	MJ_FLOWER_T MJ_FLOWER = 3

	//风牌 中发白东南西北 分别是1234567
	MJ_FLOWER_FENG MJ_FLOWER = 4

	//花牌 春夏秋冬 梅兰菊竹 分别是12345678
	MJ_FLOWER_HUA MJ_FLOWER = 5
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
	switch flowerStr {
	case "T": //筒
		p.Flower = MJ_FLOWER_T
	case "S": //索
		p.Flower = MJ_FLOWER_S
	case "W": //万
		p.Flower = MJ_FLOWER_W
	case "FENG": //风牌
		p.Flower = MJ_FLOWER_FENG
	case "HUA":
		p.Flower = MJ_FLOWER_HUA

	default:
		p.Flower = MJ_FLOWER_ERROR
	}

	//初始化大小
	p.Value = int32(numUtils.String2Int(sarry[1]))
	return nil

}

//这里得到牌的计数index
func (p *MJPAI) GetCountIndex() int32 {
	return p.Value - 1 + (int32(p.Flower-1))*9 //todo 此方法是否合适...
}

//todo
func (p *MJPAI) GetCardInfo() *mjproto.CardInfo {
	if p == nil {
		return &mjproto.CardInfo{}
	}
	cardInfo := commonNewPorot.NewCardInfo()
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

	idStr, _ := numUtils.Int2String(p.Index)
	des := ""

	switch p.Flower {
	case MJ_FLOWER_HUA:
		//花牌使用花牌的描述
		des = p.DesHua()
	case MJ_FLOWER_FENG:
		//风牌使用风牌的描述
		des = p.DesFeng()
	default:
		//其他字牌用字牌的描述
		valueStr, _ := numUtils.Int2String(p.Value)
		des = valueStr + GetFlow(int32(p.Flower))
	}

	return idStr + "-" + des
}

//风牌的描述
func (p *MJPAI) DesFeng() string {
	if p.Flower != MJ_FLOWER_FENG {
		return "未知风"
	}
	des := ""
	switch {
	case p.Value == 1:
		des = "东"
	case p.Value == 2:
		des = "南"
	case p.Value == 3:
		des = "西"
	case p.Value == 4:
		des = "北"
	case p.Value == 5:
		des = "中"
	case p.Value == 6:
		des = "白"
	case p.Value == 7:
		des = "发"
	default:
		des = "默认风"
	}
	return des
}

//花牌的描述
func (p *MJPAI) DesHua() string {
	if p.Flower != MJ_FLOWER_HUA {
		return "未知花"
	}
	des := ""
	switch {
	case p.Value == 1:
		des = "春"
	case p.Value == 2:
		des = "夏"
	case p.Value == 3:
		des = "秋"
	case p.Value == 4:
		des = "东"
	case p.Value == 5:
		des = "梅"
	case p.Value == 6:
		des = "兰"
	case p.Value == 7:
		des = "菊"
	case p.Value == 8:
		des = "竹"
	default:
		des = "默认花"
	}
	return des
}

func GetFlow(f int32) string {
	switch MJ_FLOWER(f) {
	case MJ_FLOWER_W:
		return "万"
	case MJ_FLOWER_T:
		return "筒"
	case MJ_FLOWER_S:
		return "索"
	case MJ_FLOWER_FENG:
		return "风"
	case MJ_FLOWER_HUA:
		return "花"
	default:
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
	}

	if list[i].Flower == list[j].Flower {
		if list[i].Value < list[j].Value {
			return true
		}
		if list[i].Value == list[j].Value {
			if list[i].Index < list[j].Index {
				return true
			}
			return false
		}
		return false
	}
	return false
}

func (list MjPAIList) Swap(i, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}

func (list MjPAIList) Indexs() (indexs []int32) {
	for _, i := range list {
		indexs = append(indexs, i.GetIndex())
	}
	return
}
