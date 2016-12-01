// error.go
// created by kory 2014-02-10
//

package Error

import (
	"fmt"
	"runtime/debug"
	"casino_common/common/consts"
)

//func New(errCode int, errStr string) Error {
//	return Error{errCode, errStr}
//}

// New returns an error that formats as the given text.
func New(vlist ...interface{}) Error {
	errno := int32(0)
	errstr := ""
	for _, vv := range vlist {
		switch v := vv.(type) {
		case int32:
			{
				errno = v //vlist[0].(int)
			}
		case string:
			{
				if errno == 0 {
					errno = -1
				}
				errstr = v //vlist[0].(string)
			}
		case error:
			{
				if errno == 0 {
					errno = -1
				}
				if v != nil {
					errstr = v.Error()
				}
			}
		}
	}
	return Error{errno, errstr}
}

func NewError(errCode int32, errStr string) error {
	return &Error{errCode, errStr}
}

func NewFailError(msg string) error {
	return &Error{consts.ACK_RESULT_ERROR, msg}

}

func OK() Error {
	return Error{0, ""}
}

// errorString is a trivial implementation of error.
type Error struct {
	errCode int32
	errStr  string
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if e.errCode == 0 {
		e.errStr = "OK"
	}
	return fmt.Sprintf("[%v] %v", e.errCode, e.errStr)
}

func (e *Error) Code() int32 {
	if e == nil {
		return 0
	}

	return e.errCode
}

func (e *Error) IsError() bool {
	if e == nil {
		return false
	}

	return e.errCode != 0
}

func (e *Error) Assign(err error) Error {
	if err == nil {
		e.errCode = 0
		e.errStr = ""
	} else {
		e.errCode = -1
		e.errStr = err.Error()
	}
	return *e
}

func (e *Error) SetError(code int32, err error) {
	e.errCode = code
	e.errStr = ""
}


//返回errorCode
func GetErrorCode(e error) (ret int32) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[W]", r)
			ret = consts.ACK_RESULT_ERROR
		}
	}()
	ee := e.(*Error)
	ret = ee.errCode
	return ret
}
func GetErrorMsg(e error) (ret string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[W]", r)
			ret = ""
		}
	}()
	ee := e.(*Error)
	ret = ee.errStr
	return ret
}

func ErrorRecovery(addr string) {
	// 必须要先声明defer，否则不能捕获到panic异常
	if err := recover(); err != nil {
		fmt.Println("地址[%v]开始打印日志err[%v]", addr, err) // 这里的err其实就是panic传入的内容，55
		debug.PrintStack()
	}
}