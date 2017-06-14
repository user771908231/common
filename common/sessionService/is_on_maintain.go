package sessionService

import (
	"casino_common/utils/redisUtils"
	"casino_common/common/consts"
	"casino_common/utils/numUtils"
)

//是否处于维护中,第二个返回值为维护中的提示消息
func IsOnMaintain(gid int32) (bool, string) {
	gid_str := numUtils.Int2String2(gid)
	is_on := redisUtils.GetInt64(redisUtils.K_STRING(consts.ADMIN_GAMEID_ISMAINTAIN, gid_str))
	title := redisUtils.Get(redisUtils.K_STRING(consts.ADMIN_GAMEID_MAINTAIN_INFO, gid_str))

	if is_on == 1 {
		return true, title
	}

	return false, ""
}
