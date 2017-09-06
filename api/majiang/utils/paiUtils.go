package utils

import (
	"casino_common/api/majiang"
	"github.com/golang/protobuf/proto"
	"fmt"
	"casino_common/api/majiang/api"
	"casino_common/proto/ddproto/mjproto"
)

func Mjpai2Card(p *majiang.MJPAI) *mjproto.CardInfo {
	//判断空
	if p == nil {
		return &mjproto.CardInfo{}
	}

	//返回对应的cardInfo
	return &mjproto.CardInfo{
		Type:  proto.Int32(int32(p.Flower)),
		Value: proto.Int32(p.GetClientId()),
		Id:    proto.Int32(p.Index),
	}
}

func ListMjapi2Card(ps []*majiang.MJPAI) []*mjproto.CardInfo {
	//判断空
	if len(ps) <= 0 {
		return nil
	}

	var ret []*mjproto.CardInfo
	for _, p := range ps {
		if p != nil {
			ret = append(ret, &mjproto.CardInfo{
				Type:  proto.Int32(int32(p.Flower)),
				Value: proto.Int32(p.GetClientId()),
				Id:    proto.Int32(p.Index),
			})
		}
	}
	//返回对应的cardInfo
	return ret
}


func IsXia(p *majiang.MJPAI) bool {
	if p == nil {
		return false
	}
	return p.Flower == majiang.MJ_FLOWER_W
}

//判断一张牌是不是红中
func IsHongZhong(p *majiang.MJPAI) bool {
	if p == nil {
		return false
	}
	return p.GetClientId() == 32 //32是红中的client id
}

func List2Str(pais []*majiang.MJPAI) string {
	if len(pais) <= 0 {
		return ""
	}
	s := ""
	for _, p := range pais {
		s = s + p.LogDes() + " "
	}
	//返回值
	return s
}

func ListGameData2Str(g api.MJUserGameData) string {
	/**
	u.GameData.HandPais,
		u.GameData.GangPais,
		u.GameData.PengPais,
		u.GameData.ChiPais,
		u.GameData.HuInfo,
		u.GameData.MoPai,
	 */
	return fmt.Sprintf("手牌:%v ,杠牌:%v ,碰牌:%v,吃牌:%v,胡牌:%v摸牌:%v",
		List2Str(g.GetHandPais()),
		ListGang2Str(g.GetGangPais()),
		List2Str(ListPeng2Str(g.GetPengPais())),
		"--",
		"--",
		g.GetMoPai().LogDes())
}

func ListPeng2Str(pps []*majiang.PengPai) []*majiang.MJPAI {
	var ps []*majiang.MJPAI
	for _, p := range pps {
		ps = append(ps, p.Pais...)
	}
	return ps
}
func ListGang2Str(pps []*majiang.GangPai) []*majiang.MJPAI {
	var ps []*majiang.MJPAI
	for _, p := range pps {
		ps = append(ps, p.Pais...)
	}
	return ps
}

func ListChi2Str(pps []*majiang.ChiPai) []*majiang.MJPAI {
	return nil
}
