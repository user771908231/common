package roomAgent

import (
	"casino_common/proto/ddproto"
	"errors"
	"strings"
)

//检查空闲

//创建房间
func CreateDeskByOption(owner uint32, group_id int32, game_option string) (*ddproto.CommonDeskByAgent, error) {
	switch {
	case strings.Contains(game_option, "经典跑得快"), strings.Contains(game_option, "十五张"), strings.Contains(game_option, "十六张"):
		err, desk := DoPdkKaifang(owner, group_id, game_option)
		if desk == nil {
			return nil, err
		} else {
			return desk, nil
		}
	case strings.Contains(game_option, "红中"), strings.Contains(game_option, "转转"):
		err, desk := DoZzHzKaifang(owner, group_id, game_option)
		if desk == nil {
			return nil, err
		} else {
			return desk, nil
		}
	case strings.Contains(game_option, "松江河麻将"), strings.Contains(game_option, "松江河"):
		err, desk := DoSjhKaifang(owner, group_id, game_option)
		if desk == nil {
			return nil, err
		} else {
			return desk, nil
		}
	case strings.Contains(game_option, "经典牛牛"), strings.Contains(game_option, "急速牛牛"):
		err, desk := DoNiuniuKaifang(owner, group_id, game_option)
		if desk == nil {
			return nil, err
		} else {
			return desk, nil
		}
	case strings.Contains(game_option, "血战到底"),
		strings.Contains(game_option, "长沙麻将"),
		strings.Contains(game_option, "两人两房"),
		strings.Contains(game_option, "三人两房"),
		strings.Contains(game_option, "四人两房"),
		strings.Contains(game_option, "两人三房"),
		strings.Contains(game_option, "三人三房"),
		strings.Contains(game_option, "德阳麻将"),
		strings.Contains(game_option, "倒倒胡"),
		strings.Contains(game_option, "血流成河"):
		err, desk := DoSiChuanKaifang(owner, group_id, game_option)
		if desk == nil {
			return nil, err
		} else {
			return desk, nil
		}
	}

	return nil, errors.New("开房失败！")
}
