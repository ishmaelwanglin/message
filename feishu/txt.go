package feishu

import (
	"encoding/json"
	"fmt"
	logg "powerguard/log"
	"powerguard/std"
)

type MessageText struct {
	MsgType string  `json:"msg_type"`
	Content content `json:"content"`
}

type content struct {
	Text string `json:"text"`
}

type Response struct {
	StatusCode    int         `json:"StatusCode"`
	StatusMessage string      `json:"StatusMessage"`
	Code          int         `json:"code"`
	Data          interface{} `json:"data"`
	Msg           string      `json:"msg"`
}

func NewMessageTxt(payload interface{}) *MessageText {
	var text string
	switch content := payload.(type) {
	case string:
		text = fmt.Sprintf("IPMI告警信息: %s: %s", std.IP, content)
	default:
		logg.Error("unknown payload type")
	}

	return &MessageText{
		MsgType: "text",
		Content: content{
			Text: text,
		},
	}
}

func (mesg *MessageText) payload() ([]byte, error) {
	return json.Marshal(mesg)
}

func (mesg *MessageText) Send() error {
	return sendFeishuMessage(mesg)
}
