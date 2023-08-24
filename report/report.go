package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	sceret     = "数据巡检："
    webhookURL = "保留在本地"
)

type DingTalkResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
type DingTalkMessage struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func sendDingTalkMessage(webhookURL string, message string) error {
	//timestamp := time.Now().UnixNano() / 1e6
	//sign := generateSign(secret, timestamp)

	dingTalkMessage := DingTalkMessage{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: message,
		},
	}

	payload, err := json.Marshal(dingTalkMessage)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send DingTalk message, status code: %d", resp.StatusCode)
	}
	var response DingTalkResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}
	if response.ErrCode != 0 {
		return fmt.Errorf(response.ErrMsg)
	}
	fmt.Println(response.ErrCode, response.ErrMsg)
	return nil
}

func Run(message string) {
	message = sceret + message
	err := sendDingTalkMessage(webhookURL, message)
	if err != nil {
		log.Fatal(err)
	}
}
