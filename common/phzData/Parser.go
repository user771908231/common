package phzData

import (
	"casino_common/common/log"
	"casino_common/proto/ddproto"
)

type PHZParser interface {
	CanTi(...interface{}) (interface{}, error)   //根据桌面的牌判断手牌里三张的牌是否可以跑
	GetTiPai(interface{}) interface{}            //获取手牌里原始可以提的牌
	CanPeng(...interface{}) (interface{}, error) //根据桌面的牌判断手牌里是否有可以碰的牌
	CanChi(...interface{}) (interface{}, error)  //根据桌面的牌判断手牌里是否有可以吃的牌型
	CanHu(...interface{}) (interface{}, error)   //根据桌面的牌判断是否可以胡牌
}

type ParserCore struct {
}

type HuInfo struct {
	WinUser    uint32 //
	HuXiCount  int    //赢家的虾子的数量
	LoseUser   uint32
	IsZimo     bool
	GetUserId  uint32
	SendUserId uint32
	HuDesc     string
	Fan        int32
	Score      int64
	Pai        *PHZPoker
	HuType     int32
	PaiType    []int32 //牌的类型
}

func (p *ParserCore) CountHandPais(pais []*PHZPoker) []int {
	counts := make([]int, 21) //牌值1——10:小字  11——20:大字
	for _, p := range pais {
		if p.GetPaiIndexByValue() == 0 {
			log.T("CountHandPais occur error...p.value==0")
			return nil
		}
		counts[p.GetPaiIndexByValue()]++
	}
	return counts
}

func (p *ParserCore) GetTiPai(data interface{}) interface{} {
	var tiInfo []*TiPai
	handPokers := data.([]*PHZPoker)
	paiCounts := p.CountHandPais(handPokers)
	if paiCounts != nil {
		for paiValue, paiCount := range paiCounts {
			if paiCount == 4 {
				ti := &TiPai{}
				ti.TiType = int32(ddproto.PhzEnumTipaiType_PHZ_TIPAITYPE_MING)
				//将提的牌放进提牌结构里
				for _, pai := range handPokers {
					if pai.GetValue() == int32(paiValue) {
						ti.Pais = append(ti.Pais, pai)
					}
				}
				tiInfo = append(tiInfo, ti)
			}
		}
		return tiInfo
	}
	return nil
}

func (p *ParserCore) CanTi(data interface{}, tiData interface{}) (interface{}, error) {
	ti := &TiPai{}
	handPokers := data.([]*PHZPoker)
	tiPai := tiData.(*PHZPoker)
	paiCounts := p.CountHandPais(handPokers)
	if paiCounts != nil {
		for paiValue, paiCount := range paiCounts {
			if paiCount == 3 && paiValue == int(tiPai.GetValue()) {
				ti.TiType = int32(ddproto.PhzEnumTipaiType_PHZ_TIPAITYPE_AN)
				for _, pai := range handPokers {
					if pai.GetValue() == int32(paiValue) {
						ti.Pais = append(ti.Pais, pai)
					}
				}
			}
		}
		return ti, nil
	}
	return nil, nil
}

func (p *ParserCore) CanPeng(userGameData interface{}, pengData interface{}) (interface{}, error) {
	handPokers := userGameData.([]*PHZPoker)
	pengPai := pengData.(*PHZPoker)
	pengReslut := &PengPai{}
	paiCounts := p.CountHandPais(handPokers)
	if paiCounts != nil {
		for paiValue, paiCount := range paiCounts {
			if paiValue == int(pengPai.GetValue()) && paiCount == 2 {
				for _, pai := range handPokers {
					if pai.GetValue() == int32(paiValue) {
						pengReslut.Pais = append(pengReslut.Pais, pai)
					}
				}
			}
		}
		return pengReslut, nil
	}
	return nil, nil
}

func (p *ParserCore) CanChi(userGameData interface{}, chiData interface{}) (interface{}, error) {
	/*handPokers := userGameData.([]*PHZPoker)
	chiPai := chiData.(*PHZPoker)
	chiReslut := &ChiPai{}*/
	//todo
	return nil, nil
}
