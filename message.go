package message

import (
	"powerguard/message/feishu"
)

const (
	FeiShu = "feishu"
	Ws     = "websocket"
)

type Messager interface {
	Send() error
}

func NewMessage(method string, payload interface{}) Messager {
	switch method {
	case FeiShu:
		return feishu.NewMessageTxt(payload)
	}

	return nil
}
