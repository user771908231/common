package exchangeService

import (
	"casino_common/proto/ddproto"
	"gopkg.in/mgo.v2/bson"
	"time"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"casino_common/common/userService"
)

//兑换记录
type ExchangeRecord struct {
	Id          bson.ObjectId             `bson:"_id"` //编号
	UserId      uint32                    //用户id
	Type        ddproto.HallEnumTradeType //兑换商品的类型
	Amount      float64                   //兑换数量
	Name        string                    //姓名
	Phone       string                    //电话
	WxNumber    string                    //微信号码
	Address     string                    //收货地址
	IsProcess   bool                      //是否已处理
	IsAgree     bool                      //是否已同意
	RequestTime time.Time                 //申请时间
	ProcessTime time.Time                 //审核时间
}

//新兑换记录
func NewRecord(userId uint32, goods_type ddproto.HallEnumTradeType, amount float64) (*ExchangeRecord, error) {
	user := userService.GetUserById(userId)
	new_record := &ExchangeRecord{
		Type: goods_type,
		Amount: amount,
		Name: user.GetRealName(),
		Phone: user.GetPhoneNumber(),
		WxNumber: user.GetWxNumber(),
		Address: user.GetRealAddress(),
	}
	return new_record.Insert()
}

//插入新行
func (r *ExchangeRecord) Insert() (*ExchangeRecord, error) {
	r.Id = bson.NewObjectId()
	r.IsProcess = false
	r.IsAgree = false
	r.RequestTime = time.Now()
	err := db.C(tableName.DBT_ADMIN_EXCHANGE_RECORD).Insert(r)
	return r, err
}

//更新记录
func (r *ExchangeRecord) Save() error {
	return db.C(tableName.DBT_ADMIN_EXCHANGE_RECORD).Update(bson.M{
		"_id": r.Id,
	}, r)
}
