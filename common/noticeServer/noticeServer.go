package noticeServer

import (
	"strings"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"casino_common/proto/ddproto"
	"casino_common/proto/funcsInit"
	"casino_common/common/model/noticeDao"
)

var NOTICE_TYPE_GUNDONG int32 = 1  //滚动
var NOTICE_TYPE_CHONGZHI int32 = 2 //充值信息
var NOTICE_TYPE_GONGGAO int32 = 3  //公告信息

func getRedisKey(id int32) string {
	keyPre := "public_notice_key"
	idStr, _ := numUtils.Int2String(id)
	return strings.Join([]string{keyPre, idStr}, "_")
}

//通过notice的type 来查找notice
func GetNoticeByType(noticeType int32) *ddproto.TNotice {
	//1,先从redis中获取notice
	result := redisUtils.GetObj(getRedisKey(noticeType), &ddproto.TNotice{})
	if result != nil {
		return result.(*ddproto.TNotice)
	} else {
		notice := noticeDao.FindNoticeByType(noticeType)
		if notice != nil {
			SaveNotice2Redis(notice)
		}
		return notice
	}
}

func GetCommonAckNotice(noticeType int32) *ddproto.CommonAckNotice {
	ack := commonNewPorot.NewCommonAckNotice()
	bback := GetNoticeByType(noticeType)
	if bback != nil {
		*ack.NoticeType = bback.GetNoticeType()
		*ack.NoticeTitle = bback.GetNoticeTitle()
		*ack.NoticeMemo = bback.GetNoticeMemo()
		*ack.NoticeContent = bback.GetNoticeContent()
		ack.Fileds = bback.Fileds
	}
	//返回得到的ack notice
	return ack
}

//把公告的数据保存到redis中
func SaveNotice2Redis(notice *ddproto.TNotice) error {
	redisUtils.SetObj(getRedisKey(notice.GetId()), notice)
	return nil
}

//把公告的数据保存到redis中
func SaveNotice(notice *ddproto.TNotice) error {
	noticeDao.SaveNotice2Mgo(notice)
	SaveNotice2Redis(notice)
	return nil
}
