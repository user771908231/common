package sichuan

import (
	"casino_common/api/majiang/api"
	"casino_common/api/majiang"
)

var mjpaiMap map[int]string

func init() {
	mjpaiMap = make(map[int]string, 108)
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
}

//四川的麻将解析器
type SCMJParser struct {
	*api.MJParserCore //麻将的骨架
}

func (*SCMJParser) CanChi(interface{}, interface{}) (interface{}, error) {
	panic("implement me")
}

func (*SCMJParser) CanTing(...interface{}) (interface{}, error) {
	panic("implement me")
}

func (*SCMJParser) GetJiaoInfos(...interface{}) (interface{}, error) {
	panic("implement me")
}

func (*SCMJParser) Parse(pids []int32) (interface{}, error) {
	panic("implement me")
}

func (*SCMJParser) XiPai() interface{} {
	panic("implement me")
}

func (*SCMJParser) Hu(...interface{}) (interface{}, error) {
	panic("implement me")
}

//是否能补虾子
func (p *SCMJParser) CanBu(...interface{}) (interface{}, error) {
	return nil, nil
}

//是否可以提
func (p *SCMJParser) CanTi(...interface{}) (interface{}, error) {
	return nil, nil
}

//是否可以飞
func (p *SCMJParser) CanFly(...interface{}) (bool, error) {
	return false, nil
}

//初始化牌
func (p *SCMJParser) InitMjPaiByIndex(index int32) *majiang.MJPAI {
	result := &majiang.MJPAI{
		Index: index,
		Des:   mjpaiMap[int(index)],
	}
	result.InitByDes()
	//返回结果
	return result
}

func (p *SCMJParser) CanGang(args ...interface{}) (interface{}, error) {
	//参数
	gameData := args[0].(api.MJUserGameData) //游戏数据
	var outPai *majiang.MJPAI
	if args[1] != nil {
		outPai = args[1].(*majiang.MJPAI) //别人打的牌
	}

	remainCount := args[2].(int32) //剩余牌的数量
	gangUserId := args[3].(uint32)

	//如果剩余的牌为0，那么直接不能杠
	if remainCount == 0 {
		return &api.CanGangInfo{
			CanGang: false,
		}, nil
	}

	//返回判定杠的结果
	var retCore interface{}
	var errCore error
	if outPai == nil {
		//表示自己摸牌的情况，这个时候可以巴杠
		retCore, errCore = p.MJParserCore.ZiMoGangCards(gameData)
	} else {
		//点杠
		retCore, errCore = p.MJParserCore.CanGang(gameData, outPai)
	}

	var ret *api.CanGangInfo = &api.CanGangInfo{
		CanGang:     false,
		GangsUserId: gangUserId,
	}

	if errCore != nil {
		return ret, errCore
	}

	if retCore == nil {
		return ret, errCore

	}

	//组装杠牌的代码
	if r := retCore.(*api.CanGangInfo); r.CanGang {
		for _, b := range r.GangInfoBean {
			if b != nil {
				ret.CanGang = true
				ret.GangInfoBean = append(ret.GangInfoBean, b)
				ret.GangsUserId = r.GangsUserId
				ret.OutUserId = r.OutUserId
			}
		}
	}

	return ret, errCore
}

//todo 是否能碰牌
func (p *SCMJParser) CanPeng(gameData interface{}, outPai interface{}) (bool, error) {
	if outPai == nil {
		return false, nil
	}

	return p.MJParserCore.CanPeng(gameData, outPai)
}
