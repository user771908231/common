package TStatisticBtnClickDao

import (
	"casino_common/common/Error"
	"casino_common/common/consts/tableName"
	"casino_common/common/model"
	"casino_common/utils/db"
)

//
func SaveStatisticBtnClick(click *model.T_statistic_btn_click) error {
	defer Error.ErrorRecovery("保存点击事件")
	return db.InsertMgoData(tableName.DBT_STATISTIC_BTN_CLICK, click)
}
