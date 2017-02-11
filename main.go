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
	Channel     string       `json:"channel"`
	Username    string       `json:"username"`
	Text        string       `json:"text"`
	IconEmoji   string       `json:"icon_emoji"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Fallback string `json:"fallback"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Color    string `json:"color"`
}

func main() {
	to := os.Args[1]
	subj := os.Args[2]
	msg := os.Args[3]

	r := regexp.MustCompile(`^RECOVER(Y|ED)?$`)
	emoji := ":ghost:"
	color := "warning"
	if r.MatchString(subj) {
		emoji = ":smile:"
		color = "good"
	} else if subj == "PROBLEM" {
		emoji = ":frowning:"
		color = "danger"
	}

	jsonBytes, err := json.Marshal(Payload{
		Channel:   to,
		Username:  "Zabbix",
		IconEmoji: emoji,
		Attachments: []Attachment{
			{
				Fallback: fmt.Sprintf("%s: %s", subj, msg),
				Color:    color,
				Title:    subj,
				Text:     msg,
			},
		},
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
