package iputils

import (
	"casino_common/common/log"
	"casino_common/common/sys"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func init() {
	sys.InitRedis("localhost", "test")
}

func TestIP2site(t *testing.T) {
	add := iP2site("118.112.57.224")
	t.Logf("得到的地址:%v", add)
}

func TestIP2siteAsyn(t *testing.T) {
	log.Printf("开始获取地址\n")
	iP2siteAsyn("118.112.57.224", func(add string) {
		log.Printf("这里是异步调用:%v \n ", add)
	})
	log.Printf("开始等待\n")
	time.Sleep(time.Second * 10)
	log.Printf("end\n")
}

func TestGetIPSite(t *testing.T) {
	log.Printf("%v \n", GetIPSite("118.112.57.224"))
}

func TestGetLocationByLatitudeAndLongitude(t *testing.T) {
	var Latitude float64 = 104.06407928466797
	var Longitude float64 = 30.548797607421875
	s := GetLocationByLatitudeAndLongitude(Latitude, Longitude)
	t.Logf("查询得到的地址:%v", s)

	type Person struct {
		Result struct {
			Formatted_address string
		}
	}
	//todo 这里是随便找的一个key
	ak := "DD279b2a90afdf0ae7a3796787a0742e"
	p := &Person{}
	url := "http://api.map.baidu.com/geocoder/v2/?ak=" + ak + "&location=" + fmt.Sprintf("%v", Longitude) + "," + fmt.Sprintf("%v", Latitude) + "&output=json&pois=0"
	r, err := http.Get(url)
	t.Logf("访问的url:%v", url)
	if err != nil {
		t.Logf("出错:%v", err)
	}

	//读取内容
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal([]byte(body), &p)
	if err != nil {
		t.Logf("解析json的时候err:%v", err)
	}
	t.Logf("得到的ret :%v", p)
}
