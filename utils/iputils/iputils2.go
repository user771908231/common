package iputils

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"casino_common/common/log"
	"casino_common/common/consts"
	"casino_common/utils/redisUtils"
	"casino_common/common/model/utils"
	"strings"
	"casino_common/utils/rand"
	"fmt"
)

type BackJson struct {
	Code int //错误码 0：成功，1：失败
	Data struct {
		Region string
		City   string
	}
}

type Data struct {
	Region string
	City   string
}

//ip := "180.89.94.90"
func GetIpInfo(ip string) Data {
	//unmarshal to struct
	j := &BackJson{}
	r, e := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip) //这里有可能会导致足阻塞，所以不能用同步的方式
	if e != nil {
		log.T("查找ip对应地址的时候出错...e:%v", e)
		return Data{}
	}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal([]byte(body), &j)
	if err != nil {
		log.Printf("解析json的时候err:%v", err)
	}
	if j.Code == 1 {
		log.T("查找ip对应地址的时候出错... Code错误")
	}
	return j.Data
}

//这里取出来的地址是带逗号分隔符的 如"省,城市"
func GetIPSiteRaw(ip string) string {
	key := redisUtils.K_STRING(consts.RKEY_PRE_IP_ADDR, ip)
	ret := redisUtils.Get(key)
	if ret == "" {
		go func(ip string) {
			add := ""
			//1.查询ipInfo
			ipInfo := GetIpInfo(ip)
			if ipInfo.Region != "" && ipInfo.City != "" {
				add = ipInfo.Region + "," + ipInfo.City
			}
			//2.保存到redis
			redisUtils.Set(key, add)
			//3.保存到mongo
			utils.T_ip_address{
				Ip:      ip,
				Address: add,
			}.Insert()
		}(ip)
	}
	return ret
}

func GetIPSite2(ip string) string {
	ret := GetIPSiteRaw(ip)
	//存的字符串是有逗号分隔的 取的时候要去掉
	return strings.Join(strings.Split(ret, ","), "")
}

func GetDistance(ip1, ip2 string) string {
	ipInfo1 := strings.Split(GetIPSiteRaw(ip1), ",")
	ipInfo2 := strings.Split(GetIPSiteRaw(ip2), ",")
	if len(ipInfo1) != 2 || len(ipInfo2) != 2 {
		//其中的一个地址为空
		return ""
	}
	if ipInfo1[0] == ipInfo2[0] {
		//相同省份
		if ipInfo1[1] == ipInfo2[1] {
			//相同城市
			return fmt.Sprintf("%v", rand.Rand(5, 30))
		}
		//不同城市
		return fmt.Sprintf("%v", rand.Rand(50, 400))
	}
	//不同省份
	return fmt.Sprintf("%v", rand.Rand(500, 1000))
}
