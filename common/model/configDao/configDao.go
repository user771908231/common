package configDao

import (
	"casino_common/proto/ddproto"
	"gopkg.in/mgo.v2"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_common/utils/db"
)

func GetConfigSys() *ddproto.TConfigSys {
	cfg := new(ddproto.TConfigSys)
	db.Query("", func(d *mgo.Database) {
		d.C(tableName.DBT_T_CONFIG_SYS).Find(bson.M{"id": 1}).One(cfg)
	})

	if cfg.GetId() > 0 {
		return cfg
	} else {
		return nil
	}
}

func UpsertConfigSys(data *ddproto.TConfigSys) error {
	//判断数据库中是否已经存在
	od := GetConfigSys()
	if od == nil {
		db.InsertMgoData("", tableName.DBT_T_CONFIG_SYS, data)
	} else {
		db.UpdateMgoData("", tableName.DBT_T_CONFIG_SYS, data)
	}
	return nil
}
