package openwechat

import (
	"errors"
	"strconv"
)

type Selector string

const (
	SelectorNormal          Selector = "0" // 正常
	SelectorNewMsg          Selector = "2" // 有新消息
	SelectorModContact      Selector = "4" // 联系人信息变更
	SelectorAddOrDelContact Selector = "6" // 添加或删除联系人
	SelectorModChatRoom     Selector = "7" // 进入或退出聊天室
)

type SyncCheckResponse struct {
	RetCode  string
	Selector Selector
}

func (s SyncCheckResponse) Success() bool {
	return s.RetCode == "0"
}

func (s SyncCheckResponse) NorMal() bool {
	return s.Success() && s.Selector == "0"
}

func (s SyncCheckResponse) Err() error {
	if s.Success() {
		return nil
	}
	i, err := strconv.Atoi(s.RetCode)
	if err != nil {
		return errors.New("sync check unknown error")
	}
	return Ret(i)
}
