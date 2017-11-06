package whiteListService

import (
	"casino_common/common/consts"
	"casino_common/common/log"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"fmt"
	"strings"
	"time"
)

type WhiteUser struct {
	UserId  uint32 //用户id
	WinRate int32  //胜率
}

//默认胜率
const DefaultWinRate int32 = 50

//白名单
var whiteList = map[int32][]WhiteUser{}
var lastRefreshTime time.Time

//刷新白名单
func RefreshWhiteList(gameId int32) {
	//更新频率30秒
	if time.Now().Sub(lastRefreshTime) < 5*time.Second {
		return
	}

	new_list := GetWhiteListByGid(gameId)
	whiteList[gameId] = new_list

	lastRefreshTime = time.Now()
	log.T("刷新白名单成功！%v", new_list)
}

//获取白名单列表
func GetWhiteListByGid(gameId int32) []WhiteUser {
	data := redisUtils.Get(fmt.Sprintf("%s_%02d", consts.RKEY_GAME_WHITE_LIST, gameId))
	list_str := strings.Split(data, ",")
	new_list := []WhiteUser{}

	if len(list_str) == 0 {
		log.T("刷新白名单失败，err:redis data nil")
		return new_list
	}

	for _, white_user_str := range list_str {
		white_user_arr := strings.Split(white_user_str, "=")
		new_white_user := WhiteUser{}
		if len(white_user_arr) == 1 {
			new_white_user.UserId = numUtils.String2Uint32(white_user_arr[0])
			new_white_user.WinRate = DefaultWinRate
		}
		if len(white_user_arr) == 2 {
			new_white_user.UserId = numUtils.String2Uint32(white_user_arr[0])
			new_white_user.WinRate = int32(numUtils.String2Int(white_user_arr[1]))
		}
		if new_white_user.UserId > 10000 && new_white_user.WinRate >= 0 && new_white_user.WinRate <= 100 {
			new_list = append(new_list, new_white_user)
		}
	}

	return new_list
}

//是否在白名单中
func GetWhiteUser(gameId int32, userId uint32) *WhiteUser {
	if _, ok := whiteList[gameId]; !ok {
		return nil
	}
	for _, u := range whiteList[gameId] {
		if u.UserId == userId {
			return &u
		}
	}
	return nil
}
