package model

import "time"

//按钮点击次数统计表
type T_statistic_btn_click struct {
	UserId uint32
	Ip     string
	BtnId  int32
	T      time.Time
}
