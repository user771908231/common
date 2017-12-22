package roomService

import (
	"casino_common/proto/ddproto"
	"errors"
	"casino_common/common/userService"
	"fmt"
	"casino_common/common/log"
)

//获取建房房费
func DecUserRoomcardToCreateDesk(billType ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE, gameId ddproto.CommonEnumGame, boardsCout int32, gamerNum int32, chanelId int32, createUser uint32) (err error) {
	log.T("billType:%v gameId:%v boardsCout:%v gamerNum:%v chanelId:%v createUser:%v", billType, gameId, boardsCout, gamerNum, chanelId, createUser)
	createFee := getDeskCreateFee(gameId, boardsCout, gamerNum, chanelId)
	switch billType {
	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_OWNER_PAY:
		if userService.GetUserRoomCard(createUser) < createFee {
			return errors.New("房卡不足！")
		}
		userService.DECRUserRoomcard(createUser, createFee, int32(gameId), "建房-房主扣卡")
		return nil
	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_AA_PAY:
		//不扣房卡
		return nil

	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_BIG_WINER_PAY:
		//不扣房卡
		return nil
	}

	return nil
}

//是否有足够的房卡进房
func HasEnoughRoomcardToEnterDesk(billType ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE, gameId ddproto.CommonEnumGame, boardsCout int32, gamerNum int32, chanelId int32, enterUser uint32) (err error) {
	log.T("billType:%v gameId:%v boardsCout:%v gamerNum:%v chanelId:%v enterUser:%v", billType, gameId, boardsCout, gamerNum, chanelId, enterUser)
	switch billType {
	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_OWNER_PAY:
		//房主已扣，自动进
		return nil
	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_AA_PAY:
		needAAFee := getDeskEnterAAFee(gameId, boardsCout, gamerNum, chanelId)
		if userService.GetUserRoomCard(enterUser) < needAAFee {
			return errors.New(fmt.Sprintf("房卡不足%d张，进房失败！", needAAFee))
		}

	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_BIG_WINER_PAY:
		needAAFee := getDeskCreateFee(gameId, boardsCout, gamerNum, chanelId)
		if userService.GetUserRoomCard(enterUser) < needAAFee {
			return errors.New(fmt.Sprintf("房卡不足%d张，进房失败！", needAAFee))
		}

	}

	return nil
}

//结算 扣卡类型 游戏id 最大圈数 最大人数 渠道id 所有玩家 大赢家 当前局数 房主
func DoDecUsersRoomcard(billType ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE, gameId ddproto.CommonEnumGame, boardsCout int32, gamerNum int32, chanelId int32, allUsers []uint32, winUsers []uint32, currRound int32, owner uint32) (err error) {
	log.T("billType:%v gameId:%v boardsCout:%v gamerNum:%v chanelId:%v allUsers:%v winUsers:%v currRound:%v owner:%v",
		billType, gameId, boardsCout, gamerNum, chanelId, allUsers, winUsers, currRound, owner)
	switch billType {
	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_OWNER_PAY:
		//房主已扣
		if currRound <= 1 {
			//提前结束则返还房主房卡
			needRoomcard := getDeskCreateFee(gameId, boardsCout, gamerNum, chanelId)
			userService.INCRUserRoomcard(owner, needRoomcard, int32(gameId), "提前结束，返还房主房卡。")
		}
		return nil
	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_AA_PAY:
		if currRound <= 1 {
			return nil
		}
		needRoomcard := getDeskEnterAAFee(gameId, boardsCout, gamerNum, chanelId)
		for _,u := range allUsers {
			userService.DECRUserRoomcard(u, needRoomcard, int32(gameId), "AA扣卡")
		}

	case ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE_BIG_WINER_PAY:
		if currRound <= 1 {
			return nil
		}
		needRoomcard := getDeskBigwinerOneBillFee(gameId, boardsCout, gamerNum, chanelId, allUsers, winUsers)
		for _,u := range winUsers {
			userService.DECRUserRoomcard(u, needRoomcard, int32(gameId), "大赢家扣卡")
		}
	}
	return nil
}

//===============================================房卡配置=================================================

//房卡建房配置
func getDeskCreateFee(gameId ddproto.CommonEnumGame, boardCout int32, gamerNum int32, chanelId int32) (needRoomCard int64) {
	switch gameId {
	//红中转转  长沙麻将
	case ddproto.CommonEnumGame_GID_ZHUANZHUAN, ddproto.CommonEnumGame_GID_MAHJONG:
		switch {
		case boardCout == 8 && gamerNum == 4:
			needRoomCard = 4
		case boardCout == 8 && gamerNum == 3:
			needRoomCard = 3
		case boardCout == 8 && gamerNum == 2:
			needRoomCard = 2
		case boardCout == 16 && gamerNum == 4:
			needRoomCard = 8
		case boardCout == 16 && gamerNum == 3:
			needRoomCard = 6
		case boardCout == 16 && gamerNum == 4:
			needRoomCard = 4
		}
	//牛牛 跑得快
	case ddproto.CommonEnumGame_GID_NIUNIUJINGDIAN:
		switch boardCout {
		case 10:
			needRoomCard = 4
		case 20:
			needRoomCard = 8
		}
	case ddproto.CommonEnumGame_GID_PDK:
		switch {
		case boardCout == 10 && gamerNum == 3:
			needRoomCard = 3
		case boardCout == 10 && gamerNum == 2:
			needRoomCard = 2
		case boardCout == 20 && gamerNum == 3:
			needRoomCard = 6
		case boardCout == 20 && gamerNum == 2:
			needRoomCard = 4
		}
	//跑胡子
	case ddproto.CommonEnumGame_GID_PAOHUZI:
		switch boardCout {
		case 200:
			needRoomCard = 3
		case 400:
			needRoomCard = 6
		}
	}

	return
}

//房卡AA进房配置
func getDeskEnterAAFee(gameId ddproto.CommonEnumGame, boardCout int32, gamerNum int32, chanelId int32) (needRoomCard int64) {
	switch gameId {
	case ddproto.CommonEnumGame_GID_ZHUANZHUAN, ddproto.CommonEnumGame_GID_MAHJONG:
		//红中转转  长沙麻将
		switch boardCout {
		case 8:
			needRoomCard = 1
		case 16:
			needRoomCard = 2
		}
	case ddproto.CommonEnumGame_GID_NIUNIUJINGDIAN:
		//牛牛 跑得快
		switch boardCout {
		case 10:
			needRoomCard = 1
		case 20:
			needRoomCard = 2
		}
	case ddproto.CommonEnumGame_GID_PDK:
		//跑得快
		switch boardCout {
		case 10:
			needRoomCard = 1
		case 20:
			needRoomCard = 2
		}
		//跑胡子
	case ddproto.CommonEnumGame_GID_PAOHUZI:
		switch boardCout {
		case 200:
			needRoomCard = 1
		case 400:
			needRoomCard = 2
		}
	}

	return
}

//房卡大赢家单人结算配置
func getDeskBigwinerOneBillFee(gameId ddproto.CommonEnumGame, boardCout int32, gamerNum int32, chanelId int32, allUsers []uint32, winUsers []uint32) (needRoomCard int64) {
	switch gameId {
	case ddproto.CommonEnumGame_GID_NIUNIUJINGDIAN:
		all_user_num := int64(len(allUsers))
		aaFee := getDeskEnterAAFee(gameId, boardCout, gamerNum, chanelId)
		needRoomCard = getDeskCreateFee(gameId, boardCout, gamerNum, chanelId) / all_user_num
		if needRoomCard == 0 {
			needRoomCard = aaFee
		}
		return
	}
	aaFee := getDeskEnterAAFee(gameId, boardCout, gamerNum, chanelId)
	needRoomCard = getDeskCreateFee(gameId, boardCout, gamerNum, chanelId) / int64(len(winUsers))
	if needRoomCard == 0 {
		needRoomCard = aaFee
	}

	return
}

