package noticeServer

import (
	"casino_common/common/log"
	"casino_common/common/model/noticeDao"
	"casino_common/proto/ddproto"
	"casino_common/proto/funcsInit"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"github.com/golang/protobuf/proto"
	"strings"
)

var NOTICE_TYPE_GUNDONG int32 = 1  //滚动
var NOTICE_TYPE_CHONGZHI int32 = 2 //充值信息
var NOTICE_TYPE_GONGGAO int32 = 3  //公告信息

func getRedisKey(id int32, channelId string) string {
	keyPre := "public_notice_key"
	idStr, _ := numUtils.Int2String(id)
	return strings.Join([]string{keyPre, idStr, channelId}, "_")
}

//通过notice的type 来查找notice,邮件
func GetNoticeByType(noticeType int32, channelId string) *ddproto.TNotice {
	//1,先从redis中获取notice
	result := redisUtils.GetObj(getRedisKey(noticeType, channelId), &ddproto.TNotice{})
	if result != nil {
		return result.(*ddproto.TNotice)
	} else {
		notice := noticeDao.FindNoticeByType(noticeType, channelId)
		if notice != nil {
			SaveNotice2Redis(notice)
		}
		return notice
	}
}

func GetCommonAckNotice(noticeType int32, channelId string) *ddproto.CommonAckNotice {
	ack := commonNewPorot.NewCommonAckNotice()
	//bback := GetNoticeByType(noticeType, channelId)
	//todo 临时处理，直接从数据库取数据
	bback := noticeDao.FindNoticeByType(noticeType, channelId) //手下尝试直接从数据库取出来
	log.T("noticeServer开始查询notice: notietype=%v,channelid=%v,ret:%v", noticeType, channelId, bback)

	if bback == nil {
		//查找默认值
		bback = noticeDao.FindNoticeByType(noticeType, "") //手下尝试直接从数据库取出来
	}

	if bback != nil {
		*ack.NoticeType = bback.GetNoticeType()
		*ack.NoticeTitle = bback.GetNoticeTitle()
		*ack.NoticeMemo = bback.GetNoticeMemo()
		*ack.NoticeContent = bback.GetNoticeContent()
		ack.Fileds = bback.GetNoticefileds()
		ack.Id = proto.Int32(bback.GetId())
	}

	//返回得到的ack notice
	return ack
}

//把公告的数据保存到redis中
func SaveNotice2Redis(notice *ddproto.TNotice) error {
	redisUtils.SetObj(getRedisKey(notice.GetId(), notice.GetChannelId()), notice)
	return nil
}

//把公告的数据保存到redis中
func SaveNotice(notice *ddproto.TNotice) error {
	noticeDao.SaveNotice2Mgo(notice)
	SaveNotice2Redis(notice)
	return nil
}
