package roomAgent

import (
	"strings"
)

type CreateConfig [][][]string

//解析关键词
func (c *CreateConfig) GetKeywords(owner_opt, user_opt string) (res []string) {
	for _, paras := range *c {
		default_paras := ""
		has_paras := false
		//匹配用户选项
		for _, alias := range paras {
			for _, key := range alias {
				if strings.Contains(user_opt, key) {
					has_paras = true
					default_paras = alias[0]
					//log.T("set default_paras 1 %v user_opt[%v] key[%v]", default_paras, user_opt, key)
					break
				}
			}
			if has_paras {
				break
			}
		}
		//匹配房主选项
		if has_paras == false {
			for paras_i, alias := range paras {
				for alias_i, key := range alias {
					//默认选项
					if paras_i == 0 && alias_i == 0 {
						default_paras = key
						//log.T("set default_paras 2 %v", default_paras)
					}
					if strings.Contains(owner_opt, key) {
						has_paras = true
						default_paras = alias[0]
						//log.T("set default_paras 3 %v", default_paras)
						break
					}
				}
				if has_paras {
					break
				}
			}
		}

		//log.T("%v", res)
		//插入
		res = append(res, default_paras)
	}
	return
}
