package push

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"pushbot/config"
)

// WeChatRobotMsg 企业微信机器人消息结构
type WeChatRobotMsg struct {
	MsgType  string       `json:"msgtype"`
	Text     *TextMsg     `json:"text,omitempty"`
	Markdown *MarkdownMsg `json:"markdown,omitempty"`
}

// TextMsg 文本消息内容
type TextMsg struct {
	Content       string   `json:"content"`
	MentionedList []string `json:"mentioned_list,omitempty"`
}

// MarkdownMsg Markdown消息内容
type MarkdownMsg struct {
	Content string `json:"content"`
}

// PushWX 发送消息到企业微信机器人
func PushWX(message WeChatRobotMsg) error {
	webhookURL := config.HandleYaml().Push.Weixin.Hook
	// 序列化消息为 JSON
	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// 发起 POST 请求
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
