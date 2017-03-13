package statisticsService

import (
	"casino_common/common/model"
	"time"
	"casino_common/common/model/TStatisticBtnClickDao"
)

//按钮点击事件统计
/**
	按钮的点击事件是怎么处理的?
	1,按钮总共的点击事件？这个可以区分
	2,玩家自己的点击事件?如果每个玩家，点击一次，就增加一条记录....不科学啊...

 */
type BtnClickCounter interface {
	C(userId uint32, ip string, btnId int32, t time.Time) (total int32, e error)
}

//默认的处理
type DefaultBtnClickCounter struct {
	//todo 这里可以增加其他dbconnect 的实现,考虑到以后分库分表的设计
}

func NewDefaultBtnClickCounter() *DefaultBtnClickCounter {
	return &DefaultBtnClickCounter{

	}
}

//开始统计
func (b *DefaultBtnClickCounter) C(userId uint32, ip string, btnId int32, t time.Time) {
	record := &model.T_statistic_btn_click{
		UserId: userId,
		Ip:     ip,
		BtnId:  btnId,
		T:      t,
	}
	go TStatisticBtnClickDao.SaveStatisticBtnClick(record)
}
