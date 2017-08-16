package testUtils

import (
	"casino_common/common/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime/debug"
	"casino_common/common/consts"
	"casino_common/utils/encodingUtils"
	"casino_common/utils/redisUtils"
)

type XiPai struct {
	Ids []int
}

func GetTestXiPai(gid int32) []int {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()
	//
	fmt.Println("开始读取配置信息")
	data, err := ioutil.ReadFile("/usr/local/gametest/" + fmt.Sprintf("%v", gid) + "/xipai.json")
	if err != nil {
		log.T("%v", err)
	}
	fmt.Println("开始解析文件:", data)

	var xipai = &XiPai{}
	err = json.Unmarshal(data, xipai)
	if err != nil {
		log.T("%v", err)
	}
	fmt.Println("读取到自定义的文件:", xipai)
	return xipai.Ids
}

//设置牌桌预发牌数据
func SetDeskPreSendPokers(game_id int32, desk_id int32, pokers map[uint32][]int) error {
	keys := fmt.Sprintf("%s_%d_%d", consts.RKEY_DESK_PRE_XIPAI_DATA, game_id, desk_id)

	data, err := encodingUtils.GobEncode(pokers)
	if err != nil {
		return nil
	}

	return redisUtils.SetBytes(keys, data)
}

//获取牌桌预发牌数据(调用此函数返回值后会立即从redis中删除该条数据)
func GetDeskPreSendPokers(game_id int32, desk_id int32) (pokers map[uint32][]int,err error) {
	keys := fmt.Sprintf("%s_%d_%d", consts.RKEY_DESK_PRE_XIPAI_DATA, game_id, desk_id)

	data, err := redisUtils.GetBytes(keys)
	if err != nil {
		return
	}

	err = encodingUtils.GobDecode(data, &pokers)
	//立即删除redis数据,以免数据未被及时删除
	redisUtils.Del(keys)
	return
}
