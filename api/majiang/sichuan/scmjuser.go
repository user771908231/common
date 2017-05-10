package sichuan

import (
	"casino_common/api/majiang/api"
	"github.com/name5566/leaf/gate"
	"casino_common/common/sessionService"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"reflect"
	"casino_common/common/log"
	"casino_common/api/majiang"
	"errors"
	"casino_majiang/msg/protogo"
	"casino_common/utils/iputils"
	"casino_common/api/majiang/utils"
	"casino_majiang/msg/funcsInit"
)

//玩家的游戏状态
var MJUSER_STATUS_INTOROOM int32 = 1 ///刚进入游戏
var MJUSER_STATUS_SEATED int32 = 2   //坐下游戏
var MJUSER_STATUS_READY int32 = 3    //准备游戏
var MJUSER_STATUS_DINGQUE int32 = 4  //准备定缺
var MJUSER_STATUS_GAMING int32 = 5   //正在游戏，这里的正在游戏，表示还没有胡牌..
var MJUSER_STATUS_HUPAI int32 = 6    //已经胡牌了

var MJUER_APPLYDISSOLVE_S_REFUSE int32 = -1 //拒绝解散
var MJUER_APPLYDISSOLVE_S_DEFAULT int32 = 0 //没有处理
var MJUER_APPLYDISSOLVE_S_AGREE int32 = 1   //同意解散

//四川的麻将玩家，宜宾，彭州,都江堰
type SCMJUser struct {
	*api.MJUserCore                 //coreUser
	SCGameData      *SCMJUserData   //四川麻将的牌的信息
	SCGameStatus    *SCMJUserStatus //四川麻将的状态信息
	SCGameStatistic *UserStatistic  //统计信息
	SCGameBills     map[int32]*Bill //账单
}

func NewSCMJuser(userId uint32, a gate.Agent) *SCMJUser {
	return &SCMJUser{
		MJUserCore: api.NewMJUserCore(userId, a),
	}
}

func (d *SCMJUser) UpdateAgent(a gate.Agent) {
	d.Agent = a
	d.Break = false
	d.Leave = false
}

//得到和玩家关键的 desk
func (d *SCMJUser) GetDesk() api.MJDesk {
	return d.Desk
}

//更新玩家的sessoin信息
func (u *SCMJUser) UpdateSession(gid, gameNumber int32) error {
	//3,更新userSession,返回desk 的信息
	s, _ := sessionService.UpdateSession(
		u.GetUserId(),
		int32(ddproto.COMMON_ENUM_GAMESTATUS_GAMING),
		gid,
		gameNumber,
		u.GetDesk().GetRoom().GetRoomId(),
		u.GetDesk().GetDeskId(),
		u.S,
		u.Break,
		u.Leave,
		u.GetDesk().GetRoom().GetRoomType(),
		u.GetDesk().GetPassword())
	if s != nil {
		//给agent设置session
		agent := u.Agent
		if agent != nil {
			agent.SetUserData(s)
		}
	}
	return nil
}

//发送信息
func (u *SCMJUser) WriteMsg2(p proto.Message) error {
	if u == nil {
		return nil
	}

	if p == nil {
		return nil
	}
	//todo判断条件
	if u.Break || u.Leave {
		return nil
	}
	//
	log.T("%v开始给玩家[%v]发送type[%v]，msg[%v]", u.Desk.DlogDes(), u.GetUserId(), reflect.TypeOf(p).String(), p)
	u.Agent.WriteMsg(p)
	return nil
}

//todo 删除打出去的牌
func (u *SCMJUser) DelOutPai(p *majiang.MJPAI) error {
	index := -1
	for i, pai := range u.SCGameData.OutPais {
		if pai != nil && pai.Index == p.Index {
			index = i
			break
		}
	}

	//没有找到对应的手牌
	if index < 0 {
		log.E("%v服务器错误：删除打出去的时候出错，没有找到对应打出去[%v]", u.GetDesk().DlogDes(), p)
		return errors.New("删除手牌时出错，没有找到对应的out牌...")
	}
	//删除手牌
	u.SCGameData.OutPais = append(u.SCGameData.OutPais[:index], u.SCGameData.OutPais[index+1:]...)
	return nil
}

func (u *SCMJUser) GetGameData() interface{} {
	return u.SCGameData
}

func (u *SCMJUser) GetPlayerCard(showHand bool) *mjproto.PlayerCard {
	playerCard := newProto.NewPlayerCard()
	playerCard.UserId = proto.Uint32(u.GetUserId())
	playerCard.HandCardCount = proto.Int32(int32(len(u.SCGameData.GetHandPais()))) //手牌的长度

	//得到所有的手牌手牌
	if showHand {
		for _, pai := range u.SCGameData.GetHandPais() {
			if pai != nil {
				playerCard.HandCard = append(playerCard.HandCard, pai.GetCardInfo())
			}
		}
	}

	//类型（1,碰，2,明杠，3,暗杠）
	//得到碰牌
	for _, info := range u.SCGameData.GetPengPais() {
		if info != nil {
			com := newProto.NewComposeCard()
			*com.Value = info.Pais[0].GetClientId()
			*com.Type = int32(mjproto.ComposeCardType_C_PENG) //todo ,需要把type 放置在常量里面 这里代表的是碰牌
			playerCard.ComposeCard = append(playerCard.ComposeCard, com)
		}
	}

	//得到杠牌
	for _, info := range u.SCGameData.GetGangPais() {
		if info != nil {
			com := newProto.NewComposeCard()
			*com.Value = info.Pais[0].GetClientId()

			if info.GangType == api.GANGTYPE_MING {
				*com.Type = int32(mjproto.ComposeCardType_C_MINGGANG) // 明杠
			} else if info.GangType == api.GANGTYPE_AN {
				*com.Type = int32(mjproto.ComposeCardType_C_ANGANG) // 暗杠
			} else if info.GangType == api.GANGTYPE_BA {
				*com.Type = int32(mjproto.ComposeCardType_C_BAGANG) // 巴杠
			}
			playerCard.ComposeCard = append(playerCard.ComposeCard, com)
		}
	}

	//得到胡牌
	for _, hu := range u.SCGameData.GetHuPais() {
		if hu != nil {
			playerCard.HuCard = append(playerCard.HuCard, hu.Pai.GetCardInfo())
		}
	}

	//打出去的牌
	for _, pai := range u.SCGameData.GetOutPais() {
		if pai != nil {
			playerCard.OutCard = append(playerCard.OutCard, pai.GetCardInfo())
		}
	}

	return playerCard
}

//返回一个用户信息
func (u *SCMJUser) GetPlayerInfo(showHand bool, reconnect mjproto.RECONNECT_TYPE) *mjproto.PlayerInfo {
	log.T("开始得到user[%v]的牌的信息，showHand[%v]", u.GetUserId(), showHand)
	info := newProto.NewPlayerInfo()
	if u.Ready {
		info.BReady = proto.Int32(1)
	}
	*info.NHuPai = 0     //捉虾子用不着的
	*info.BDingQue = 0   //捉虾子用不着的
	*info.BExchanged = 0 //捉虾子用不着的
	*info.Coin = u.GetCoin()
	*info.IsBanker = u.Banker
	info.PlayerCard = u.GetPlayerCard(showHand) //获取游戏的数据
	*info.NickName = u.GetNickName()
	*info.UserId = u.GetUserId()
	*info.Sex = u.GetSex()
	info.Ip = proto.String(u.GetIp())
	info.Address = proto.String(iputils.GetIPSite(u.GetIp()))
	info.IsOwner = proto.Bool(u.IsOwner)
	info.WxInfo = &mjproto.WeixinInfo{
		NickName: proto.String(u.GetNickName()),
		HeadUrl:  proto.String(u.GetHeadUrl()),
		OpenId:   proto.String(u.GetOpenId()),
	}
	if reconnect == mjproto.RECONNECT_TYPE_RECONNECT {
		//如果是断线重连
	}
	//如果是断线重连，需要把手牌中的摸牌给去掉
	return info
}

//发送OverTurn
func (u *SCMJUser) SendOverTurn(p proto.Message) error {
	u.WriteMsg2(p)
	return nil
}

func (u *SCMJUser) AddHandPai(p *majiang.MJPAI) error {
	if p != nil {
		u.SCGameData.HandPais = append(u.SCGameData.HandPais, p)
	}
	return nil
}

//删除手牌
func (u *SCMJUser) DelHandPai(p *majiang.MJPAI) error {
	index := -1
	for i, pai := range u.SCGameData.HandPais {
		if pai != nil && pai.Index == p.Index {
			index = i
			break
		}
	}

	//没有找到对应的手牌
	if index < 0 {
		log.E("%v服务器错误：删除手牌的时候出错，没有找到对应的手牌[%v]", u.GetDesk().DlogDes(), p.LogDes())
		return errors.New("删除手牌时出错，没有找到对应的手牌...")
	}

	//删除手牌
	u.SCGameData.HandPais = append(u.SCGameData.HandPais[:index], u.SCGameData.HandPais[index+1:]...)
	return nil
}

//删除鹏的牌
func (u *SCMJUser) DelPengPai(p *majiang.PengPai) error {
	index := -1
	for i, pai := range u.SCGameData.PengPais {
		if pai != nil && pai.Pais[0].GetClientId() == p.Pais[0].GetClientId() {
			index = i
			break
		}
	}

	//没有找到对应的碰牌
	if index < 0 {
		log.E("%v服务器错误：删除碰牌的时候出错，没有找到对应的碰牌[%v]", u.GetDesk().DlogDes(), p)
		return errors.New("删除手牌时出错，没有找到对应的手牌...")
	}

	//删除手牌
	u.SCGameData.PengPais = append(u.SCGameData.PengPais[:index], u.SCGameData.PengPais[index+1:]...)
	return nil
}

//描述手牌的信息
func (u *SCMJUser) DescAllPais() string {
	return utils.ListGameData2Str(u.SCGameData)
}
