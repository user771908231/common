package utils

import (
	"casino_common/proto/ddproto"
	//"casino_common/common/log"
)

/*

	gid : 游戏id
	roomType : 金币场 还是朋友桌
 */
func GetGameName(gid int32, roomType int32) string {
	//log.T("查询游戏的名字 gid:%v roomType %v", gid, roomType)
	if gid == int32(ddproto.CommonEnumGame_GID_TEXAS) {
		return "德州扑克"
	} else if gid == int32(ddproto.CommonEnumGame_GID_MAHJONG) {
		if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN) {
			return "麻将金币场"
		} else if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND) {
			return "麻将朋友桌"
		} else {
			return "麻将"
		}
	} else if gid == int32(ddproto.CommonEnumGame_GID_DDZ) {
		if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN) {
			return "斗地主金币场"
		} else if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND) {
			return "斗地主朋友桌"
		} else {
			return "斗地主"
		}
	} else if gid == int32(ddproto.CommonEnumGame_GID_ZJH) {
		return "炸金花"
	} else if gid == int32(ddproto.CommonEnumGame_GID_NIUNIUJINGDIAN) {
		return "牛牛朋友桌"
	} else if gid == int32(ddproto.CommonEnumGame_GID_PDK) {
		if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN) {
			return "跑得快金币场"
		} else if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND) {
			return "跑得快朋友桌"
		} else {
			return "跑得快"
		}
	}
	return "其他"

}
