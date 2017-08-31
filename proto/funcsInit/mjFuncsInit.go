package commonNewPorot

import (
	"casino_common/proto/ddproto"
)

func SuccessHeader() *ddproto.ProtoHeader {
	header := NewHeader()
	*header.Code = 0
	return header
}

func ErrorHeader() *ddproto.ProtoHeader {
	header := NewHeader()
	*header.Code = -1
	return header
}

func NewGame_AckCreateRoom() *ddproto.Game_AckCreateRoom {
	ret := &ddproto.Game_AckCreateRoom{}
	ret.Header = NewHeader()
	ret.DeskId = new(int32)
	ret.Password = new(string)
	ret.UserBalance = new(int64)
	ret.CreateFee = new(int64)
	return ret
}

func NewGame_AckReady() *ddproto.Game_AckReady {
	ret := &ddproto.Game_AckReady{}
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	return ret
}

func NewGame_AckEnterRoom() *ddproto.Game_AckEnterRoom {
	ret := &ddproto.Game_AckEnterRoom{}
	ret.Header = NewHeader()
	return ret
}

func NewGame_BroadcastBeginDingQue() *ddproto.Game_BroadcastBeginDingQue {
	ret := &ddproto.Game_BroadcastBeginDingQue{}
	ret.Reconnect = new(bool)
	return ret
}

func NewGame_DingQue() *ddproto.Game_DingQue {
	que := &ddproto.Game_DingQue{}
	que.Header = NewHeader()
	que.Color = new(int32)
	que.UserId = new(uint32)
	return que
}

func NewRoomTypeInfo() *ddproto.RoomTypeInfo {
	ret := &ddproto.RoomTypeInfo{}
	ret.BaseValue = new(int64)
	ret.BoardsCout = new(int32)
	ret.CapMax = new(int64)
	ret.CardsNum = new(int32)
	ret.MjRoomType = new(ddproto.MJRoomType)
	ret.PlayOptions = NewPlayOptions()
	ret.Settlement = new(int32)
	return ret
}

func NewPlayOptions() *ddproto.PlayOptions {
	ret := &ddproto.PlayOptions{}
	ret.DianGangHuaRadio = new(int32)
	ret.HuRadio = new(int32)
	ret.ZiMoRadio = new(int32)
	return ret
}

func NewGame_SendGameInfo() *ddproto.Game_SendGameInfo {
	ret := &ddproto.Game_SendGameInfo{}
	ret.Header = NewHeader()
	ret.SenderUserId = new(uint32)
	ret.IsReconnect = new(ddproto.RECONNECT_TYPE)
	return ret
}

//返回deskGameInfo
func NewDeskGameInfo() *ddproto.MJDeskGameInfo {
	ret := &ddproto.MJDeskGameInfo{}
	ret.GameStatus = new(int32)
	ret.PlayerNum = new(int32)
	ret.ActiveUserId = new(uint32)
	ret.ActionTime = new(int32)
	ret.DelayTime = new(int32)
	ret.NInitActionTime = new(int32)
	ret.NInitDelayTime = new(int32)
	ret.InitRoomCoin = new(int64)
	ret.CurrPlayCount = new(int32)
	ret.TotalPlayCount = new(int32)
	ret.RoomNumber = new(string)
	ret.Banker = new(uint32)
	ret.RemainCards = new(int32)
	return ret
}

//返回一个playerInfo
func NewPlayerInfo() *ddproto.MJPlayerInfo {
	info := &ddproto.MJPlayerInfo{}
	info.IsBanker = new(bool)
	info.Coin = new(int64)
	info.NickName = new(string)
	info.UserId = new(uint32)
	info.IsOwner = new(bool)
	info.BReady = new(int32)
	info.BDingQue = new(int32)
	info.BExchanged = new(int32)
	info.NHuPai = new(int32)
	info.NickName = new(string)
	info.QuePai = new(int32)
	info.Sex = new(int32)
	return info
}

//麻将card
func NewPlayerCard() *ddproto.PlayerCard {
	card := &ddproto.PlayerCard{}
	card.UserId = new(uint32)
	return card
}

func NewCardInfo() *ddproto.CardInfo {
	cardInfo := &ddproto.CardInfo{}
	cardInfo.Id = new(int32)
	cardInfo.Type = new(int32)
	cardInfo.Value = new(int32)
	return cardInfo
}

func NewComposeCard() *ddproto.ComposeCard {
	ret := &ddproto.ComposeCard{}
	ret.Type = new(int32)
	ret.Value = new(int32)
	return ret
}

func NewGame_OverTurn() *ddproto.Game_OverTurn {
	ret := &ddproto.Game_OverTurn{}
	ret.ActType = new(int32)
	ret.CanGang = new(bool)
	ret.CanPeng = new(bool)
	ret.CanHu = new(bool)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	ret.NextUserId = new(uint32)
	ret.Time = new(int32)
	ret.PaiCount = new(int32)
	ret.CanGuo = new(bool)
	ret.CanBu = new(bool)
	ret.CanChi = new(bool)
	ret.CanTing = new(bool)
	ret.IsBaoTingAgent = new(bool)
	ret.CanFly = new(bool)
	ret.CanTi = new(bool)
	return ret
}

func NewSJHGame_OverTurn() *ddproto.Game_OverTurn {
	ret := &ddproto.Game_OverTurn{}
	ret.ActType = new(int32)
	ret.CanGang = new(bool)
	ret.CanPeng = new(bool)
	ret.CanHu = new(bool)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	ret.NextUserId = new(uint32)
	ret.Time = new(int32)
	ret.PaiCount = new(int32)
	ret.CanGuo = new(bool)
	ret.CanChi = new(bool)
	return ret
}

//转转麻将overturn
func NewMJZhZhGame_OverTurn() *ddproto.Game_OverTurn {
	ret := &ddproto.Game_OverTurn{}
	ret.ActType = new(int32)
	ret.CanGang = new(bool)
	ret.CanPeng = new(bool)
	ret.CanHu = new(bool)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	ret.NextUserId = new(uint32)
	ret.Time = new(int32)
	ret.PaiCount = new(int32)
	ret.CanGuo = new(bool)
	return ret
}

func NewGame_DealCards() *ddproto.Game_DealCards {
	ret := &ddproto.Game_DealCards{}
	ret.Header = NewHeader()
	ret.DealerUserId = new(uint32)
	return ret
}

func NewGame_Opening() *ddproto.Game_Opening {
	ret := &ddproto.Game_Opening{}
	ret.Header = NewHeader()
	ret.CurrPlayCount = new(int32)
	ret.Dice1 = new(int32)
	ret.Dice2 = new(int32)
	return ret
}

func NewGame_AckQuickConn() *ddproto.Game_AckQuickConn {
	ret := &ddproto.Game_AckQuickConn{}
	ret.CurrVersion = new(int32)
	ret.DownloadUrl = new(string)
	ret.Header = NewHeader()
	ret.IsMaintain = new(int32)
	ret.IsUpdate = new(int32)
	ret.ReleaseTag = new(int32)
	ret.VersionInfo = new(string)
	return ret
}

func NewGame_DingQueEnd() *ddproto.Game_DingQueEnd {
	ret := &ddproto.Game_DingQueEnd{}
	return ret
}

func NewDingQueEndBean() *ddproto.DingQueEndBean {
	ret := &ddproto.DingQueEndBean{}
	ret.UserId = new(uint32)
	ret.Flower = new(int32)
	return ret
}

func NewGame_AckActHu() *ddproto.Game_AckActHu {
	ret := &ddproto.Game_AckActHu{}
	ret.Header = NewHeader()
	ret.UserIdIn = new(uint32)
	ret.UserIdOut = new(uint32)
	ret.HuType = new(int32)
	ret.IsZiMo = new(bool)
	return ret
}

func NewGame_AckSendOutCard() *ddproto.Game_AckSendOutCard {
	ret := &ddproto.Game_AckSendOutCard{}
	ret.Header = NewHeader()
	ret.Result = new(int32)
	ret.UserId = new(uint32)
	return ret
}

func NewGame_AckActGang() *ddproto.Game_AckActGang {
	ret := &ddproto.Game_AckActGang{}
	ret.GangType = new(int32)
	ret.Header = NewHeader()
	ret.UserIdIn = new(uint32)
	ret.UserIdOut = new(uint32)
	return ret
}

func NewGame_AckActPeng() *ddproto.Game_AckActPeng {
	ret := &ddproto.Game_AckActPeng{}
	ret.Header = NewHeader()
	ret.UserIdIn = new(uint32)
	ret.UserIdOut = new(uint32)
	ret.JiaoInfos = nil
	return ret
}

func NewGame_AckActChi() *ddproto.Game_AckActChi {
	ret := &ddproto.Game_AckActChi{}
	ret.Header = NewHeader()
	ret.UserIdIn = new(uint32)
	ret.UserIdOut = new(uint32)
	ret.JiaoInfos = nil
	return ret
}

func NewGame_AckLogin() *ddproto.Game_AckLogin {
	ret := &ddproto.Game_AckLogin{}
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	ret.NickName = new(string)
	ret.RoomPassword = new(string)
	ret.CostCreateRoom = new(int64)
	ret.CostRebuy = new(int64)
	ret.Championship = new(bool)
	ret.Chip = new(int64)
	ret.MailCount = new(int32)
	ret.Notice = new(string)
	ret.GameStatus = new(int32)
	return ret
}

func NewGame_SendCurrentResult() *ddproto.Game_SendCurrentResult {
	ret := &ddproto.Game_SendCurrentResult{}
	ret.Header = NewHeader()
	return ret

}

func NewGame_SendEndLottery() *ddproto.Game_SendEndLottery {
	ret := &ddproto.Game_SendEndLottery{}
	ret.Header = NewHeader()
	return ret
}

func NewWinCoinInfo() *ddproto.WinCoinInfo {
	ret := &ddproto.WinCoinInfo{}
	ret.CardTitle = new(string)
	ret.Coin = new(int64)
	ret.HuCount = new(int32)
	ret.IsDealer = new(bool)
	ret.UserId = new(uint32)
	ret.NickName = new(string)
	ret.WinCoin = new(int64)
	return ret
}

func NewEndLotteryInfo() *ddproto.EndLotteryInfo {
	ret := &ddproto.EndLotteryInfo{}
	ret.UserId = new(uint32)
	ret.NickName = new(string)
	ret.BigWin = new(bool)
	ret.BestGunner = new(bool)
	ret.IsOwner = new(bool)
	ret.WinCoin = new(int64)
	ret.CountHu = new(int32)
	ret.CountZiMo = new(int32)
	ret.CountDianPao = new(int32)
	ret.CountAnGang = new(int32)
	ret.CountMingGang = new(int32)
	ret.CountDianGang = new(int32)
	ret.CountChaJiao = new(int32)
	ret.CountLianZhuang = new(int32)
	return ret
}

func NewGame_AckDissolveDesk() *ddproto.Game_AckDissolveDesk {
	ret := &ddproto.Game_AckDissolveDesk{}
	ret.DeskId = new(int32)
	ret.Header = NewHeader()
	ret.PassWord = new(string)
	ret.UserId = new(uint32)
	return ret
}

func NewGame_AckGameRecord() *ddproto.Game_AckGameRecord {
	ret := &ddproto.Game_AckGameRecord{}
	ret.UserId = new(uint32)
	ret.GameId = new(int32)
	ret.Header = NewHeader()
	return ret
}

//麻将战绩的bean
func NewBeanGameRecord() *ddproto.BeanGameRecord {
	ret := &ddproto.BeanGameRecord{}
	ret.Header = NewHeader()
	ret.DeskId = new(int32)
	ret.BeginTime = new(string)
	ret.Id = new(int32)
	return ret
}

func NewBeanUserRecord() *ddproto.BeanUserRecord {
	result := &ddproto.BeanUserRecord{}
	result.NickName = new(string)
	result.Header = NewHeader()
	result.UserId = new(uint32)
	result.WinAmount = new(int64)
	return result
}

func NewGame_AckActGuo() *ddproto.Game_AckActGuo {
	ret := &ddproto.Game_AckActGuo{}
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	return ret
}

func NewGame_AckExchangeCards() *ddproto.Game_AckExchangeCards {
	ret := &ddproto.Game_AckExchangeCards{}
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	return ret
}

func NewGame_ExchangeCardsEnd() *ddproto.Game_ExchangeCardsEnd {
	ret := &ddproto.Game_ExchangeCardsEnd{}
	ret.Header = NewHeader()
	ret.ExchangeNum = new(int32)
	ret.ExchangeType = new(int32)
	return ret
}

func NewGame_BroadcastBeginExchange() *ddproto.Game_BroadcastBeginExchange {
	ret := &ddproto.Game_BroadcastBeginExchange{}
	ret.Reconnect = new(bool)
	return ret
}

func NewGame_AckJiaoInfos() *ddproto.GameAckJiaoinfos {
	ret := &ddproto.GameAckJiaoinfos{}
	ret.Header = NewHeader()
	ret.JiaoInfos = nil
	return ret
}
