package configService

import (
	"testing"
	"reflect"
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"gopkg.in/mgo.v2/bson"
	db_init "casino_common/common/db"
	"casino_common/common/sys"
	"errors"
	login_conf "casino_login/conf"
	"time"
)

func init() {
	ConfigMap = map[string]ConfInfo {
		tableName.DBT_GAME_CONFIG_LOGIN: ConfInfo{
			Name: "登陆服配置",
			List: &[]login_conf.ConfStruct{},
			Row: &login_conf.ConfStruct{},
		},
		tableName.DBT_GAME_CONFIG_LOGIN_LIST: ConfInfo{
			Name: "登陆服游戏列表",
			List: &[]login_conf.ServerInfo{},
			Row: &login_conf.ServerInfo{},
		},
	}
	db_init.InitMongoDb("192.168.199.200", 27017, "test", "id",[]string{})
	sys.InitRedis("192.168.199.200:6379","test")
}

//列信息
type FieldInfo struct {
	Name string
	Type string
	Title string
	Info string
	Value interface{}
}

//配置信息
type ConfInfo struct {
	Name string //名称
	lastTime *time.Time  //上次更新配置的时间
	List interface{}  //表类型
	Row interface{}   //单条配置类型
}

var ConfigMap map[string]ConfInfo

//获列字段信息
func GetColInfo(col_struct interface{}) []FieldInfo {
	cols := []FieldInfo{}
	val := reflect.TypeOf(col_struct)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	val_val := reflect.ValueOf(col_struct)
	if val_val.Kind() == reflect.Ptr {
		val_val = val_val.Elem()
	}
	for i:=0;i<val.NumField();i++ {
		col_info := FieldInfo{}
		col := val.Field(i)
		col_info.Name = col.Name
		if field := col.Tag.Get("bson"); field != "" {
			col_info.Name = field
		}
		col_info.Type = col.Type.Kind().String()
		col_info.Title = col.Tag.Get("title")
		col_info.Info = col.Tag.Get("info")

		val_field := val_val.Field(i)
		if val_field.Kind() == reflect.Ptr {
			val_field = val_field.Elem()
		}
		if val_field.CanInterface() {
			col_info.Value = val_field.Interface()
		}else {
			col_info.Value = reflect.Zero(val_field.Type())
		}

		cols = append(cols, col_info)
	}

	return cols
}

//返回列表colinfo
func GetSliceField(slice_struct interface{}) [][]FieldInfo {
	val := reflect.ValueOf(slice_struct)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	cols := [][]FieldInfo{}
	if val.Type().Kind() != reflect.Slice {
		return cols
	}
	for i:=0; i<val.Len(); i++ {
		if val.Index(i).CanInterface() {
			cols = append(cols, GetColInfo(val.Index(i).Interface()))
		}
	}
	return cols
}

//从
func TestGetCol(t *testing.T) {
	conf_info := ConfigMap[tableName.DBT_T_USER]

	list := conf_info.List
	row := conf_info.Row

	err := db.C(tableName.DBT_T_USER).FindAll(bson.M{"id": 10082}, list)
	db.C(tableName.DBT_T_USER).Find(bson.M{"id": 10082}, row)

	col_info := GetColInfo(row)

	t.Log(err, list)
	t.Log(row)
	row_elem := row.(*ddproto.User)
	t.Log(row_elem.GetNickName())
	t.Log(col_info)

}

//获取配置
func GetConfig(table_name string) (conf ConfInfo, err error) {
	if conf_info,ok := ConfigMap[table_name]; ok {
		//10分钟配置缓存时间
		if conf_info.lastTime == nil || time.Now().Sub(*conf_info.lastTime) > 10 * time.Minute {
			err = db.C(table_name).FindAll(bson.M{}, conf_info.List)
			if err == nil {
				now := time.Now()
				conf_info.lastTime = &now
			}
		}

		return conf_info, err
	}else {
		return ConfInfo{}, errors.New("错误！")
	}
}

//注册配置
func Regist(tableName string, iPtr interface{}) error {
	val_conf := reflect.ValueOf(iPtr)

	conf,err := GetConfig(tableName)
	if err != nil {
		return err
	}
	conf_val := reflect.ValueOf(conf.List)

	if val_conf.Kind() != reflect.Ptr {
		return errors.New("必须传指针进来！")
	}
	val_conf.Elem().Set(conf_val.Elem())
	return nil
}



//pull config

//update config key

//getconfig

