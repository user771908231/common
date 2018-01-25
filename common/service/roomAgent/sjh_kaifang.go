package roomAgent

import (
	"casino_common/common/consts"
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/service/rpcService"
	"casino_common/proto/ddproto"
	"casino_common/proto/ddproto/mjproto"
	"casino_common/utils/db"
	"errors"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

//转转红中配置
var SjhConf = CreateConfig{
	[][]string{
		[]string{"2圈", "二圈"},
		[]string{"4圈", "四圈"},
		[]string{"8圈", "八圈"},
	},
	[][]string{
		[]string{"点炮包三家", "包三家"},
		[]string{"坐车也输分"},
		[]string{"点炮自己拿"},
	},
	[][]string{
		[]string{"16封顶", "十六封顶"},
		[]string{"32封顶", "三十二封顶"},
	},
}

//松江河开房
func DoSjhKaifang(owner uint32, group_id int32, option_str string) (error, *ddproto.CommonDeskByAgent) {
	var opt_gamer_number int = 4
	var opt_room_type mjproto.MJRoomType = mjproto.MJRoomType_roomType_mj_songjianghe
	var opt_circle_num int32 = 2
	var opt_cap_max int64 = 16
	var opt_payment int32 = 1
	var opt_is_dianpaobaofen bool = false
	//解析关键词
	sjh_keywords := SjhConf.GetKeywords("", option_str)
	for _, v := range sjh_keywords {
		switch v {
		case "2圈":
			opt_circle_num = 2
		case "4圈":
			opt_circle_num = 4
		case "8圈":
			opt_circle_num = 8
		case "点炮包三家": //1
			opt_is_dianpaobaofen = true
			opt_payment = 1
		case "坐车也输分": //2
			opt_is_dianpaobaofen = false
			opt_payment = 2
		case "点炮自己拿": //3
			opt_payment = 3
		case "16封顶":
			opt_cap_max = 16
		case "32封顶":
			opt_cap_max = 32
		}
	}
	//检查是否有空闲房间
	ex_room := GetAgentFreeRoomByOption(group_id, opt_gamer_number, sjh_keywords)
	if ex_room != nil {
		return errors.New("has_ex_room"), ex_room
	}

	rpc_req := &mjproto.Game_CreateRoom{
		Header: &mjproto.ProtoHeader{
			UserId: proto.Uint32(owner),
		},
		RoomTypeInfo: &mjproto.RoomTypeInfo{
			MjRoomType: opt_room_type.Enum(),
			BoardsCout: proto.Int32(opt_circle_num),
			BaiShanPlayOptions: &mjproto.BaiShanPlayOptions{
				DianPaoBaoFen: proto.Bool(opt_is_dianpaobaofen),
				PaymentOption: proto.Int32(opt_payment),
			},
			CapMax: proto.Int64(opt_cap_max),
		},
		IsDaiKai: proto.Bool(true),
	}

	res, err := rpcService.GetSjhMj().CreateRoom(context.Background(), rpc_req)
	log.T("rpc req:%v res:%v res-err:%v", rpc_req, res, err)

	if err != nil {
		log.E("rpc err:%v", err)
		return errors.New("开房失败，请联系管理员。"), nil
	}

	if res == nil {
		log.E("rpc res is nil")
		return errors.New("开房失败，请联系管理员。"), nil
	}

	if res.Header.GetCode() == consts.ACK_RESULT_SUCC {
		if group_id > 0 {
			//更新tag
			err = db.C(tableName.DBT_AGENT_CREATED_ROOM).Update(bson.M{
				"password": res.GetPassword(),
			}, bson.M{
				"$set": bson.M{"groupid": group_id},
			})
			if err != nil {
				log.E("update groupid err:%v", err)
				return err, nil
			}
			//更新redis
			saveToRedisByGroupid(group_id)
		}
		ex_room := GetAgentFreeRoomByOption(group_id, opt_gamer_number, sjh_keywords)
		if ex_room != nil {
			return nil, ex_room
		} else {
			log.E("rpc success, but ex_room is nil.")
			return errors.New("开房失败，请联系管理员。"), nil
		}
	} else {
		log.E("rpc code err:%v", res)
		return errors.New(res.Header.GetError()), nil
	}
}
