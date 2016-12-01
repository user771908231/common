//author saplmm

package db

import (
	"github.com/name5566/leaf/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/name5566/leaf/log"
	"casino_common/common/model"
)

var mongoConfig struct {
	ip                   string
	port                 int
	dbname               string
	DB_ENSURECOUNTER_KEY string
}

func Oninit(ip string, port int, dbname string, key string) {
	log.Debug("初始化mongoDb的地址  ip[%v],port[%v]", ip, port)
	mongoConfig.ip = ip
	mongoConfig.port = port
	mongoConfig.dbname = dbname
	mongoConfig.DB_ENSURECOUNTER_KEY = key
}

//活的链接
func GetMongoConn() (*mongodb.DialContext, error) {
	return mongodb.Dial(mongoConfig.ip, mongoConfig.port)
}


//保存数据
func InsertMgoData(dbt string, data interface{}) error {
	//得到连接
	c, err := GetMongoConn()
	if err != nil {
		return err
	}
	defer c.Close()

	// 获取回话 session
	s := c.Ref()
	defer c.UnRef(s)

	error := s.DB(mongoConfig.dbname).C(dbt).Insert(data)
	return error

}

//更新数据通过_id来更新
func UpdateMgoData(dbt string, data model.BaseMode) error {
	c, err := GetMongoConn()
	if err != nil {
		return err
	}
	defer c.Close()

	// 获取回话 session
	s := c.Ref()
	defer c.UnRef(s)

	error := s.DB(mongoConfig.dbname).C(dbt).Update(bson.M{"_id": data.GetMid()}, data)
	return error
}


//得到序列号
func GetNextSeq(dbt string) (int32, error) {
	//连接数据库
	c, err := GetMongoConn()
	if err != nil {
		return 0, err
	}
	defer c.Close()

	//获取session
	s := c.Ref()
	defer c.UnRef(s)
	id, _ := c.NextSeq(mongoConfig.dbname, dbt, mongoConfig.DB_ENSURECOUNTER_KEY)
	return int32(id), nil
}


//查询一个list
func Query(f func(*mgo.Database)) {
	c, err := GetMongoConn()
	if err != nil {
	}
	defer c.Close()

	//获取session
	s := c.Ref()
	defer c.UnRef(s)
	f(s.DB(mongoConfig.dbname))
}