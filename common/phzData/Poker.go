package phzData

import (
	"casino_common/common/log"
	"casino_common/utils/chessUtils"
	"casino_common/utils/numUtils"
	"strings"
)

const PHZ_ALLPOKER_NUM int32 = 80 //所有牌的数量

var PokerMap map[int32]string

func init() {
	PokerMap = make(map[int32]string, PHZ_ALLPOKER_NUM)
	PokerMap[0] = ""
	PokerMap[1] = "Black_Y_11"
	PokerMap[2] = "Black_Y_11"
	PokerMap[3] = "Black_Y_11"
	PokerMap[4] = "Black_Y_11"
	PokerMap[5] = "Black_N_1"
	PokerMap[6] = "Black_N_1"
	PokerMap[7] = "Black_N_1"
	PokerMap[8] = "Black_N_1"
	PokerMap[9] = "Red_Y_12"
	PokerMap[10] = "Red_Y_12"
	PokerMap[11] = "Red_Y_12"
	PokerMap[12] = "Red_Y_12"
	PokerMap[13] = "Red_N_2"
	PokerMap[14] = "Red_N_2"
	PokerMap[15] = "Red_N_2"
	PokerMap[16] = "Red_N_2"
	PokerMap[17] = "Black_Y_13"
	PokerMap[18] = "Black_Y_13"
	PokerMap[19] = "Black_Y_13"
	PokerMap[20] = "Black_Y_13"
	PokerMap[21] = "Black_N_3"
	PokerMap[22] = "Black_N_3"
	PokerMap[23] = "Black_N_3"
	PokerMap[24] = "Black_N_3"
	PokerMap[25] = "Black_Y_14"
	PokerMap[26] = "Black_Y_14"
	PokerMap[27] = "Black_Y_14"
	PokerMap[28] = "Black_Y_14"
	PokerMap[29] = "Black_N_4"
	PokerMap[30] = "Black_N_4"
	PokerMap[31] = "Black_N_4"
	PokerMap[32] = "Black_N_4"
	PokerMap[33] = "Black_Y_15"
	PokerMap[34] = "Black_Y_15"
	PokerMap[35] = "Black_Y_15"
	PokerMap[36] = "Black_Y_15"
	PokerMap[37] = "Black_N_5"
	PokerMap[38] = "Black_N_5"
	PokerMap[39] = "Black_N_5"
	PokerMap[40] = "Black_N_5"
	PokerMap[41] = "Black_Y_16"
	PokerMap[42] = "Black_Y_16"
	PokerMap[43] = "Black_Y_16"
	PokerMap[44] = "Black_Y_16"
	PokerMap[45] = "Black_N_6"
	PokerMap[46] = "Black_N_6"
	PokerMap[47] = "Black_N_6"
	PokerMap[48] = "Black_N_6"
	PokerMap[49] = "Red_Y_17"
	PokerMap[50] = "Red_Y_17"
	PokerMap[51] = "Red_Y_17"
	PokerMap[52] = "Red_Y_17"
	PokerMap[53] = "Red_N_7"
	PokerMap[54] = "Red_N_7"
	PokerMap[55] = "Red_N_7"
	PokerMap[56] = "Red_N_7"
	PokerMap[57] = "Black_Y_18"
	PokerMap[58] = "Black_Y_18"
	PokerMap[59] = "Black_Y_18"
	PokerMap[60] = "Black_Y_18"
	PokerMap[61] = "Black_N_8"
	PokerMap[62] = "Black_N_8"
	PokerMap[63] = "Black_N_8"
	PokerMap[64] = "Black_N_8"
	PokerMap[65] = "Black_Y_19"
	PokerMap[66] = "Black_Y_19"
	PokerMap[67] = "Black_Y_19"
	PokerMap[68] = "Black_Y_19"
	PokerMap[69] = "Black_N_9"
	PokerMap[70] = "Black_N_9"
	PokerMap[71] = "Black_N_9"
	PokerMap[72] = "Black_N_9"
	PokerMap[73] = "Red_Y_20"
	PokerMap[74] = "Red_Y_20"
	PokerMap[75] = "Red_Y_20"
	PokerMap[76] = "Red_Y_20"
	PokerMap[77] = "Red_N_10"
	PokerMap[78] = "Red_N_10"
	PokerMap[79] = "Red_N_10"
	PokerMap[80] = "Red_N_10"
}

type PengPai struct {
	PengType  int32 //碰牌的类型：偎牌，碰，臭偎
	Pais      []*PHZPoker
	OutUserId uint32
	InUserId  uint32
	HuXi      int32
}

type PaoPai struct {
	Pais      []*PHZPoker
	PaoType   int32 //跑牌类型
	OutUserId uint32
	InUserId  uint32
	HuXi      int32
}

type TiPai struct {
	TiType    int32 //提牌的类型
	Pais      []*PHZPoker
	OutUserId uint32
	InUserId  uint32
	HuXi      int32
}

type ChiPai struct {
	Pais      []*PHZPoker
	Pai       *PHZPoker
	OutUserId uint32
	InUserId  uint32
	HuXi      int32
}

type WeiPai struct {
	Pais []*PHZPoker
	Pai  *PHZPoker
	HuXi int32
}

type TiInfo struct {
}

type PHZPoker struct {
	Id      int32  //每张牌唯一的ID
	Value   int32  //牌值
	Flower  FLOWER //花色
	BigWord bool   //true：大字  false：小字
	Des     string //描述
}

type FLOWER int32

var (
	MJ_FLOWER_ERROR FLOWER = 0
	FLOWER_R        FLOWER = 1 //红字
	FLOWER_B        FLOWER = 2 //黑字
)

func (p *PHZPoker) GetId() int32 {
	return p.Id
}

func (p *PHZPoker) GetValue() int32 {
	return p.Value
}

func (p *PHZPoker) GetFlower() int32 {
	return int32(p.Flower)
}

func (p *PHZPoker) GetDes() string {
	return p.Des
}

func (p *PHZPoker) IsBig() bool {
	return p.BigWord
}

func initByDes(des string) FLOWER {
	//初始化花色
	var ret FLOWER
	switch des {
	case "Red": //红字
		ret = FLOWER_R
	case "Black": //黑字
		ret = FLOWER_B
	default:
		ret = MJ_FLOWER_ERROR
	}
	return ret
}

func initBigOrSmall(des string) bool {
	var ret bool
	switch des {
	case "N":
		ret = false
	case "Y":
		ret = true
	default:
		ret = false
	}
	return ret
}

func parserByIndex(index int32) (id int32, value int32, isBig bool, flower FLOWER, des string) {
	var pokerString string = PokerMap[index]
	sarry := strings.Split(pokerString, "_")
	flower = initByDes(sarry[0])
	isBig = initBigOrSmall(sarry[1])
	value = int32(numUtils.String2Int(sarry[2]))
	return index, value, isBig, flower, pokerString
}

func InitPaiByIndex(index int32) *PHZPoker {
	if index < 0 || index > PHZ_ALLPOKER_NUM {
		log.E("解析牌失败...ID非法")
		return nil
	}
	_, value, isBig, flower, des := parserByIndex(index)

	poker := &PHZPoker{
		Id:      index,
		Value:   value,
		Flower:  flower,
		BigWord: isBig,
		Des:     des,
	}
	return poker
}

func InitPaisByArray(array []int32) []*PHZPoker {
	if array == nil {
		log.E("根据索引数组初始化失败...")
		return nil
	}
	pokers := []*PHZPoker{}
	for _, index := range array {
		p := InitPaiByIndex(index)
		pokers = append(pokers, p)
	}
	return pokers
}

func XiPai() []*PHZPoker {
	paiIndex := chessUtils.Xipai(1, int(PHZ_ALLPOKER_NUM))
	log.T("一局洗出来的牌是：[%v]", paiIndex)
	pokerPais := []*PHZPoker{}
	for _, index := range paiIndex {
		p := InitPaiByIndex(index)
		if p == nil {
			return nil
		}
		pokerPais = append(pokerPais, p)
	}
	return pokerPais
}

func (p *PHZPoker) GetLogDes() string {
	suit, _ := numUtils.Int2String(p.GetId())
	switch p.GetFlower() {
	case int32(FLOWER_B):
		suit = "黑"
	case int32(FLOWER_R):
		suit = "红"
	default:
	}
	if p.IsBig() {
		suit += "大"
	} else {
		suit += "小"
	}
	valueString := ""
	//valueString, _ := numUtils.Int2String(p.GetValue())
	if p.GetValue() > 10 {
		valueString, _ = numUtils.Int2String(p.GetValue() - 10)
	} else {
		valueString, _ = numUtils.Int2String(p.GetValue())
	}
	suit += valueString
	return suit
}
func Cards2String(cs []*PHZPoker) string {
	if cs == nil || len(cs) <= 0 {
		return "没有牌"
	}

	s := ""
	for _, c := range cs {
		if c != nil {
			s += " " + c.GetLogDes()
		}
	}
	return s
}

func Card2String(p *PHZPoker) string {
	return p.GetLogDes()
}

func (p *PHZPoker) GetPaiIndexByValue() int32 {
	//根据牌的value和花色返回index，此处的index用于统计牌的张数
	return p.GetValue()
}
