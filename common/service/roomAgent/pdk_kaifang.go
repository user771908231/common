package roomAgent

import (
	"casino_common/common/consts"
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/service/rpcService"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"errors"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

//跑得快配置
var PdkConf = CreateConfig{
	[][]string{
		[]string{"十五张跑得快", "十五张", "15张"},
		[]string{"经典跑得快", "十六张", "16张"},
	},
	[][]string{
		[]string{"20局", "二十局"},
		[]string{"10局", "十局"},
	},
	[][]string{
		[]string{"2人", "二人", "两人"},
		[]string{"3人", "三人"},
	},
	[][]string{
		[]string{"首出黑桃3", "首出黑桃三"},
		[]string{"不出黑桃3"},
	},
	[][]string{
		[]string{"不抓鸟"},
		[]string{"红桃10抓鸟", "红桃十抓鸟", "抓鸟"},
	},
	[][]string{
		[]string{"不显示余牌"},
		[]string{"显示余牌"},
	},
}

//经典跑得快开房
func DoPdkKaifang(owner uint32, group_id int32, option_str string) (error, *ddproto.CommonDeskByAgent) {
	var opt_gamer_number int = 2
	var opt_room_type ddproto.PdkEnumRoomType = ddproto.PdkEnumRoomType_PDK_T_FIFTEEN_PDK
	var opt_circle_num int = 20
	var opt_ht3_shouchu bool = false
	var opt_ht10_zhuaniao bool = false
	var opt_show_yupai bool = false

	//解析关键词
	pdk_keywords := PdkConf.GetKeywords("", option_str)
	for _, v := range pdk_keywords {
		switch v {
		case "十五张跑得快":
			opt_room_type = ddproto.PdkEnumRoomType_PDK_T_FIFTEEN_PDK
		case "经典跑得快":
			opt_room_type = ddproto.PdkEnumRoomType_PDK_T_NORMAL_PDK
		case "10局":
			opt_circle_num = 10
		case "20局":
			opt_circle_num = 20
		case "2人":
			opt_gamer_number = 2
		case "3人":
			opt_gamer_number = 3
		case "首出黑桃3":
			opt_ht3_shouchu = true
		case "红桃10抓鸟":
			opt_ht10_zhuaniao = true
		case "显示余牌":
			opt_show_yupai = true
		}
	}

	//检查是否有空闲房间
	ex_room := GetAgentFreeRoomByOption(group_id, opt_gamer_number, pdk_keywords)
	if ex_room != nil {
		return errors.New("has_ex_room"), ex_room
	}

	rpc_req := &ddproto.PdkReqCreateDesk{
		Header: &ddproto.ProtoHeader{
			UserId: proto.Uint32(owner),
		},
		RoomTypeInfo: &ddproto.PdkBaseRoomTypeInfo{
			RoomType:       opt_room_type.Enum(),
			BoardsCount:    proto.Int32(int32(opt_circle_num)),
			UserCountLimit: proto.Int32(int32(opt_gamer_number)),
			IsDaikai:       proto.Bool(true),
			IsShowCardsNum: proto.Bool(opt_show_yupai),
			IsZhuaNiao:     proto.Bool(opt_ht10_zhuaniao),
			IsSpadeThree:   proto.Bool(opt_ht3_shouchu),
		},
	}
	res, err := rpcService.GetPdk().CreateRoom(context.Background(), rpc_req)
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
		ex_room := GetAgentFreeRoomByOption(group_id, opt_gamer_number, pdk_keywords)
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
