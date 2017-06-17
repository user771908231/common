package sessionService

import (
	"casino_common/test"
	"testing"
	"casino_common/utils/redisUtils"
	"casino_common/common/consts"
	"casino_common/utils/numUtils"
)

func init() {
	test.TestInit()
}

func TestGetSession(t *testing.T) {
	userId := uint32(10404)
	//s := GetSession(10507, 2)
	fs := GetSession(userId, 1)
	t.Logf("得到的session： %v", fs)
	cs := GetSession(userId, 2)
	t.Logf("得到的session： %v", cs)
}

func TestIsOnMaintain(t *testing.T) {
	var gid int32 = 8
	redisUtils.SetInt64(consts.ADMIN_GAMEID_ISMAINTAIN+"_"+numUtils.Int2String2(gid), 0)
	t.Log(IsOnMaintain(gid))
	redisUtils.SetInt64(consts.ADMIN_GAMEID_ISMAINTAIN+"_"+numUtils.Int2String2(gid), 1)
	t.Log(IsOnMaintain(gid))
	redisUtils.SetInt64(consts.ADMIN_GAMEID_ISMAINTAIN+"_"+numUtils.Int2String2(gid), 0)
	t.Log(IsOnMaintain(gid))
}
