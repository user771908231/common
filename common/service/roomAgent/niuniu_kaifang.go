package roomAgent

import (
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"casino_common/common/service/rpcService"
	"casino_common/common/log"
	"golang.org/x/net/context"
	"errors"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
)

//跑得快配置
var NiuniuConf = CreateConfig{
	[][]string{
		[]string{"经典牛牛"},
		[]string{"急速牛牛"},
	},
	[][]string{
		[]string{"20局", "二十局"},
		[]string{"10局", "十局"},
	},
	[][]string{
		[]string{"6人", "六人"},
		[]string{"2人", "二人", "两人"},
		[]string{"3人", "三人"},
		[]string{"4人", "四人"},
		[]string{"5人", "五人"},
	},
	[][]string{
		[]string{"看牌抢庄", "抢庄"},
		[]string{"牛牛换庄","换庄"},
		[]string{"随机坐庄", "随机"},
		[]string{"房主定庄", "定庄"},
	},
	[][]string{
		[]string{"允许中途加入"},
		[]string{"禁止中途加入"},
	},
	[][]string{
		[]string{"有花牌"},
		[]string{"无花牌"},
	},
	[][]string{
		[]string{"花式玩法"},
		[]string{"普通玩法"},
	},
}

//经典跑得快开房
func DoNiuniuKaifang(owner uint32, group_id int32, option_str string) (error, *ddproto.CommonDeskByAgent) {
	var opt_has_animation = false
	var opt_circle_num int = 20
	var opt_gamer_number int = 6
	var opt_bank_rule = ddproto.NiuniuEnumBankerRule_QIANG_ZHUANG
	var opt_deny_half_join = false
	var opt_has_flower = false
	var opt_flower_play = false

	//解析关键词
	niuniu_keywords := PdkConf.GetKeywords("", option_str)
	for _, v := range niuniu_keywords {
		switch v {
		case "经典牛牛":
			opt_has_animation = true
		case "急速牛牛":
			opt_has_animation = false
		case "10局":
			opt_circle_num = 10
		case "20局":
			opt_circle_num = 20
		case "2人":
			opt_gamer_number = 2
		case "3人":
			opt_gamer_number = 3
		case "4人":
			opt_gamer_number = 4
		case "5人":
			opt_gamer_number = 5
		case "6人":
			opt_gamer_number = 6
		case "看牌抢庄":
			opt_bank_rule = ddproto.NiuniuEnumBankerRule_QIANG_ZHUANG
		case "牛牛换庄":
			opt_bank_rule = ddproto.NiuniuEnumBankerRule_DING_ZHUANG
		case "随机坐庄":
			opt_bank_rule = ddproto.NiuniuEnumBankerRule_SUI_JI_ZUO_ZHUANG
		case "房主定庄":
			opt_bank_rule = ddproto.NiuniuEnumBankerRule_FANGZHU_DINGZHUANG
		case "禁止中途加入":
			opt_deny_half_join = true
		case "允许中途加入":
			opt_deny_half_join = false
		case "有花牌":
			opt_has_flower = true
		case "无花牌":
			opt_has_flower = false
		case "花式玩法":
			opt_flower_play = true
		case "普通玩法":
			opt_flower_play = false
		}
	}

	//检查是否有空闲房间
	ex_room := GetAgentFreeRoomByOption(group_id, opt_gamer_number, niuniu_keywords)
	if ex_room != nil {
		return errors.New("has_ex_room"), ex_room
	}

	rpc_req := &ddproto.NiuCreateDeskReq{
		Header: &ddproto.ProtoHeader{
			UserId: proto.Uint32(owner),
		},
		Option: &ddproto.NiuniuDeskOption{
			MinUser:proto.Int32(2),
			MaxUser:proto.Int32(int32(opt_gamer_number)),
			MaxCircle:proto.Int32(int32(opt_circle_num)),
			HasFlower:proto.Bool(opt_has_flower),
			BankRule: opt_bank_rule.Enum(),
			IsFlowerPlay:proto.Bool(opt_flower_play),
			IsJiaoFenJiaBei:proto.Bool(false),
			HasAnimation:proto.Bool(opt_has_animation),
			IsCoinRoom:proto.Bool(false),
			BaseScore:proto.Int32(0),
			NeedPwd:proto.Bool(false),
			MinEnterScore:proto.Int32(0),
			MaxQzScore:proto.Int32(4),
			DenyHalfJoin:proto.Bool(opt_deny_half_join),
		},
		IsDaikai: proto.Bool(true),
	}
	res,err := rpcService.GetNiuniu().CreateRoom(context.Background(), rpc_req)
	log.T("rpc req:%v res:%v res-err:%v", rpc_req, res, err)

	if err != nil {
		log.E("rpc err:%v", err)
		return errors.New("开房失败，请联系管理员。"), nil
	}

	if res == nil {
		log.E("rpc res is nil")
		return errors.New("开房失败，请联系管理员。"), nil
	}

	if res.Header.GetCode() >= 0 {
		if group_id > 0 {
			//更新tag
			err = db.C(tableName.DBT_AGENT_CREATED_ROOM).Update(bson.M{
				"password": res.DeskState.GetDeskNumber(),
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
		ex_room := GetAgentFreeRoomByOption(group_id, opt_gamer_number, niuniu_keywords)
		if ex_room != nil {
			return nil, ex_room
		}else {
			log.E("rpc success, but ex_room is nil.")
			return errors.New("开房失败，请联系管理员。"), nil
		}
	}else {
		log.E("rpc code err:%v", res)
		return errors.New(res.Header.GetError()), nil
	}
}

