//author saplmm

package db

import (
	"github.com/name5566/leaf/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/name5566/leaf/log"
	"reflect"
	"errors"
)

type BaseMode interface {
	GetId() int32
}

type BaseModeu32 interface {
	GetId() uint32
}


var mongoConfig struct {
	ip                   string
	port                 int
	dbname               string
	DB_ENSURECOUNTER_KEY string
	dialc                *mongodb.DialContext
}

func Oninit(ip string, port int, dbname string, key string) {
	log.Debug("初始化mongoDb的地址  ip[%v],port[%v]", ip, port)
	mongoConfig.ip = ip
	mongoConfig.port = port
	mongoConfig.dbname = dbname
	mongoConfig.DB_ENSURECOUNTER_KEY = key
	mongoConfig.dialc, _ = mongodb.Dial(mongoConfig.ip, mongoConfig.port)
}

//活的链接
func GetMongoConn() (*mongodb.DialContext, error) {
	//return mongodb.Dial(mongoConfig.ip, mongoConfig.port)
	if mongoConfig.dialc == nil {
		mongoConfig.dialc, _ = mongodb.Dial(mongoConfig.ip, mongoConfig.port)
	}
	return mongoConfig.dialc, nil
}

//保存数据
func InsertMgoData(dbt string, data interface{}) error {
	//得到连接
	c, err := GetMongoConn()
	if err != nil {
		return err
	}
	//defer c.Close()

	// 获取回话 session
	s := c.Ref()
	defer c.UnRef(s)

	error := s.DB(mongoConfig.dbname).C(dbt).Insert(data)
	return error
}

//批量保存数据
func InsertMgoDatas(dbt string, datas []interface{}) (err error, count int) {
	c, err := GetMongoConn()
	if err != nil {
		return err, -1
	}
	//defer c.Close()

	// 获取回话 session
	s := c.Ref()
	defer c.UnRef(s)

	for _, data := range datas {

		err = s.DB(mongoConfig.dbname).C(dbt).Insert(data)
		if err == nil {
			count++
		}
	}
	return nil, count
}

//更新数据通过_id来更新
func UpdateMgoData(dbt string, data BaseMode) error {
	c, err := GetMongoConn()
	if err != nil {
		return err
	}
	//defer c.Close()

	// 获取回话 session
	s := c.Ref()
	defer c.UnRef(s)

	error := s.DB(mongoConfig.dbname).C(dbt).Update(bson.M{"id": data.GetId()}, data)
	return error
}

func UpdateMgoDataU32(dbt string, data BaseModeu32) error {
	c, err := GetMongoConn()
	if err != nil {
		return err
	}
	//defer c.Close()

	// 获取回话 session
	s := c.Ref()
	defer c.UnRef(s)

	error := s.DB(mongoConfig.dbname).C(dbt).Update(bson.M{"id": data.GetId()}, data)
	return error
}

//得到序列号
func GetNextSeq(dbt string) (int32, error) {
	//连接数据库
	c, err := GetMongoConn()
	if err != nil {
		return 0, err
	}
	//defer c.Close()

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
	//defer c.Close()

	//获取session
	s := c.Ref()
	defer c.UnRef(s)
	f(s.DB(mongoConfig.dbname))
}

//通过k==v查询一条数据
func FindByKV(dbt string, key string, value interface{}, obj interface{}) error {
	var err error = nil
	Query(func(mgo *mgo.Database) {
		err = mgo.C(dbt).Find(bson.M{
			key: value,
		}).One(obj)
	})
	return err
}

//通过k==v查询多条数据
func SelectByKV(dbt string, key string, value interface{}, obj interface{}) error {
	var err error = nil
	Query(func(mgo *mgo.Database) {
		err = mgo.C(dbt).Find(bson.M{
			key: value,
		}).All(obj)
	})
	return err
}

func CloseMGO() {
	log.Release("mgo close...")
	mongoConfig.dialc.Close()
}

//模拟类似ThinkPHP的M()方法
type C string

//查询单条
func (c C) Find(query interface{}, result interface{}) error {
	var err error = nil
	Query(func(mgo *mgo.Database) {
		err = mgo.C(string(c)).Find(query).One(result)
	})
	return err
}

//查询多条
func (c C) FindAll(query interface{}, result interface{}) error {
	var err error = nil
	Query(func(mgo *mgo.Database) {
		err = mgo.C(string(c)).Find(query).All(result)
	})
	return err
}

//查询分页
func (c C) Page(query interface{}, result interface{}, sort string, page int, limit int) (err error, count int) {
	Query(func(mgo *mgo.Database) {
		q := mgo.C(string(c)).Find(query)
		//统计数量
		count, err = q.Count()

		if sort != "" {
			q = q.Sort(sort)
		}
		if limit > 0 && page >= 1 {
			q = q.Skip((page - 1) * limit).Limit(limit)
		}
		err = q.All(result)
	})
	return err, count
}

//更新一条
func (c C) Update(query interface{}, update interface{}) error {
	var err error = nil
	Query(func(mgo *mgo.Database) {
		err = mgo.C(string(c)).Update(query, update)
	})
	return err
}

//更新多条
func (c C) UpdateAll(query interface{}, update interface{}) (change_info *mgo.ChangeInfo, err error) {
	Query(func(mgo *mgo.Database) {
		change_info, err = mgo.C(string(c)).UpdateAll(query, update)
	})
	return
}

//插入一条或多条
func (c C) Insert(doc ...interface{}) error {
	var err error = nil
	Query(func(mgo *mgo.Database) {
		err = mgo.C(string(c)).Insert(doc...)
	})
	return err
}

//一次性插入多条数据
func (c C) InsertAll(doc_slice interface{}) (error, int) {
	if reflect.TypeOf(doc_slice).Kind() != reflect.Slice {
		return errors.New("doc_slice must be a slice."), 0
	}
	val := reflect.ValueOf(doc_slice)
	doc_arr := []interface{}{}
	doc_len := val.Len()
	for i:=0; i<doc_len; i++ {
		doc_arr = append(doc_arr, val.Index(i).Interface())
	}
	return c.Insert(doc_arr...), doc_len
}

//删除单条数据
func (c C) Remove(query interface{}) error {
	var err error = nil
	Query(func(mgo *mgo.Database) {
		err = mgo.C(string(c)).Remove(query)
	})
	return err
}

//删除多条数据
func (c C) RemoveAll(query interface{}) (change_info *mgo.ChangeInfo, err error) {
	Query(func(mgo *mgo.Database) {
		change_info, err = mgo.C(string(c)).RemoveAll(query)
	})
	return
}

//计数
func (c C) Count(query interface{}) (count int, err error) {
	Query(func(mgo *mgo.Database) {
		count, err = mgo.C(string(c)).Find(query).Count()
	})
	return
}

//管道
func (c C)PipeAll(query interface{}, result interface{}) (err error) {
	Query(func(mgo *mgo.Database) {
		pipe := mgo.C(string(c)).Pipe(query)
		err = pipe.All(result)
	})
	return
}
//管道
func (c C)Pipe(query interface{}, result interface{}) (err error) {
	Query(func(mgo *mgo.Database) {
		pipe := mgo.C(string(c)).Pipe(query)
		err = pipe.One(result)
	})
	return
}

//删除表
func (c C) Drop() (err error) {
	Query(func(mgo *mgo.Database) {
		err = mgo.C(string(c)).DropCollection()
	})
	return
}
