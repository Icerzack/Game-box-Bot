package main

import (
	"VK-bot/internal/app"
	"VK-bot/internal/config"
)

const (
	token   = "YOUR_TOKEN"
	apiURL  = "https://api.vk.com/method/"
	apiVer  = "5.131"
	groupID = "YOUR_GROUP_ID"
	wait    = "25"
)

func main() {
	bot := app.NewBot(config.Config{
		Token:   token,
		ApiURL:  apiURL,
		ApiVer:  apiVer,
		GroupID: groupID,
		Wait:    wait,
	})

	bot.SetDebugMode(true)
	bot.Start()
}
