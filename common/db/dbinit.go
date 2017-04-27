package db

import (
	"fmt"
	"casino_common/utils/db"
)

func InitMongoDb(mongoIp string, DBNAME string, DB_ENSURECOUNTER_KEY string, seqTableName []string) error {
	fmt.Println("4，初始化mongo...")
	//初始化地址...
	db.Oninit(mongoIp, DBNAME, DB_ENSURECOUNTER_KEY)
	//0,活的数据库连接
	c, err := db.GetMongoConn()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//defer c.Close()

	//初始化自增键
	if len(seqTableName) > 0 {
		for _, n := range seqTableName {
			err = c.EnsureCounter(DBNAME, n, DB_ENSURECOUNTER_KEY)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
