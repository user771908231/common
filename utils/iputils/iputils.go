package iputils

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"casino_common/common/log"
	"casino_common/common/consts"
	"casino_common/utils/redisUtils"
	"casino_common/common/model/utils"
)

type Person struct {
	Data struct {
		City string
	}
}

//ip := "180.89.94.90"
func iP2site(ip string) string {
	//unmarshal to struct
	p := &Person{}
	r, e := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip) //这里有可能会导致足阻塞，所以不能用同步的方式
	if e != nil {
		log.T("查找ip对应地址的时候出错...e:%v", e)
		return ""
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		log.Printf("解析json的时候err:%v", err)
	}
	return p.Data.City
}

//通过异步的方式来更新地址
func iP2siteAsyn(ip string, cb func(add string)) {
	go func(i string, cb2 func(a string)) {
		add := iP2site(i)
		if cb2 != nil {
			cb2(add)
		}
	}(ip, cb)
}

func GetIPSite(ip string) string {
	key := redisUtils.K_STRING(consts.RKEY_PRE_IP_ADDR, ip)
	ret := redisUtils.Get(key)
	if ret == "" {
		iP2siteAsyn(ip, func(add string) {
			//保存到redis
			redisUtils.Set(key, add)
			//保存到mongo
			utils.T_ip_address{
				Ip:      ip,
				Address: add,
			}.Insert()
		})
	}
	return ret
}
