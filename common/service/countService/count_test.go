package countService

import (
	"testing"
	"casino_common/common/service/countService/countType"
	ldb "casino_common/common/db"
)

func init() {
	ldb.InitMongoDb("192.168.199.200", 27017, "test", "id",[]string{})
}

func TestGetUserCounterByType(t *testing.T) {
	Add(1, countType.ALL_GAME_COUNT, 1)
	t.Log(GetAllCounter(1))
	t.Log(GetUserCounterByType(1, countType.ALL_GAME_COUNT))
}
