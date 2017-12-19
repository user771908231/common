package roomService

import (
	"casino_common/proto/ddproto"
	"errors"
	"casino_common/common/userService"
	"fmt"
	"casino_common/common/log"
)

//房费扣除配置，用于支持AA扣房卡和大赢家扣房卡
type RoomcardBillOpt struct {
	BillType ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE  //结算类型
	GameId ddproto.CommonEnumGame  //游戏id
	ChanelId int32  //渠道id
	GamerNum int32  //游戏人数
	BoardsCout int32  //房间局数
	IsPreEnd bool  //是否为提前结束（解散为True不扣房卡）
	Owner uint32  //房主
	Users []uint32  //当前房间需要扣费的人
	BigWiner []uint32
}

//是否能够进房
func (opt *RoomcardBillOpt) HasEnoughRoomcard(user uint32) error {
	if opt.ChanelId != 73 && opt.ChanelId != 74 {
		log.T("仅支持黄圣棋牌， %v", opt)
		return errors.New("仅支持黄圣")
	}
	userRoomCard := userService.GetUserRoomCard(user)
	var needRoomCard int64 = 0
	switch opt.BillType {
	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_AA_PAY:
		//AA付
		needRoomCard = getHuangshengRoomcardConfig(opt.GameId, opt.BoardsCout) / int64(opt.GamerNum)
		if userRoomCard < needRoomCard {
			return errors.New(fmt.Sprintf("房卡不足%d张，进入房间失败。", needRoomCard))
		}
	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_BIG_WINER_PAY:
		//大赢家付
		needRoomCard = getHuangshengRoomcardConfig(opt.GameId, opt.BoardsCout)
		if userRoomCard < needRoomCard {
			return errors.New(fmt.Sprintf("房卡不足%d张，进入房间失败。", needRoomCard))
		}
	}
	return nil
}

//扣除房卡
func (opt *RoomcardBillOpt) DecrUserRoomCard() error {
	if opt.ChanelId != 73 && opt.ChanelId != 74 {
		log.T("仅支持黄圣棋牌， %v", opt)
		return errors.New("仅支持皇圣")
	}
	//提前结束不扣除房卡
	if opt.IsPreEnd == true {
		return nil
	}
	needRoomCard := getHuangshengRoomcardConfig(opt.GameId, opt.BoardsCout)
	switch opt.BillType {
	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_AA_PAY:
		//AA付
		one_user_need_roomcard := needRoomCard / int64(len(opt.Users))
		for _,u := range opt.Users {
			userService.DECRUserRoomcard(u, one_user_need_roomcard, int32(opt.GameId), "AA扣房卡")
		}
	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_BIG_WINER_PAY:
		//大赢家付
		one_user_need_roomcard := needRoomCard / int64(len(opt.BigWiner))
		for _,u := range opt.Users {
			userService.DECRUserRoomcard(u, one_user_need_roomcard, int32(opt.GameId), "大赢家扣房卡")
		}
	}
	return nil
}


//黄圣 房卡配置
func getHuangshengRoomcardConfig(gameId ddproto.CommonEnumGame, boardCout int32) (needRoomCard int64) {
	switch gameId {
	//红中转转  长沙麻将
	case ddproto.CommonEnumGame_GID_ZHUANZHUAN, ddproto.CommonEnumGame_GID_MAHJONG:
		switch boardCout {
		case 4:
			needRoomCard = 4
		case 8:
			needRoomCard = 8
		}
	//牛牛 跑得快
	case ddproto.CommonEnumGame_GID_NIUNIUJINGDIAN, ddproto.CommonEnumGame_GID_PDK:
		switch boardCout {
		case 10:
			needRoomCard = 4
		case 20:
			needRoomCard = 8
		}
	//跑胡子
	case ddproto.CommonEnumGame_GID_PAOHUZI:
		needRoomCard = 8
	}

	return
}
