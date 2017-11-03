package aipVerifyUtils

import (
	"github.com/Pallinder/go-iap"
	"casino_common/common/log"
	"time"
)

//内部调用
func verifyReceipt(receipt_code string) (receipt *goiap.Receipt, err error) {
	receipt, err = goiap.VerifyReceipt(receipt_code,false)

	goiapErr, ok := err.(goiap.ErrorWithCode)

	//沙盒测试处理
	if ok && goiapErr.Code() == goiap.SandboxReceiptOnProd {
		receipt, err = goiap.VerifyReceipt(receipt_code, true)
	}

	if err != nil {
		return nil, err
	}

	return
}

//外部调用，多次验证，防止苹果服务器异常导致验证失败
func VerifyReceipt(receipt_code string) (receipt *goiap.Receipt, err error) {
	for i:=0;i<3;i++{
		receipt, err = verifyReceipt(receipt_code)
		if err != nil {
			//验证错误
			log.E("AppleIAP Verify Fail No.%d. receipt_code:%v receipt:%v err:%v", i+1, receipt_code, receipt, err)
			//延时再次验证
			<-time.After(1*time.Second)
		}else {
			//验证成功
			log.E("AppleIAP Verify Success No.%d. receipt_code:%v receipt:%v err:%v", i+1, receipt_code, receipt, err)
			//直接返回
			return
		}
	}

	return
}
