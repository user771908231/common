package service

func init() {
	WXConfig.AppId = "wx4d338762d89927b4"
	WXConfig.MchId = "1417107302"
	WXConfig.ApiKey = "DD18Secoh98GdO5Ikl1Kehj4IqPop123"

}

var WXConfig struct {
	AppId  string
	MchId  string
	ApiKey string
}

const (
	DEFAULT_DEVICEINFO string = "WEB" //默认web
	//WXPAY_NOTIFYURL    string = "http://wxpaytest.tondeen.com/wxpaycb"
	WXPAY_NOTIFYURL string = "https://www.dongzhile.cn/wxpaycb"
)
