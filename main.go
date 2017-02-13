package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Payload struct {
	Channel     string       `json:"channel"`
	Username    string       `json:"username"`
	IconEmoji   string       `json:"icon_emoji"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Fallback string `json:"fallback"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Color    string `json:"color"`
	Footer   string `json:"footer"`
	Ts       int64  `json:"ts"`
}

func main() {
	usage := "Usage: SLACK_WEBHOOK_URL=\"YOUR_WEBHOOK_URL\" go-zabbix-slack-alert <CHANNEL> <SUBJECT> <MESSAGE>"
	emoji := ":ghost:"
	color := "warning"
	timeout := 5 * time.Second

	if len(os.Args) != 4 {
		log.Fatal(usage)
	}
	to := os.Args[1]
	subj := os.Args[2]
	msg := os.Args[3]
	url := os.Getenv("SLACK_WEBHOOK_URL")

	if url == "" {
		log.Fatal(usage)
	} else if to == "" {
		log.Fatal(usage)
	} else if subj == "" {
		log.Fatal(usage)
	} else if msg == "" {
		log.Fatal(usage)
	}

	if strings.Contains(subj, "OK:") {
		emoji = ":smile:"
		color = "good"
	} else if strings.Contains(subj, "PROBLEM:") {
		emoji = ":frowning:"
		color = "danger"
	}

	payload := Payload{
		Channel:   to,
		Username:  "Zabbix",
		IconEmoji: emoji,
		Attachments: []Attachment{
			{
				Fallback: fmt.Sprintf("%s: %s", subj, msg),
				Color:    color,
				Title:    subj,
				Text:     msg,
				Footer:   "go-zabbix-slack-alert",
				Ts:       time.Now().Unix(),
			},
		},
	}
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
