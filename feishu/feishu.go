package feishu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"powerguard/conf"
	"time"
)

type feishuMessager interface {
	payload() ([]byte, error)
	Send() error
}
type Response struct {
	StatusCode    int         `json:"StatusCode"`
	StatusMessage string      `json:"StatusMessage"`
	Code          int         `json:"code"`
	Data          interface{} `json:"data"`
	Msg           string      `json:"msg"`
}

// 通用的发送函数
func sendFeishuMessage(msg feishuMessager) error {
	var (
		payload []byte
		err     error
	)
	if payload, err = msg.payload(); err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, conf.Config.Message.WebHook, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	p, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	respData := new(Response)
	if err = json.Unmarshal(p, respData); err != nil {
		return err
	}

	if respData.Code != 0 {
		return fmt.Errorf("feishu send message failed")
	}

	return nil
}
