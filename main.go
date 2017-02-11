package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"
)

type Payload struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
}

func main() {
	username := "Zabbix"
	to := os.Args[1]
	subj := os.Args[2]
	msg := os.Args[3]
	text := fmt.Sprintf("*%s:*\n%s", subj, msg)

	r := regexp.MustCompile(`^RECOVER(Y|ED)?$`)
	emoji := ":ghost:"
	if r.MatchString(subj) {
		emoji = ":smile:"
	} else if subj == "PROBLEM" {
		emoji = ":frowning:"
	}

	jsonBytes, err := json.Marshal(Payload{
		Channel:   to,
		Username:  username,
		Text:      text,
		IconEmoji: emoji,
	})
	if err != nil {
		panic(err)
	}

	url := os.Getenv("SLACK_WEBHOOK_URL")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
