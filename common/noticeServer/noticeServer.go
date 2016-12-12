package noticeServer

import (
	"strings"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"casino_common/common/model"
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"casino_common/utils/db"
	"casino_common/proto/ddproto"
	"casino_common/proto/funcsInit"
)

var NOTICE_TYPE_GUNDONG int32 = 1        //滚动
var NOTICE_TYPE_CHONGZHI int32 = 2        //充值信息
var NOTICE_TYPE_GONGGAO int32 = 3        //公告信息

func getRedisKey(id int32) string {
	keyPre := "public_notice_key"
	idStr, _ := numUtils.Int2String(id)
	return strings.Join([]string{keyPre, idStr}, "_")
}


//通过notice的type 来查找notice
func GetNoticeByType(noticeType int32) *ddproto.TNotice {
	//1,先从redis中获取notice
	result := redisUtils.GetObj(getRedisKey(noticeType), &ddproto.TNotice{})
	if result == nil {
		//2,如果notice 不存在则从数据库中获取,并且保存在redis中
		notice := &model.T_th_notice{}
		db.Query(func(d *mgo.Database) {
			d.C(tableName.DBT_T_TH_NOTICE).Find(bson.M{"noticetype" : noticeType}).One(notice)
		})
		if notice.Id == 0 {
			log.T("在mongo中没有查询到notice[%v].", noticeType)
			result = nil
		} else {
			log.T("数据库中查询到noticeType[%v]的 notice[%v]", noticeType, notice)
			//把从数据获得的结果填充到redis的model中
			result = tnotice2Rnotice(notice)
			if result != nil {
				SaveNotice2Redis(result.(*ddproto.TNotice))
			}
		}
	}

	//3,返回数据
	if result == nil {
		return nil
	} else {
		return result.(*ddproto.TNotice)
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

func tnotice2Rnotice(notice *model.T_th_notice) *ddproto.TNotice {
	result := &ddproto.TNotice{}
	result.Id = new(int32)
	result.NoticeType = new(int32)
	result.NoticeTitle = new(string)
	result.NoticeContent = new(string)
	result.NoticeMemo = new(string)

	*result.Id = notice.Id
	*result.NoticeContent = notice.NoticeContent
	*result.NoticeTitle = notice.NoticeTitle
	*result.NoticeMemo = notice.NoticeMemo
	*result.NoticeType = notice.NoticeType
	result.Fileds = notice.NoticeFileds

	return result
}


//把公告的数据保存到redis中
func SaveNotice2Redis(notice *ddproto.TNotice) error {
	redisUtils.SetObj(getRedisKey(notice.GetId()), notice)
	return nil
}

//把公告的数据保存到redis中
func SaveNotice(tnotice *model.T_th_notice) error {
	saveNotice2mongos(tnotice)        //保存到mongo
	rnotice := tnotice2Rnotice(tnotice)
	SaveNotice2Redis(rnotice)
	return nil
}

//保存公告到mongo数据库
func saveNotice2mongos(tnotice *model.T_th_notice) error {
	log.T("开始保存公告信息到mogno,notice[%v]", tnotice)
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_TH_NOTICE).Insert(tnotice)
	})
	return nil
}




