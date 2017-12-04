package ignoreSensitiveWordsUtils

import (
	"casino_common/common/log"
	"fmt"
	"github.com/importcjj/sensitive"
)

var Filter *sensitive.Filter

func init() {
	Filter = sensitive.New()
	if err := Filter.LoadWordDict("../conf/dict.txt"); err != nil {
		log.E("过滤敏感词加载文件错误：%v", err)
		fmt.Printf("过滤敏感词加载文件错误：%v", err)
	}
}
