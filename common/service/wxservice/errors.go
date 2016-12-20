package service

import "errors"

var ERR_ENDTRADE_REPETITION error = errors.New("微信异步重复回调")
