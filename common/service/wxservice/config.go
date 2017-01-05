package service

func init() {
	WXConfig.AppId = "wx4d338762d89927b4"
	WXConfig.MchId = "1417107302"
	WXConfig.ApiKey = "DD18Secoh98GdO5Ikl1Kehj4IqPop123"
	WXConfig.WXPAY_NOTIFYURL = "http://wxpaytest.tondeen.com/wxpaycb"
}

var WXConfig struct {
	AppId           string
	MchId           string
	ApiKey          string
	WXPAY_NOTIFYURL string
}

//初始化微信
func DoInitWxpay(url string) {
	if url != "" {
		WXConfig.WXPAY_NOTIFYURL = url
	}
}
