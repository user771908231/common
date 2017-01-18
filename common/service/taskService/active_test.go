package taskService

import (
	"testing"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
	ldb "casino_common/common/db"
)

func init() {
	ldb.InitMongoDb("192.168.199.200", 27017, "test", "id",[]string{})
}

func TestActiveInsert(t *testing.T) {
	list := []*ddproto.HallItemEvent{
		&ddproto.HallItemEvent{
			Id:      proto.Int32(0),
			Type:    ddproto.HallEnumEvent_TYPE_TIME.Enum(),
			Reward:  ddproto.HallEnum_Reward_RE_EXP.Enum(),
			RichText:[]string{" 神经游戏为广大玩家提供休闲娱乐的平台，旗下所属的游戏均为模拟竞技类游戏，我们希望每一位玩家在此能感受到竞技的乐趣。游戏中使用的游戏金币为游戏道具，不具有任何的实际价值，只限玩家本人在游戏中使用。", "" },
			Title:   proto.String("重要公告"),
		},
		&ddproto.HallItemEvent{
			Id:      proto.Int32(1),
			Type:    ddproto.HallEnumEvent_TYPE_TIME.Enum(),
			Reward:  ddproto.HallEnum_Reward_RE_EXP.Enum(),
			RichText:[]string{"神经游戏迎新年，丰厚豪礼送不停！每日登录领取金币奖励，多玩游戏即可获得更多丰厚奖励！", ""},
			Title:   proto.String("迎新春"),
		},
		&ddproto.HallItemEvent{
			Id:      proto.Int32(2),
			Type:    ddproto.HallEnumEvent_TYPE_TIME.Enum(),
			Reward:  ddproto.HallEnum_Reward_RE_EXP.Enum(),
			RichText:[]string{"活邀请好友注册游戏，可获金币奖励！好友完游戏还可获得更多奖励，钻石、房卡、奖券送不停！", ""},
			Title:   proto.String("邀请奖励"),
		},
		&ddproto.HallItemEvent{
			Id:      proto.Int32(3),
			Type:    ddproto.HallEnumEvent_TYPE_TIME.Enum(),
			Reward:  ddproto.HallEnum_Reward_RE_EXP.Enum(),
			RichText:[]string{"登录即可抽取幸运转盘大奖连续登录7天更有惊喜大奖！奖品不限量，人人有份拿！", ""},
			Title:   proto.String("连续登录"),
		},
	}
	for _,task := range list {
		var err error = nil
		db.Query(func(d *mgo.Database) {
			err = d.C(tableName.DBT_T_ACTIVE_LIST).Insert(task)
		})
		t.Log(err)
	}
}

//测试：活动列表
func TestGetActiveList(t *testing.T) {
	t.Log(GetActiveList())
}
