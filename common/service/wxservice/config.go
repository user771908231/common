package service

func init() {
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
