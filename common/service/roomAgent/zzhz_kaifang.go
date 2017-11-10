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
var ZzhzConf = CreateConfig{
	[][]string{
		[]string{"红中麻将", "红中"},
		[]string{"转转麻将", "转转"},
	},
	[][]string{
		[]string{"8局", "八局"},
		[]string{"16局", "十六局"},
	},
	[][]string{
		[]string{"4人", "四人"},
		[]string{"3人", "三人"},
		[]string{"2人", "二人", "两人"},
	},
	[][]string{
		[]string{"鸟不翻倍", "不翻倍", "不翻"},
		[]string{"抓鸟翻倍", "翻倍"},
	},
	[][]string{
		[]string{"一码全中", "全中", "一码"},
		[]string{"抓2鸟", "抓二鸟", "2鸟", "二鸟"},
		[]string{"抓4鸟", "抓四鸟", "4鸟", "四鸟"},
		[]string{"抓6鸟", "抓六鸟", "6鸟", "六鸟"},
	},
	[][]string{
		[]string{"金鸟"},
		[]string{""},
	},
	[][]string{
		[]string{"可接炮"},
		[]string{""},
	},
	[][]string{
		[]string{"飘分"},
		[]string{""},
	}, [][]string{
		[]string{"自由飘分"},
		[]string{""},
	},
	[][]string{
		[]string{"8红中", "八红中"},
		[]string{""},
	},
}

//红中开房
func DoZzHzKaifang(owner uint32, group_id int32, option_str string) (error, *ddproto.CommonDeskByAgent) {
	var opt_gamer_number int = 4
	var opt_room_type mjproto.MJRoomType = mjproto.MJRoomType_roomType_mj_hongzhong
	var opt_circle_num int = 8
	var opt_zhama_jiabei bool = false
	var opt_bird_num int = 1
	var opt_is_gold_bird = false
	var opt_is_piao = false
	var opt_is_piao_free = false
	var opt_can_jiepao = false
	var opt_is_bahongzhong = false

	//解析关键词
	zzhz_keywords := ZzhzConf.GetKeywords("", option_str)
	for _, v := range zzhz_keywords {
		switch v {
		case "红中麻将":
			opt_room_type = mjproto.MJRoomType_roomType_mj_hongzhong
		case "转转麻将":
			opt_room_type = mjproto.MJRoomType_roomType_mj_zhuanzhuan
		case "8局":
			opt_circle_num = 8
		case "16局":
			opt_circle_num = 16
		case "3人":
			opt_gamer_number = 3
		case "2人":
			opt_gamer_number = 2
		case "4人":
			opt_gamer_number = 4
		case "鸟不翻倍":
			opt_zhama_jiabei = false
		case "抓鸟翻倍":
			opt_zhama_jiabei = true
		case "一码全中":
			opt_bird_num = 1
		case "抓2鸟":
			opt_bird_num = 2
		case "抓4鸟":
			opt_bird_num = 4
		case "抓6鸟":
			opt_bird_num = 6
		case "金鸟":
			opt_is_gold_bird = true
		case "飘分":
			opt_is_piao = true
		case "自由飘分":
			opt_is_piao_free = true
		case "可接炮":
			opt_can_jiepao = true
		case "8红中":
			opt_is_bahongzhong = true
		}
	}
	//检查是否有空闲房间
	ex_room := GetAgentFreeRoomByOption(group_id, opt_gamer_number, zzhz_keywords)
	if ex_room != nil {
		return errors.New("has_ex_room"), ex_room
	}

	rpc_req := &mjproto.Game_CreateRoom{
		Header: &mjproto.ProtoHeader{
			UserId: proto.Uint32(owner),
		},
		RoomTypeInfo: &mjproto.RoomTypeInfo{
			MjRoomType: opt_room_type.Enum(),
			BoardsCout: proto.Int32(int32(opt_circle_num)),
			BaseValue:  proto.Int64(1),
			ChangShaPlayOptions: &mjproto.ChangShaPlayOptions{
				PlayerCount: proto.Int32(int32(opt_gamer_number)),
			},
			ZhuanzhuanPlayOptions: &mjproto.ZhuanZhuanPlayOptions{
				ZhaMa:         proto.Int32(int32(opt_bird_num)),
				IsZhaMaJiaBei: proto.Bool(opt_zhama_jiabei),
				IsGoldBird:    proto.Bool(opt_is_gold_bird),
				CanJiePaoHu:   proto.Bool(opt_can_jiepao),
				IsBaHongZhong: proto.Bool(opt_is_bahongzhong),
			},
			IsPiaoFen:     proto.Bool(opt_is_piao),
			IsPiaoFenFree: proto.Bool(opt_is_piao_free),
		},
		IsDaiKai: proto.Bool(true),
	}

	res, err := rpcService.GetZzHz().CreateRoom(context.Background(), rpc_req)
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
		ex_room := GetAgentFreeRoomByOption(group_id, opt_gamer_number, zzhz_keywords)
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
