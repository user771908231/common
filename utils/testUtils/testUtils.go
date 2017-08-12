package testUtils

import (
	"casino_common/common/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime/debug"
)

type XiPai struct {
	Ids []int
}

func GetTestXiPai(gid int32) []int {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()
	//
	fmt.Println("开始读取配置信息")
	data, err := ioutil.ReadFile("/usr/local/gametest/" + fmt.Sprintf("%v", gid) + "/xipai.json")
	if err != nil {
		log.T("%v", err)
	}
	fmt.Println("开始解析文件:", data)

	var xipai = &XiPai{}
	err = json.Unmarshal(data, xipai)
	if err != nil {
		log.T("%v", err)
	}
	fmt.Println("读取到自定义的文件:", xipai)
	return xipai.Ids
}
