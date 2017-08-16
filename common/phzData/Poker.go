package phzData

const PHZ_ALLPOKER_NUM int32 = 80 //所有牌的数量

var PokerMap map[int32]string

func init() {
	PokerMap = make(map[int32]string, PHZ_ALLPOKER_NUM)
	PokerMap[0] = ""
	PokerMap[1] = "一_1"
	PokerMap[2] = "一_1"
	PokerMap[3] = "一_1"
	PokerMap[4] = "一_1"
	PokerMap[5] = "1_1"
	PokerMap[6] = "1_1"
	PokerMap[7] = "1_1"
	PokerMap[8] = "1_1"
	PokerMap[9] = "二_2"
	PokerMap[10] = "二_2"
	PokerMap[11] = "二_2"
	PokerMap[12] = "二_2"
	PokerMap[13] = "2_2"
	PokerMap[14] = "2_2"
	PokerMap[15] = "2_2"
	PokerMap[16] = "2_2"
	PokerMap[17] = "三_3"
	PokerMap[18] = "三_3"
	PokerMap[19] = "三_3"
	PokerMap[20] = "三_3"
	PokerMap[21] = "3_3"
	PokerMap[22] = "3_3"
	PokerMap[23] = "3_3"
	PokerMap[24] = "3_3"
	PokerMap[25] = "四_4"
	PokerMap[26] = "四_4"
	PokerMap[27] = "四_4"
	PokerMap[28] = "四_4"
	PokerMap[29] = "4_4"
	PokerMap[30] = "4_4"
	PokerMap[31] = "4_4"
	PokerMap[32] = "4_4"
	PokerMap[33] = "五_5"
	PokerMap[33] = "五_5"
	PokerMap[33] = "五_5"
	PokerMap[33] = "五_5"
	PokerMap[33] = "5_5"
	PokerMap[33] = "5_5"
	PokerMap[33] = "5_5"
	PokerMap[33] = "5_5"
	PokerMap[33] = "六_6"
	PokerMap[33] = "六_6"
	PokerMap[33] = "六_6"
	PokerMap[33] = "六_6"
	PokerMap[33] = "六_6"
	PokerMap[33] = "6_6"
	PokerMap[33] = "6_6"
	PokerMap[33] = "6_6"
	PokerMap[33] = "6_6"
	PokerMap[33] = "七_7"
	PokerMap[33] = "七_7"
	PokerMap[33] = "七_7"
	PokerMap[33] = "七_7"
	PokerMap[33] = "7_7"
	PokerMap[33] = "7_7"
	PokerMap[33] = "7_7"
	PokerMap[33] = "7_7"
	PokerMap[33] = "八_8"
	PokerMap[33] = "八_8"
	PokerMap[33] = "八_8"
	PokerMap[33] = "八_8"
	PokerMap[33] = "8_8"
	PokerMap[33] = "8_8"
	PokerMap[33] = "8_8"
	PokerMap[33] = "8_8"
	PokerMap[33] = "九_9"
	PokerMap[33] = "九_9"
	PokerMap[33] = "九_9"
	PokerMap[33] = "九_9"
	PokerMap[33] = "9_9"
	PokerMap[33] = "9_9"
	PokerMap[33] = "9_9"
	PokerMap[33] = "9_9"
	PokerMap[33] = "十_10"
	PokerMap[33] = "十_10"
	PokerMap[33] = "十_10"
	PokerMap[33] = "十_10"
	PokerMap[33] = "10_10"
	PokerMap[33] = "10_10"
	PokerMap[33] = "10_10"
	PokerMap[33] = "10_10"
}

type PHZPoker struct {
	Id     int32 //每张牌唯一的ID
	Value  int32 //牌值
	Flower int32 //花色
}

func (p *PHZPoker) GetId() int32 {
	return p.Id
}

func (p *PHZPoker) GetValue() int32 {
	return p.Value
}

func (p *PHZPoker) GetFlower() int32 {
	return p.Flower
}
