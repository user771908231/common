package ddproto

import (
	"errors"
	"sort"
	"casino_common/common/log"
	"casino_common/utils/pokerUtils"
	"casino_common/utils/numUtils"
)

//排序的问题
type DdzPokerOutList []*CommonSrvPokerPai;

func (list DdzPokerOutList) Less(i, j  int) bool {
	if list[i].GetValue() < list[j].GetValue() {
		return true
	} else {
		return false
	}
}

// Len 为集合内元素的总数
func (list DdzPokerOutList)Len() int {
	return len(list)
}

// Swap 交换索引为 i 和 j 的元素
func (list DdzPokerOutList) Swap(i, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}


//通过牌来初始化其他的参数
func (out *DdzSrvOutPokerPais) Init() error {
	//1,先判断是否有牌
	if out.PokerPais == nil || len(out.PokerPais) <= 0 {
		return errors.New("初始化出的牌失败...没有找到牌。")
	}

	//2，通过牌来进行初始化
	out.sortPais()        //首先对牌进行排序
	out.initTypeAndKeyValue()        //初始化类型和比较值
	return nil
}

//得到牌的张数
func (out *DdzSrvOutPokerPais) getPaiCount() int32 {
	return int32(len(out.PokerPais))
}

//对牌进行排序,左边小，右边大的值进行排序
func (out *DdzSrvOutPokerPais) sortPais() error {
	var list DdzPokerOutList = out.PokerPais
	sort.Sort(list)        //进行排序
	return nil
}


//初始化类型
func (out *DdzSrvOutPokerPais) initTypeAndKeyValue() error {
	//统计数据
	counts := make([]int32, 18) //value作索引 索引最高17 所以需要18个空间

	//统计牌的数量
	for _, pai := range out.PokerPais {
		counts[pai.GetValue()]++ //把牌的value作为索引
	}

	var liangzhangs []int32
	var sanzhangs []int32
	var sizhangs []int32

	for paiValue, paiCount := range counts {

		value := int32(paiValue)

		if paiCount == 1 {
			*out.CountYizhang ++
		}

		if paiCount == 2 {
			*out.CountDuizi ++
			liangzhangs = append(liangzhangs, value)
			//liangzhangs[paiValue] = paiCount
		}

		if paiCount == 3 {
			*out.CountSanzhang ++
			sanzhangs = append(sanzhangs, value)
			//sanzhangs[paiValue] = paiCount
		}

		if paiCount == 4 {
			*out.CountSizhang ++
			sizhangs = append(sizhangs, value)
			//sizhangs[paiValue] = paiCount
		}
	}

	/*** 根据count判断牌的基本类型 ***/

	//判断是否是单张
	isDanZhang := false
	if out.getPaiCount() == 1 {
		isDanZhang = true
	}

	isDuiZi := false
	if out.getPaiCount() == 2 && out.GetCountDuizi() == 1 {
		isDuiZi = true
	}

	isWangZha := false
	if out.getPaiCount() == 2 && counts[16] == 1 && counts[17] == 1 {
		isWangZha = true
	}

	isSanZhang := false
	if out.getPaiCount() == 3 && out.GetCountSanzhang() == 1 {
		isSanZhang = true
	}

	isSanDaiYi := false
	if out.getPaiCount() == 4 && out.GetCountSanzhang() == 1 && out.GetCountYizhang() == 1 {
		isSanDaiYi = true
	}

	isSiZhang := false
	if out.getPaiCount() == 4 && out.GetCountSizhang() == 1 {
		isSiZhang = true
	}

	isSanDaiDui := false
	if out.getPaiCount() == 5 && out.GetCountSanzhang() == 1 && out.GetCountDuizi() == 1 {
		isSanDaiDui = true
		if counts[16] > 0 || counts[17] > 0 {
			//三带对 不能带对王
			isSanDaiDui = false
		}
	}


	//判断是否是单顺 5张或以上
	isDanShun := false
	if (out.GetCountYizhang() == out.getPaiCount()) && (out.GetCountYizhang() >= 5) {
		//单张的数量等于出牌的数量 即牌型的每个值都是独立的
		boolFlag := true
		for k := 0; k < int(out.getPaiCount() - 1); k++ {
			if out.PokerPais[k].GetValue() + 1 != out.PokerPais[k + 1].GetValue() {
				boolFlag = false
				break
			}
		}
		if counts[15] > 0 || counts[16] > 0 || counts[17] > 0 {
			//不能包含2 blackjoker redjoker
			boolFlag = false
		}
		isDanShun = boolFlag
	}

	//判断是否是双顺 三对或以上
	isLianDui := false
	if (out.GetCountDuizi() * 2 == out.getPaiCount())  && (out.GetCountDuizi() >= 3) {
		boolFlag := true
		for i := 0; i < int(out.GetCountDuizi() - 1); i++ {
			if liangzhangs[i] + 1 != liangzhangs[i + 1] {
				boolFlag = false
				break
			}
		}
		if counts[15] > 0 || counts[16] > 0 || counts[17] > 0 {
			//不能包含2 blackjoker redjoker
			boolFlag = false
		}
		isLianDui = boolFlag
	}

	//飞机
	isFeiji := false
	if (out.GetCountSanzhang() * 3 == out.getPaiCount()) && out.GetCountSanzhang() >= 2 {
		isFeiji = true
		if counts[15] > 0 {
			isFeiji = false //飞机不能包含2
		}
	}

	//飞机带单张
	isFeijiChibang := false
	if out.GetCountSanzhang() * 4 == out.getPaiCount() && out.GetCountSanzhang() >= 2 {
		boolFlag := true
		for i := 0; i < int(out.GetCountSanzhang()) - 1; i++ {
			if sanzhangs[i] + 1 != sanzhangs[i + 1] {
				boolFlag = false
				break
			}
			if sanzhangs[i] == 15 || sanzhangs[i + 1] == 15 {
				//三张里不能包含2
				boolFlag = false
				break
			}
		}
		isFeijiChibang = boolFlag
	}

	//飞机带对子
	isFeijiDuizi := false
	if out.GetCountSanzhang() * 5 == out.getPaiCount() && out.GetCountSanzhang() == out.GetCountDuizi() && out.GetCountSanzhang() >= 2 {
		boolFlag := true
		for i := 0; i < int(out.GetCountSanzhang() - 1); i++ {
			if sanzhangs[i] + 1 != sanzhangs[i + 1] {
				boolFlag = false
				break
			}
			if sanzhangs[i] == 15 || sanzhangs[i + 1] == 15 {
				//三张里不能包含2
				boolFlag = false
				break
			}
		}
		isFeijiDuizi = boolFlag
	}

	isSiDaiLiangDui := false
	if out.getPaiCount() == 8 && out.GetCountSizhang() == 1 && out.GetCountDuizi() == 2 {
		isSiDaiLiangDui = true        //四个带两队
		if counts[16] > 0 || counts[17] > 0 {
			//不能包含王
			isSiDaiLiangDui = true
		}
	}

	isSiDaiYiDui := false
	if out.getPaiCount() == 6 && out.GetCountSizhang() == 1 && out.GetCountDuizi() == 1 {
		isSiDaiYiDui = true        //四个带一对
		if counts[16] > 0 || counts[17] > 0 {
			//不能包含王
			isSiDaiYiDui = true
		}
	}

	isSiDaiLiangZhang := false
	if out.getPaiCount() == 6 && out.GetCountSizhang() == 1 && out.GetCountYizhang() == 2 {
		isSiDaiLiangZhang = true        //四个带两张
	}



	/*** 设置牌型type、关键牌值keyValue ***/

	//switch 牌型独立 没有交集
	switch {
	case isDanZhang : //单张
		*out.Type = int32(DdzEnumPaiType_SINGLECARD) //单张
		*out.KeyValue = out.GetPokerPais()[0].GetValue()
		log.T("user[%v] 出牌类型为单张, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isDuiZi : //对子
		*out.Type = int32(DdzEnumPaiType_DOUBLECARD) // 对子
		*out.KeyValue = out.GetPokerPais()[0].GetValue() //or liangzhangs[0]
		log.T("user[%v] 出牌类型为对子, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isWangZha : //王炸
		*out.Type = int32(DdzEnumPaiType_SUPERBOMB)
		*out.KeyValue = out.GetPokerPais()[0].GetValue()
		*out.IsBomb = true
		log.T("user[%v] 出牌类型为王炸, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isDanShun : //单顺
		*out.Type = int32(DdzEnumPaiType_CONNECTCARD)        // 顺子
		*out.KeyValue = out.GetPokerPais()[0].GetValue()
		log.T("user[%v] 出牌类型为顺子, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isLianDui : //双顺
		*out.Type = int32(DdzEnumPaiType_COMPANYCARD) //连队
		*out.KeyValue = out.GetPokerPais()[0].GetValue()
		log.T("user[%v] 出牌类型为双顺, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isSanZhang : //三张
		*out.Type = int32(DdzEnumPaiType_THREECARD)                //三张
		*out.KeyValue = out.GetPokerPais()[0].GetValue() //or sanzhangs[0]
		log.T("user[%v] 出牌类型为三张, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isSanDaiYi : //三带一
		*out.Type = int32(DdzEnumPaiType_THREEONECARD)        // 三带一
		*out.KeyValue = sanzhangs[0]
		log.T("user[%v] 出牌类型为三带一, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isSiZhang : //四张
		*out.Type = int32(DdzEnumPaiType_BOMBCARD)
		*out.KeyValue = out.GetPokerPais()[0].GetValue() //or sizhangs[0]
		*out.IsBomb = true
		log.T("user[%v] 出牌类型为四张, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isSiDaiLiangDui : //四带两队
		*out.Type = int32(DdzEnumPaiType_BOMBTWOOOCARD)       // 四带两队
		*out.KeyValue = sizhangs[0]
		log.T("user[%v] 出牌类型为四带两队, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isSiDaiYiDui || isSiDaiLiangZhang : //四带两张
		*out.Type = int32(DdzEnumPaiType_BOMBTWOCARD)      // 四带二个单张
		*out.KeyValue = sizhangs[0]
		log.T("user[%v] 出牌类型为四带两张, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isSanDaiDui : //三带对
		*out.Type = int32(DdzEnumPaiType_THREETWOCARD)
		*out.KeyValue = sanzhangs[0]
		log.T("user[%v] 出牌类型为三带对, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isFeiji : //飞机
		*out.Type = int32(DdzEnumPaiType_AIRCRAFTCARD)
		*out.KeyValue = sanzhangs[0]
		log.T("user[%v] 出牌类型为飞机, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	case isFeijiChibang || isFeijiDuizi: //飞机带翅膀
		*out.Type = int32(DdzEnumPaiType_AIRCRAFTSINGLECARD)
		*out.KeyValue = sanzhangs[0]
		log.T("user[%v] 出牌类型为飞机带翅膀, keyValue为[%v]", out.GetUserId(), *out.KeyValue)

	default:
		*out.Type = int32(DdzEnumPaiType_ERRORCARD)   //默认错误牌型
		*out.KeyValue = 0
		log.T("user[%v] 出牌类型错误", out.GetUserId())
	}
	return nil;
}

//比较两幅牌的大小
func (out *DdzSrvOutPokerPais)  GT(outb *DdzSrvOutPokerPais) (bool, error) {
	if out.GetType() == outb.GetType() {
		return out.GetKeyValue() > outb.GetKeyValue(), nil
	} else {
		//比较类型不同的情况
		if out.GetIsBomb() || outb.GetIsBomb() {
			return out.GetIsBomb(), nil
		} else {
			log.E("牌型[%v]和牌型[%v] 无法比较...", out.GetType(), outb.GetType())
			return false, errors.New("无法比较，类型有错误..")
		}
	}

}

func (out *DdzSrvOutPokerPais) GetClientPokers() []*ClientBasePoker {
	//return ServerPoker2ClienPoker(out.PokerPais)
	return nil
}


//牌相关的方法

func ( p *CommonSrvPokerPai) GetClientPoker() *ClientBasePoker {
	ret := new(ClientBasePoker)
	ret.Suit = new(CommonEnumPokerColor)
	ret.Id = p.Id
	ret.Num = p.Value
	*ret.Suit = p.GetSuit()
	return ret
}

func (p *CommonSrvPokerPai) GetSuit() CommonEnumPokerColor {
	if p.GetFlower() == pokerUtil.FLOWER_DIAMOND {
		return CommonEnumPokerColor_FANGKUAI
	} else if p.GetFlower() == pokerUtil.FLOWER_CLUB {
		return CommonEnumPokerColor_MEIHUA
	} else if p.GetFlower() == pokerUtil.FLOWER_HEART {
		return CommonEnumPokerColor_HONGTAO
	} else if p.GetFlower() == pokerUtil.FLOWER_SPADE {
		return CommonEnumPokerColor_HEITAO
	} else if p.GetFlower() == pokerUtil.FLOWER_BLACKJOKER {
		return CommonEnumPokerColor_BLACKBIGJOKER
	} else {
		return CommonEnumPokerColor_REDJOKER
	}
}

func (p *CommonSrvPokerPai) GetLogDes() string {
	suit, _ := numUtils.Int2String(p.GetId())
	switch p.GetSuit() {
	case CommonEnumPokerColor_FANGKUAI:
		suit = "方块"
	case CommonEnumPokerColor_BLACKBIGJOKER:
		suit = "小王"
	case CommonEnumPokerColor_HEITAO:
		suit = "黑桃"
	case CommonEnumPokerColor_HONGTAO:
		suit = "红桃"
	case CommonEnumPokerColor_REDJOKER:
		suit = "大王"
	case CommonEnumPokerColor_MEIHUA:
		suit = "梅花"
	default:
	}
	suit += p.GetName()
	return suit
}

