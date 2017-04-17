package inco

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// Message is object for incoming webhook.
type Message struct {
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
	Username  string `json:"username"`
	Channel   string `json:"channel"`
}

// Incoming post to incoming webhook.
func Incoming(url string, msg *Message) (err error) {
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	return
}
