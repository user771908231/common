package commonNewPorot

import (
	"casino_common/proto/ddproto/mjproto"
)

func SuccessHeader() *mjproto.ProtoHeader {
	header := NewMJHeader()
	*header.Code = 0
	return header
}

func ErrorHeader() *mjproto.ProtoHeader {
	header := NewMJHeader()
	*header.Code = -1
	return header
}

func NewMJHeader() *mjproto.ProtoHeader {
	ret := &mjproto.ProtoHeader{}
	ret.UserId = new(uint32)
	ret.Code = new(int32)
	ret.Error = new(string)
	return ret
}

func NewGame_AckCreateRoom() *mjproto.Game_AckCreateRoom {
	ret := &mjproto.Game_AckCreateRoom{}
	ret.Header = NewMJHeader()
	ret.DeskId = new(int32)
	ret.Password = new(string)
	ret.UserBalance = new(int64)
	ret.CreateFee = new(int64)
	return ret
}

func NewGame_AckReady() *mjproto.Game_AckReady {
	ret := &mjproto.Game_AckReady{}
	ret.Header = NewMJHeader()
	ret.UserId = new(uint32)
	return ret
}

func NewGame_AckEnterRoom() *mjproto.Game_AckEnterRoom {
	ret := &mjproto.Game_AckEnterRoom{}
	ret.Header = NewMJHeader()
	return ret
}

func NewGame_BroadcastBeginDingQue() *mjproto.Game_BroadcastBeginDingQue {
	ret := &mjproto.Game_BroadcastBeginDingQue{}
	ret.Reconnect = new(bool)
	return ret
}

func NewGame_DingQue() *mjproto.Game_DingQue {
	que := &mjproto.Game_DingQue{}
	que.Header = NewMJHeader()
	que.Color = new(int32)
	que.UserId = new(uint32)
	return que
}

func NewRoomTypeInfo() *mjproto.RoomTypeInfo {
	ret := &mjproto.RoomTypeInfo{}
	ret.BaseValue = new(int64)
	ret.BoardsCout = new(int32)
	ret.CapMax = new(int64)
	ret.CardsNum = new(int32)
	ret.MjRoomType = new(mjproto.MJRoomType)
	ret.PlayOptions = NewPlayOptions()
	ret.Settlement = new(int32)
	return ret
}

func NewPlayOptions() *mjproto.PlayOptions {
	ret := &mjproto.PlayOptions{}
	ret.DianGangHuaRadio = new(int32)
	ret.HuRadio = new(int32)
	ret.ZiMoRadio = new(int32)
	return ret
}

func NewGame_SendGameInfo() *mjproto.Game_SendGameInfo {
	ret := &mjproto.Game_SendGameInfo{}
	ret.Header = NewMJHeader()
	ret.SenderUserId = new(uint32)
	ret.IsReconnect = new(mjproto.RECONNECT_TYPE)
	return ret
}

//返回deskGameInfo
func NewDeskGameInfo() *mjproto.DeskGameInfo {
	ret := &mjproto.DeskGameInfo{}
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
func NewPlayerInfo() *mjproto.PlayerInfo {
	info := &mjproto.PlayerInfo{}
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
func NewPlayerCard() *mjproto.PlayerCard {
	card := &mjproto.PlayerCard{}
	card.UserId = new(uint32)
	return card
}

func NewCardInfo() *mjproto.CardInfo {
	cardInfo := &mjproto.CardInfo{}
	cardInfo.Id = new(int32)
	cardInfo.Type = new(int32)
	cardInfo.Value = new(int32)
	return cardInfo
}

func NewComposeCard() *mjproto.ComposeCard {
	ret := &mjproto.ComposeCard{}
	ret.Type = new(int32)
	ret.Value = new(int32)
	return ret
}

func NewGame_OverTurn() *mjproto.Game_OverTurn {
	ret := &mjproto.Game_OverTurn{}
	ret.ActType = new(int32)
	ret.CanGang = new(bool)
	ret.CanPeng = new(bool)
	ret.CanHu = new(bool)
	ret.Header = NewMJHeader()
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

func NewSJHGame_OverTurn() *mjproto.Game_OverTurn {
	ret := &mjproto.Game_OverTurn{}
	ret.ActType = new(int32)
	ret.CanGang = new(bool)
	ret.CanPeng = new(bool)
	ret.CanHu = new(bool)
	ret.Header = NewMJHeader()
	ret.UserId = new(uint32)
	ret.NextUserId = new(uint32)
	ret.Time = new(int32)
	ret.PaiCount = new(int32)
	ret.CanGuo = new(bool)
	ret.CanChi = new(bool)
	return ret
}

func NewHaiNanGame_OverTurn() *mjproto.Game_OverTurn {
	ret := &mjproto.Game_OverTurn{}
	ret.ActType = new(int32)
	ret.CanGang = new(bool)
	ret.CanPeng = new(bool)
	ret.CanHu = new(bool)
	ret.Header = NewMJHeader()
	ret.UserId = new(uint32)
	ret.NextUserId = new(uint32)
	ret.Time = new(int32)
	ret.PaiCount = new(int32)
	ret.CanGuo = new(bool)
	ret.CanChi = new(bool)
	ret.CanBu = new(bool)
	return ret
}

//转转麻将overturn
func NewMJZhZhGame_OverTurn() *mjproto.Game_OverTurn {
	ret := &mjproto.Game_OverTurn{}
	ret.ActType = new(int32)
	ret.CanGang = new(bool)
	ret.CanPeng = new(bool)
	ret.CanHu = new(bool)
	ret.Header = NewMJHeader()
	ret.UserId = new(uint32)
	ret.NextUserId = new(uint32)
	ret.Time = new(int32)
	ret.PaiCount = new(int32)
	ret.CanGuo = new(bool)
	return ret
}

func NewGame_DealCards() *mjproto.Game_DealCards {
	ret := &mjproto.Game_DealCards{}
	ret.Header = NewMJHeader()
	ret.DealerUserId = new(uint32)
	return ret
}

func NewGame_Opening() *mjproto.Game_Opening {
	ret := &mjproto.Game_Opening{}
	ret.Header = NewMJHeader()
	ret.CurrPlayCount = new(int32)
	ret.Dice1 = new(int32)
	ret.Dice2 = new(int32)
	return ret
}

func NewGame_AckQuickConn() *mjproto.Game_AckQuickConn {
	ret := &mjproto.Game_AckQuickConn{}
	ret.CurrVersion = new(int32)
	ret.DownloadUrl = new(string)
	ret.Header = NewMJHeader()
	ret.IsMaintain = new(int32)
	ret.IsUpdate = new(int32)
	ret.ReleaseTag = new(int32)
	ret.VersionInfo = new(string)
	return ret
}

func NewGame_DingQueEnd() *mjproto.Game_DingQueEnd {
	ret := &mjproto.Game_DingQueEnd{}
	return ret
}

func NewDingQueEndBean() *mjproto.DingQueEndBean {
	ret := &mjproto.DingQueEndBean{}
	ret.UserId = new(uint32)
	ret.Flower = new(int32)
	return ret
}

func NewGame_AckActHu() *mjproto.Game_AckActHu {
	ret := &mjproto.Game_AckActHu{}
	ret.Header = NewMJHeader()
	ret.UserIdIn = new(uint32)
	ret.UserIdOut = new(uint32)
	ret.HuType = new(int32)
	ret.IsZiMo = new(bool)
	return ret
}

func NewGame_AckSendOutCard() *mjproto.Game_AckSendOutCard {
	ret := &mjproto.Game_AckSendOutCard{}
	ret.Header = NewMJHeader()
	ret.Result = new(int32)
	ret.UserId = new(uint32)
	return ret
}

func NewGame_AckActGang() *mjproto.Game_AckActGang {
	ret := &mjproto.Game_AckActGang{}
	ret.GangType = new(int32)
	ret.Header = NewMJHeader()
	ret.UserIdIn = new(uint32)
	ret.UserIdOut = new(uint32)
	return ret
}

func NewGame_AckActPeng() *mjproto.Game_AckActPeng {
	ret := &mjproto.Game_AckActPeng{}
	ret.Header = NewMJHeader()
	ret.UserIdIn = new(uint32)
	ret.UserIdOut = new(uint32)
	ret.JiaoInfos = nil
	return ret
}

func NewGame_AckActChi() *mjproto.Game_AckActChi {
	ret := &mjproto.Game_AckActChi{}
	ret.Header = NewMJHeader()
	ret.UserIdIn = new(uint32)
	ret.UserIdOut = new(uint32)
	ret.JiaoInfos = nil
	return ret
}

func NewGame_AckLogin() *mjproto.Game_AckLogin {
	ret := &mjproto.Game_AckLogin{}
	ret.Header = NewMJHeader()
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

func NewGame_SendCurrentResult() *mjproto.Game_SendCurrentResult {
	ret := &mjproto.Game_SendCurrentResult{}
	ret.Header = NewMJHeader()
	return ret

}

func NewGame_SendEndLottery() *mjproto.Game_SendEndLottery {
	ret := &mjproto.Game_SendEndLottery{}
	ret.Header = NewMJHeader()
	return ret
}

func NewWinCoinInfo() *mjproto.WinCoinInfo {
	ret := &mjproto.WinCoinInfo{}
	ret.CardTitle = new(string)
	ret.Coin = new(int64)
	ret.HuCount = new(int32)
	ret.IsDealer = new(bool)
	ret.UserId = new(uint32)
	ret.NickName = new(string)
	ret.WinCoin = new(int64)
	return ret
}

func NewEndLotteryInfo() *mjproto.EndLotteryInfo {
	ret := &mjproto.EndLotteryInfo{}
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

func NewGame_AckDissolveDesk() *mjproto.Game_AckDissolveDesk {
	ret := &mjproto.Game_AckDissolveDesk{}
	ret.DeskId = new(int32)
	ret.Header = NewMJHeader()
	ret.PassWord = new(string)
	ret.UserId = new(uint32)
	return ret
}

func NewGame_AckGameRecord() *mjproto.Game_AckGameRecord {
	ret := &mjproto.Game_AckGameRecord{}
	ret.UserId = new(uint32)
	ret.GameId = new(int32)
	ret.Header = NewMJHeader()
	return ret
}

//麻将战绩的bean
func NewBeanGameRecord() *mjproto.BeanGameRecord {
	ret := &mjproto.BeanGameRecord{}
	ret.Header = NewMJHeader()
	ret.DeskId = new(int32)
	ret.BeginTime = new(string)
	ret.Id = new(int32)
	return ret
}

func NewBeanUserRecord() *mjproto.BeanUserRecord {
	result := &mjproto.BeanUserRecord{}
	result.NickName = new(string)
	result.Header = NewMJHeader()
	result.UserId = new(uint32)
	result.WinAmount = new(int64)
	return result
}

func NewGame_AckActGuo() *mjproto.Game_AckActGuo {
	ret := &mjproto.Game_AckActGuo{}
	ret.Header = NewMJHeader()
	ret.UserId = new(uint32)
	return ret
}

func NewGame_AckExchangeCards() *mjproto.Game_AckExchangeCards {
	ret := &mjproto.Game_AckExchangeCards{}
	ret.Header = NewMJHeader()
	ret.UserId = new(uint32)
	return ret
}

func NewGame_ExchangeCardsEnd() *mjproto.Game_ExchangeCardsEnd {
	ret := &mjproto.Game_ExchangeCardsEnd{}
	ret.Header = NewMJHeader()
	ret.ExchangeNum = new(int32)
	ret.ExchangeType = new(int32)
	return ret
}

func NewGame_BroadcastBeginExchange() *mjproto.Game_BroadcastBeginExchange {
	ret := &mjproto.Game_BroadcastBeginExchange{}
	ret.Reconnect = new(bool)
	return ret
}

func NewGame_AckJiaoInfos() *mjproto.GameAckJiaoinfos {
	ret := &mjproto.GameAckJiaoinfos{}
	ret.Header = NewMJHeader()
	ret.JiaoInfos = nil
	return ret
}
