package majiang

import (
	"fmt"
	"sort"
	"testing"
)

var TOTALPAICOUNT int = 136
var mjpaiMap map[int]string

func init() {
	mjpaiMap = make(map[int]string, TOTALPAICOUNT)
	mjpaiMap[0] = "T_1"
	mjpaiMap[1] = "T_1"
	mjpaiMap[2] = "T_1"
	mjpaiMap[3] = "T_1"
	mjpaiMap[4] = "T_2"
	mjpaiMap[5] = "T_2"
	mjpaiMap[6] = "T_2"
	mjpaiMap[7] = "T_2"
	mjpaiMap[8] = "T_3"
	mjpaiMap[9] = "T_3"
	mjpaiMap[10] = "T_3"
	mjpaiMap[11] = "T_3"
	mjpaiMap[12] = "T_4"
	mjpaiMap[13] = "T_4"
	mjpaiMap[14] = "T_4"
	mjpaiMap[15] = "T_4"
	mjpaiMap[16] = "T_5"
	mjpaiMap[17] = "T_5"
	mjpaiMap[18] = "T_5"
	mjpaiMap[19] = "T_5"
	mjpaiMap[20] = "T_6"
	mjpaiMap[21] = "T_6"
	mjpaiMap[22] = "T_6"
	mjpaiMap[23] = "T_6"
	mjpaiMap[24] = "T_7"
	mjpaiMap[25] = "T_7"
	mjpaiMap[26] = "T_7"
	mjpaiMap[27] = "T_7"
	mjpaiMap[28] = "T_8"
	mjpaiMap[29] = "T_8"
	mjpaiMap[30] = "T_8"
	mjpaiMap[31] = "T_8"
	mjpaiMap[32] = "T_9"
	mjpaiMap[33] = "T_9"
	mjpaiMap[34] = "T_9"
	mjpaiMap[35] = "T_9"
	mjpaiMap[36] = "S_1"
	mjpaiMap[37] = "S_1"
	mjpaiMap[38] = "S_1"
	mjpaiMap[39] = "S_1"
	mjpaiMap[40] = "S_2"
	mjpaiMap[41] = "S_2"
	mjpaiMap[42] = "S_2"
	mjpaiMap[43] = "S_2"
	mjpaiMap[44] = "S_3"
	mjpaiMap[45] = "S_3"
	mjpaiMap[46] = "S_3"
	mjpaiMap[47] = "S_3"
	mjpaiMap[48] = "S_4"
	mjpaiMap[49] = "S_4"
	mjpaiMap[50] = "S_4"
	mjpaiMap[51] = "S_4"
	mjpaiMap[52] = "S_5"
	mjpaiMap[53] = "S_5"
	mjpaiMap[54] = "S_5"
	mjpaiMap[55] = "S_5"
	mjpaiMap[56] = "S_6"
	mjpaiMap[57] = "S_6"
	mjpaiMap[58] = "S_6"
	mjpaiMap[59] = "S_6"
	mjpaiMap[60] = "S_7"
	mjpaiMap[61] = "S_7"
	mjpaiMap[62] = "S_7"
	mjpaiMap[63] = "S_7"
	mjpaiMap[64] = "S_8"
	mjpaiMap[65] = "S_8"
	mjpaiMap[66] = "S_8"
	mjpaiMap[67] = "S_8"
	mjpaiMap[68] = "S_9"
	mjpaiMap[69] = "S_9"
	mjpaiMap[70] = "S_9"
	mjpaiMap[71] = "S_9"
	mjpaiMap[72] = "W_1"
	mjpaiMap[73] = "W_1"
	mjpaiMap[74] = "W_1"
	mjpaiMap[75] = "W_1"
	mjpaiMap[76] = "W_2"
	mjpaiMap[77] = "W_2"
	mjpaiMap[78] = "W_2"
	mjpaiMap[79] = "W_2"
	mjpaiMap[80] = "W_3"
	mjpaiMap[81] = "W_3"
	mjpaiMap[82] = "W_3"
	mjpaiMap[83] = "W_3"
	mjpaiMap[84] = "W_4"
	mjpaiMap[85] = "W_4"
	mjpaiMap[86] = "W_4"
	mjpaiMap[87] = "W_4"
	mjpaiMap[88] = "W_5"
	mjpaiMap[89] = "W_5"
	mjpaiMap[90] = "W_5"
	mjpaiMap[91] = "W_5"
	mjpaiMap[92] = "W_6"
	mjpaiMap[93] = "W_6"
	mjpaiMap[94] = "W_6"
	mjpaiMap[95] = "W_6"
	mjpaiMap[96] = "W_7"
	mjpaiMap[97] = "W_7"
	mjpaiMap[98] = "W_7"
	mjpaiMap[99] = "W_7"
	mjpaiMap[100] = "W_8"
	mjpaiMap[101] = "W_8"
	mjpaiMap[102] = "W_8"
	mjpaiMap[103] = "W_8"
	mjpaiMap[104] = "W_9"
	mjpaiMap[105] = "W_9"
	mjpaiMap[106] = "W_9"
	mjpaiMap[107] = "W_9"

	//东
	mjpaiMap[108] = "FENG_1_EAST"
	mjpaiMap[109] = "FENG_1_EAST"
	mjpaiMap[110] = "FENG_1_EAST"
	mjpaiMap[111] = "FENG_1_EAST"

	//南
	mjpaiMap[112] = "FENG_2_SOUTH"
	mjpaiMap[113] = "FENG_2_SOUTH"
	mjpaiMap[114] = "FENG_2_SOUTH"
	mjpaiMap[115] = "FENG_2_SOUTH"

	//西
	mjpaiMap[116] = "FENG_3_WEST"
	mjpaiMap[117] = "FENG_3_WEST"
	mjpaiMap[118] = "FENG_3_WEST"
	mjpaiMap[119] = "FENG_3_WEST"

	//北
	mjpaiMap[120] = "FENG_4_NORTH"
	mjpaiMap[121] = "FENG_4_NORTH"
	mjpaiMap[122] = "FENG_4_NORTH"
	mjpaiMap[123] = "FENG_4_NORTH"

	//中
	mjpaiMap[124] = "FENG_5_Z"
	mjpaiMap[125] = "FENG_5_Z"
	mjpaiMap[126] = "FENG_5_Z"
	mjpaiMap[127] = "FENG_5_Z"

	//白
	mjpaiMap[128] = "FENG_6_B"
	mjpaiMap[129] = "FENG_6_B"
	mjpaiMap[130] = "FENG_6_B"
	mjpaiMap[131] = "FENG_6_B"

	//发
	mjpaiMap[132] = "FENG_7_F"
	mjpaiMap[133] = "FENG_7_F"
	mjpaiMap[134] = "FENG_7_F"
	mjpaiMap[135] = "FENG_7_F"
}

func TestSort(t *testing.T) {
	ids := []int32{1, 47, 3, 42, 5, 2, 6, 20, 4, 8, 10}
	pais := []*MJPAI{}
	for _, id := range ids {
		pais = append(pais, InitMjPaiByIndex(id))
	}
	list := MjPAIList{}
	list = pais
	println(fmt.Sprintf("1 %v", List2Str(list)))
	sort.Sort(list)
	println(fmt.Sprintf("2 %v", List2Str(list)))
}

func InitMjPaiByIndex(index int32) *MJPAI {
	result := &MJPAI{
		Index: index,
		Des:   mjpaiMap[int(index)],
	}
	result.InitByDes()
	//返回结果
	return result
}

func List2Str(pais []*MJPAI) string {
	if len(pais) <= 0 {
		return ""
	}
	s := ""
	for _, p := range pais {
		s = s + p.LogDes() + " "
	}
	//返回值
	return s
}
