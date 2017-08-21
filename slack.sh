#!/bin/bash

export SLACK_WEBHOOK_URL="YOUR_SLACK_WEBHOOK_URL"
/path/to/go-zabbix-slack-alert "$1" "$2" "$3"
