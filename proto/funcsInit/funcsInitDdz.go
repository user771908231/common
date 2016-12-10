package commonNewPorot

import (
	"casino_common/proto/ddproto"
)


//备份用
func NewPDdzbak() *ddproto.DdzSrvBak {
	ret := new(ddproto.DdzSrvBak)
	return ret
}

//Desk
func NewPDdzDesk() *ddproto.DdzSrvDesk {
	ret := new(ddproto.DdzSrvDesk)
	ret.DeskId = new(int32)
	ret.Key = new(string)
	ret.UserCountLimit = new(int32)
	ret.Owner = new(uint32)
	ret.DizhuPaiUser = new(uint32)
	ret.DiZhuUserId = new(uint32)
	ret.ActiveUserId = new(uint32)
	ret.BaseValue = new(int64)
	ret.QingDizhuValue = new(int64)
	ret.WinValue = new(int64)
	ret.DdzType = new(int32)
	ret.RoomType = new(int32)
	ret.BoardsCount = new(int32)
	ret.CurrPlayCount = new(int32)
	ret.TotalPlayCount = new(int32)
	ret.CapMax = new(int64)
	ret.IsJiaoFen = new(bool)
	ret.FootRate = new(int32)
	return ret
}

func NewPDdzUser() *ddproto.DdzSrvUser {
	user := new(ddproto.DdzSrvUser)
	user.UserId = new(uint32)
	user.DeskId = new(int32)
	user.RoomId = new(int32)
	user.Status = new(int32)
	user.IsBreak = new(bool)
	user.IsLeave = new(bool)
	user.QiangDiZhuStatus = new(int32)
	user.JiabeiStatus = new(int32)
	user.LaDaoStatus = new(int32)
	user.PlayRate = new(int32)
	user.Bomb = new(int32)
	user.JiaoScore = new(int32)
	return user
}

func NewDdzSrvUserStatistics() *ddproto.DdzSrvUserStatistics {
	ret := new(ddproto.DdzSrvUserStatistics)
	ret.CountLose = new(int32)
	ret.CountWin = new(int32)
	return ret
}

func NewPDdzDeskTongJi() *ddproto.DdzSrvDeskTongJi {
	pDdzDeskTongJi := new(ddproto.DdzSrvDeskTongJi)
	pDdzDeskTongJi.CountQiangDiZhu = new(int32)
	return pDdzDeskTongJi
}

func NewPPokerPai() *ddproto.CommonSrvPokerPai {
	pai := new(ddproto.CommonSrvPokerPai)
	pai.Des = new(string)
	pai.Flower = new(int32)
	pai.Id = new(int32)
	pai.Name = new(string)
	pai.Value = new(int32)
	return pai
}

func NewPOutPokerPais() *ddproto.DdzSrvOutPokerPais {
	ret := new(ddproto.DdzSrvOutPokerPais)
	ret.KeyValue = new(int32)
	ret.Type = new(int32)
	ret.IsBomb = new(bool)

	ret.CountYizhang = new(int32)
	ret.CountDuizi = new(int32)
	ret.CountSanzhang = new(int32)
	ret.CountSizhang = new(int32)

	ret.UserId = new(uint32)
	return ret
}

func NewPDdzBill() *ddproto.DdzSrvBill {
	ret := new(ddproto.DdzSrvBill)
	ret.WinCoin = new(int64)
	return ret
}

func NewPDdzBillBean() *ddproto.DdzSrvBillBean {
	b := new(ddproto.DdzSrvBillBean)
	b.Coin = new(int64)
	b.WinUser = new(uint32)
	b.LoseUser = new(uint32)
	b.Desc = new(string)
	return b
}

func NewDdzSrvGameData() *ddproto.DdzSrvGameData {
	ret := new(ddproto.DdzSrvGameData)
	return ret
}

func NewRoomTypeInfo() *ddproto.DdzBaseRoomTypeInfo {
	ret := new(ddproto.DdzBaseRoomTypeInfo)
	ret.BaseValue = new(int64)
	ret.BoardsCount = new(int32)
	ret.CapMax = new(int64)
	ret.IsJiaoFen = new(bool)
	ret.RoomType = new(ddproto.DdzEnumRoomType)
	return ret
}



//创建房间
func NewDdzAckCreateDesk() *ddproto.DdzAckCreateDesk {
	ret := new(ddproto.DdzAckCreateDesk)
	ret.CreateFee = new(int64)
	ret.DeskId = new(int32)
	ret.Header = NewHeader()
	ret.Password = new(string)
	ret.UserBalance = new(int64)
	return ret
}

//进入房间成功
func NewDdzAckEnterDesk() *ddproto.DdzAckEnterDesk {
	ret := new(ddproto.DdzAckEnterDesk)
	ret.Header = NewHeader()
	return ret
}

//发送过的回复
func NewDdzActGuoAck() *ddproto.DdzAckGuoAck {
	ret := new(ddproto.DdzAckGuoAck)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	return ret
}


//发送发牌
func NewDdzOverTurn() *ddproto.DdzBcOverTurn {
	ret := new(ddproto.DdzBcOverTurn)
	ret.UserId = new(uint32)
	ret.ActType = new(ddproto.DdzEnumActType)
	ret.CanDouble = new(bool)
	ret.CanOutCards = new(bool)
	ret.PullOrPush = new(int32)
	ret.Time = new(int32)
	ret.JiaoScore = new(ddproto.DdzEnumJdScore)
	return ret
}


//准备回复
func NewDdzAckReady() *ddproto.DdzAckReady {
	ret := new(ddproto.DdzAckReady)
	ret.UserId = new(uint32)
	ret.Msg = new(string)
	return ret
}


//发送消息
func NewGameMessage() *ddproto.DdzBcMessage {
	ret := new(ddproto.DdzBcMessage)
	ret.Header = NewHeader()
	ret.Id = new(int32)
	ret.Msg = new(string)
	ret.MsgType = new(int32)
	ret.UserId = new(uint32)
	return ret
}

func NewDdzSendGameInfo() *ddproto.DdzBcGameInfo {
	ret := new(ddproto.DdzBcGameInfo)
	ret.Header = NewHeader()
	ret.IsReconnect = new(int32)
	ret.SenderUserId = new(uint32)
	return ret
}

func NewDdzOpening() *ddproto.DdzBcOpening {
	ret := new(ddproto.DdzBcOpening)
	ret.Header = NewHeader()
	return ret
}

func NewDdzDealCards() *ddproto.DdzBcDealCards {
	ret := new(ddproto.DdzBcDealCards)
	ret.Header = NewHeader()
	return ret
}

func NewDdzJiaoDiZhuAck() *ddproto.DdzAckJiaoDiZhu {
	ret := new(ddproto.DdzAckJiaoDiZhu)
	ret.Jiao = new(bool)
	ret.UserId = new(uint32)
	ret.Header = NewHeader()
	return ret
}

func NewDdzRobDiZhuAck() *ddproto.DdzAckRobDiZhu {
	ret := new(ddproto.DdzAckRobDiZhu)
	ret.UserId = new(uint32)
	ret.Header = NewHeader()
	ret.Rob = new(bool)
	return ret
}

func NewDdzDouble() *ddproto.DdzAckDouble {
	ret := new(ddproto.DdzAckDouble)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	ret.Double = new(ddproto.DdzEnumDoubleType)
	return ret
}

func NewDdzDoubleAck() *ddproto.DdzAckDouble {
	ret := new(ddproto.DdzAckDouble)
	ret.Double = new(ddproto.DdzEnumDoubleType)
	ret.UserId = new(uint32)
	ret.Header = NewHeader()
	return ret
}

func NewDdzShowHandPokers() *ddproto.DdzReqShowHandPokers {
	ret := new(ddproto.DdzReqShowHandPokers)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	return ret
}

func NewDdzShowHandPokersAck() *ddproto.DdzAckShowHandPokers {
	ret := new(ddproto.DdzAckShowHandPokers)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	return ret
}

func NewDdzMenuZhua() *ddproto.DdzReqMenuZhua {
	ret := new(ddproto.DdzReqMenuZhua)
	ret.UserId = new(uint32)
	ret.Header = NewHeader()
	return ret
}

func NewDdzMenuZhuaAck() *ddproto.DdzAckMenuZhua {
	ret := new(ddproto.DdzAckMenuZhua)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	return ret
}

func NewDdzSeeCards() *ddproto.DdzReqSeeCards {
	ret := new(ddproto.DdzReqSeeCards)
	ret.UserId = new(uint32)
	ret.Header = NewHeader()
	return ret
}

func NewDdzSeeCardsAck() *ddproto.DdzAckSeeCards {
	ret := new(ddproto.DdzAckSeeCards)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	return ret
}

func NewDdzPull() *ddproto.DdzReqPull {
	ret := new(ddproto.DdzReqPull)
	ret.UserId = new(uint32)
	ret.Header = NewHeader()
	return ret
}

func NewDdzPullAck() *ddproto.DdzAckPull {
	ret := new(ddproto.DdzAckPull)
	ret.Act = new(ddproto.DdzEnumLaDaoType)
	ret.UserId = new(uint32)
	ret.Header = NewHeader()
	return ret
}

func NewDdzOutCards() *ddproto.DdzReqOutCards {
	ret := new(ddproto.DdzReqOutCards)
	ret.Header = NewHeader()
	return ret
}

func NewDdzOutCardsAck() *ddproto.DdzAckOutCards {
	ret := new(ddproto.DdzAckOutCards)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	ret.CardType = new(ddproto.DdzEnumPaiType)
	ret.RemainPokers = new(int32)
	return ret
}

func NewDdzActGuo() *ddproto.DdzReqActGuo {
	ret := new(ddproto.DdzReqActGuo)
	ret.UserId = new(uint32)
	ret.Header = NewHeader()
	return ret
}

func NewDdzStartPlay() *ddproto.DdzBcStartPlay {
	ret := new(ddproto.DdzBcStartPlay)
	ret.Header = NewHeader()
	ret.FootRate = new(int32)
	ret.Dizhu = new(uint32)
	return ret
}

func NewDdzSendCurrentResult() *ddproto.DdzBaCurrentResult {
	ret := new(ddproto.DdzBaCurrentResult)
	ret.Header = NewHeader()
	return ret
}

func NewDdzSendEndLottery() *ddproto.DdzBcEndLottery {
	ret := new(ddproto.DdzBcEndLottery)
	ret.Header = NewHeader()
	return ret
}

func NewEndLotteryInfo() *ddproto.DdzBaseEndLotteryInfo {
	ret := new(ddproto.DdzBaseEndLotteryInfo)
	ret.UserId = new(uint32)
	ret.NickName = new(string)
	ret.BigWin = new(bool)
	ret.WinCoin = new(int64)
	ret.MaxWinCoin = new(int32)
	ret.CountBomb = new(int32)
	ret.CountWin = new(int32)
	ret.CountLose = new(int32)
	return ret
}

func NewDdzDissolveDesk() *ddproto.DdzReqDissolveDesk {
	ret := new(ddproto.DdzReqDissolveDesk)
	ret.Header = NewHeader()
	ret.UserId = new(uint32)
	return ret
}

func NewDdzAckDissolveDesk() *ddproto.DdzAckDissolveDesk {
	ret := new(ddproto.DdzAckDissolveDesk)
	ret.UserId = new(uint32)
	ret.DeskId = new(int32)
	ret.PassWord = new(string)
	return ret
}

func NewDdzLeaveDesk() *ddproto.DdzReqLeaveDesk {
	ret := new(ddproto.DdzReqLeaveDesk)
	ret.Header = NewHeader()
	return ret
}

func NewDdzAckLeaveDesk() *ddproto.DdzAckLeaveDesk {
	ret := new(ddproto.DdzAckLeaveDesk)
	ret.Header = NewHeader()
	return ret
}

func NewDdzMessage() *ddproto.DdzReqMessage {
	ret := new(ddproto.DdzReqMessage)
	ret.Header = NewHeader()
	ret.DeskId = new(int32)
	ret.UserId = new(uint32)
	ret.Id = new(int32)
	ret.Msg = new(string)
	ret.MsgType = new(int32)
	return ret
}

func NewDdzSendMessage() *ddproto.DdzBcMessage {
	ret := new(ddproto.DdzBcMessage)
	ret.MsgType = new(int32)
	ret.Msg = new(string)
	ret.Id = new(int32)
	ret.UserId = new(uint32)
	ret.Header = NewHeader()
	return ret
}

func NewDdzAckGameRecord() *ddproto.DdzAckGameRecord {
	ret := new(ddproto.DdzAckGameRecord)
	ret.UserId = new(uint32)
	ret.Header = NewHeader()
	return ret
}

func NewDdzDeskInfo() *ddproto.DdzBaseDeskInfo {
	ret := new(ddproto.DdzBaseDeskInfo)
	ret.GameStatus = new(int32)
	ret.PlayerNum = new(int32)
	ret.ActiveUserId = new(uint32)
	ret.ActionTime = new(int32)
	ret.NInitActionTime = new(int32)
	ret.InitRoomCoin = new(int64)
	ret.CurrPlayCount = new(int32)
	ret.TotalPlayCount = new(int32)
	ret.RoomNumber = new(string)
	ret.DiZhuUserId = new(uint32)
	ret.FootRate = new(int32)
	ret.PlayRate = new(int32)
	return ret
}

func NewPlayerInfo() *ddproto.DdzBasePlayerInfo {
	ret := new(ddproto.DdzBasePlayerInfo)
	ret.IsDiZhu = new(bool)
	ret.Coin = new(int64)
	ret.NickName = new(string)
	ret.Sex = new(int32)
	ret.UserId = new(uint32)
	ret.IsOwner = new(bool)
	ret.BReady = new(int32)
	ret.OnlineStatus = new(int32)
	ret.PlayRate = new(int32)
	ret.RemainPaiCount = new(int32)
	return ret
}

func NewPoker() *ddproto.ClientBasePoker {
	p := new(ddproto.ClientBasePoker)
	p.Num = new(int32)
	p.Id = new(int32)
	p.Suit = new(ddproto.CommonEnumPokerColor)
	return p
}

func NewWinCoinInfo() *ddproto.DdzBaseWinCoinInfo {
	ret := new(ddproto.DdzBaseWinCoinInfo)
	ret.BaseValue = new(int32)
	ret.Coin = new(int64)
	ret.Description = new(string)
	ret.IsDiZhu = new(bool)
	ret.NickName = new(string)
	ret.Rate = new(int32)
	ret.UserId = new(uint32)
	ret.WinCoin = new(int64)
	ret.Bomb = new(int32)
	return ret
}

func NewDdzSrvDesk() *ddproto.DdzSrvDesk {
	ret := new(ddproto.DdzSrvDesk)
	ret.DeskId = new(int32)
	ret.Key = new(string)
	ret.UserCountLimit = new(int32)
	ret.Owner = new(uint32)
	ret.DizhuPaiUser = new(uint32)
	ret.DiZhuUserId = new(uint32)
	ret.ActiveUserId = new(uint32)
	ret.BaseValue = new(int64)
	ret.QingDizhuValue = new(int64)
	ret.WinValue = new(int64)
	ret.DdzType = new(int32)
	ret.RoomType = new(int32)
	ret.BoardsCount = new(int32)
	ret.CurrPlayCount = new(int32)
	ret.TotalPlayCount = new(int32)
	ret.CapMax = new(int64)
	ret.IsJiaoFen = new(bool)
	ret.FootRate = new(int32)
	return ret
}
