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

//血战到底 底分1 局数8 2番封顶 点杠炮 换三张
//四川长沙麻将配置
var SiChuanConf = CreateConfig{
	[][]string{
		[]string{"长沙麻将", "长沙"},
		[]string{"血战到底", "血战"},
		[]string{"两人两房", "2人两房", "两人2房", "2人2房", "二人两房"},
		[]string{"三人两房", "3人两房", "三人2房", "3人2房", "三人二房"},
		[]string{"四人两房", "4人两房", "四人2房", "4人2房", "四人二房"},
		[]string{"两人三房", "2人三房", "两人2房", "2人2房", "二人三房"},
		[]string{"三人三房", "3人三房", "三人3房", "3人3房"},
		[]string{"德阳麻将", "德阳"},
		[]string{"倒倒胡", "到到胡"},
		[]string{"血流成河", "血流"},
	},
	[][]string{
		[]string{"底分1", "底分一"},
		[]string{"底分2", "底分二"},
		[]string{"底分5", "底分五"},
		[]string{"底分10", "底分十"},
	},
	[][]string{
		[]string{"区分庄闲", "庄闲"},
		[]string{"不区分庄闲"},
		[]string{""},
	},
	[][]string{
		[]string{"抓0鸟", "抓零鸟"},
		[]string{"抓2鸟", "抓两鸟", "抓二鸟"},
		[]string{"抓4鸟", "抓四鸟"},
		[]string{"抓6鸟", "抓六鸟"},
		[]string{""},
	},
	[][]string{
		[]string{"4局", "四局", "局数4", "局数四"},
		[]string{"8局", "八局", "局数8", "局数八"},
		[]string{"12局", "十二局", "局数12", "局数十二"},
		[]string{"16局", "十六局", "局数16", "局数十六"},
		[]string{""},
	},
	[][]string{
		[]string{"2人", "两人", "二人"},
		[]string{"3人", "三人"},
		[]string{"4人", "四人"},
		[]string{""},
	},
	[][]string{
		[]string{"2番封顶", "两番封顶", "二番封顶"},
		[]string{"3番封顶", "三番封顶"},
		[]string{"4番封顶", "四番封顶"},
		[]string{""},
	},
	[][]string{
		[]string{"自摸加底"},
		[]string{"自摸加番"},
		[]string{""},
	},
	[][]string{
		[]string{"点杠炮"},
		[]string{"点杠花"},
		[]string{""},
	},
	[][]string{
		[]string{"幺九将对"},
		[]string{""},
	},
	[][]string{
		[]string{"天地胡"},
		[]string{""},
	},
	[][]string{
		[]string{"换三张"},
		[]string{""},
	},
	[][]string{
		[]string{"门清中张"},
		[]string{""},
	},
	[][]string{
		[]string{"卡二条"},
		[]string{""},
	},
	[][]string{
		[]string{"7张牌", "七张牌"},
		[]string{"10张牌", "十张牌"},
		[]string{"13张牌", "十三张牌"},
		[]string{""},
	},
}

//四川开房
func DoSiChuanKaifang(owner uint32, group_id int32, option_str string) (error, *ddproto.CommonDeskByAgent) {
	//必须填的参数
	//m.GetRoomTypeInfo().GetMjRoomType(),
	//m.GetRoomTypeInfo().GetBoardsCout(),
	//m.GetRoomTypeInfo().GetCapMax(),
	//m.GetRoomTypeInfo().GetCardsNum(),
	//m.GetRoomTypeInfo().GetSettlement(),
	//m.GetRoomTypeInfo().GetBaseValue(),
	//m.GetRoomTypeInfo().GetPlayOptions().GetZiMoRadio(),
	//m.GetRoomTypeInfo().GetPlayOptions().GetOthersCheckBox(),
	//m.GetRoomTypeInfo().GetPlayOptions().GetHuRadio(),
	//m.GetRoomTypeInfo().GetPlayOptions().GetDianGangHuaRadio(),
	//m.GetRoomTypeInfo().GetBoardsCout(),
	//m.GetRoomTypeInfo().GetChangShaPlayOptions()
	//PlayerCount      *int32 `protobuf:"varint,1,opt,name=playerCount" json:"playerCount,omitempty"`
	//IgnoreBank       *bool  `protobuf:"varint,2,opt,name=ignoreBank" json:"ignoreBank,omitempty"`
	//BirdCount        *int32 `protobuf:"varint,3,opt,name=birdCount" json:"birdCount,omitempty"`
	//BirdMultiple

	log.T("DoSiChuanKaifang group_id:%v, option_str:%v", group_id, option_str)

	var opt_room_type mjproto.MJRoomType = mjproto.MJRoomType_roomType_xueZhanDaoDi
	var opt_boards_count int32 = 8 //BoardsCout
	var opt_cap_max int64 = 2
	var opt_cards_num int32 = 13
	var opt_settlement int32 = 0
	var opt_base_value int64 = 1
	var opt_zimo_radio int32 = 0
	var opt_others_check_box []int32 = []int32{}
	var opt_hu_radio int32 = int32(mjproto.MJOption_ZIMO_JIA_DI)
	var opt_diangang_hu_radio int32 = 4
	var opt_player_count int32 = 4
	var opt_ignore_bank bool = false
	var opt_bird_count int32 = 0
	var opt_bird_multiple int32 = 1

	//解析关键词
	sichuan_keywords := SiChuanConf.GetKeywords("", option_str)
	for _, v := range sichuan_keywords {
		switch v {
		case "长沙麻将":
			opt_room_type = mjproto.MJRoomType_roomType_changSha
		case "血战到底":
			opt_room_type = mjproto.MJRoomType_roomType_xueZhanDaoDi
		case "两人两房":
			opt_room_type = mjproto.MJRoomType_roomType_liangRenLiangFang
		case "三人两房":
			opt_room_type = mjproto.MJRoomType_roomType_sanRenLiangFang
		case "四人两房":
			opt_room_type = mjproto.MJRoomType_roomType_siRenLiangFang
		case "两人三房":
			opt_room_type = mjproto.MJRoomType_roomType_liangRenSanFang
		case "三人三房":
			opt_room_type = mjproto.MJRoomType_roomType_sanRenSanFang
		case "德阳麻将":
			opt_room_type = mjproto.MJRoomType_roomType_deYangMaJiang
		case "倒倒胡":
			opt_room_type = mjproto.MJRoomType_roomType_daoDaoHu
		case "血流成河":
			opt_room_type = mjproto.MJRoomType_roomType_xueLiuChengHe
		case "底分1":
			opt_base_value = 1
		case "底分2":
			opt_base_value = 2
		case "底分5":
			opt_base_value = 5
		case "底分10":
			opt_base_value = 10
		case "区分庄闲":
			opt_ignore_bank = true
		case "不区分庄闲":
			opt_ignore_bank = false
		case "抓0鸟":
			opt_bird_count = 0
		case "抓2鸟":
			opt_bird_count = 2
		case "抓4鸟":
			opt_bird_count = 4
		case "抓6鸟":
			opt_bird_count = 6
		case "4局":
			opt_boards_count = 4
		case "8局":
			opt_boards_count = 8
		case "12局":
			opt_boards_count = 12
		case "16局":
			opt_boards_count = 16
		case "2人":
			opt_player_count = 2
		case "3人":
			opt_player_count = 3
		case "4人":
			opt_player_count = 4
		case "2番封顶":
			opt_cap_max = 2
		case "3番封顶":
			opt_cap_max = 3
		case "4番封顶":
			opt_cap_max = 4
		case "自摸加底":
			opt_hu_radio = int32(mjproto.HuType_H_ZiMoJiaDi)
		case "自摸加番":
			opt_hu_radio = int32(mjproto.HuType_H_ZiMoJiaFan)
		case "点杠炮":
			opt_diangang_hu_radio = 5
		case "点杠花":
			opt_diangang_hu_radio = 4
		case "幺九将对":
			opt_others_check_box = append(opt_others_check_box, int32(mjproto.MJOption_YAOJIU_JIANGDUI))
		case "天地胡":
			opt_others_check_box = append(opt_others_check_box, int32(mjproto.MJOption_TIAN_DI_HU))
		case "换三张":
			opt_others_check_box = append(opt_others_check_box, int32(mjproto.MJOption_EXCHANGE_CARDS))
		case "门清中张":
			opt_others_check_box = append(opt_others_check_box, int32(mjproto.MJOption_MENQING_MID_CARD))
		case "卡二条":
			opt_others_check_box = append(opt_others_check_box, int32(mjproto.MJOption_KA_ER_TIAO))
		case "7张牌":
			opt_cards_num = 7
		case "10张牌":
			opt_cards_num = 10
		case "13张牌":
			opt_cards_num = 13
		default:

		}
	}
	//检查是否有空闲房间
	ex_room := GetAgentFreeRoomByOption(group_id, int(opt_player_count), sichuan_keywords)
	if ex_room != nil {
		return errors.New("has_ex_room"), ex_room
	}

	rpc_req := &mjproto.Game_CreateRoom{
		Header: &mjproto.ProtoHeader{
			UserId: proto.Uint32(owner),
		},
		RoomTypeInfo: &mjproto.RoomTypeInfo{
			MjRoomType: opt_room_type.Enum(),
			BoardsCout: proto.Int32(opt_boards_count),
			CapMax:     proto.Int64(opt_cap_max),
			CardsNum:   proto.Int32(opt_cards_num),
			Settlement: proto.Int32(opt_settlement),
			BaseValue:  proto.Int64(opt_base_value),

			PlayOptions: &mjproto.PlayOptions{
				ZiMoRadio:        proto.Int32(opt_zimo_radio),
				OthersCheckBox:   opt_others_check_box,
				HuRadio:          proto.Int32(opt_hu_radio),
				DianGangHuaRadio: proto.Int32(opt_diangang_hu_radio),
			},

			ChangShaPlayOptions: &mjproto.ChangShaPlayOptions{
				PlayerCount:  proto.Int32(opt_player_count),
				IgnoreBank:   proto.Bool(opt_ignore_bank),
				BirdCount:    proto.Int32(opt_bird_count),
				BirdMultiple: proto.Int32(opt_bird_multiple),
			},
		},
		IsDaiKai: proto.Bool(true),
	}

	res, err := rpcService.GetSiChuan().CreateRoom(context.Background(), rpc_req)
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
		ex_room := GetAgentFreeRoomByOption(group_id, int(opt_player_count), sichuan_keywords)
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
