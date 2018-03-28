package utils

import (
	"casino_common/common/Error"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
)

type T_ip_address struct {
	Ip      string
	Address string
}

func (t T_ip_address) Insert() {
	//保存数据
	go func(d *T_ip_address) {
		Error.ErrorRecovery("保存ip地址表")
		db.InsertMgoData(tableName.DBT_IP_ADDRESS, d)
	}(&t)
}
