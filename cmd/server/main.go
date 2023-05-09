package main

import (
	"VK-bot/internal/app"
	"VK-bot/internal/config"
	"VK-bot/tools"
	"os"
)

const (
	apiURL = "https://api.vk.com/method/"
	apiVer = "5.131"
	wait   = "25"
)

func main() {
	tools.LoadEnv()

	bot := app.NewBot(config.Config{
		Token:   os.Getenv("TOKEN"),
		ApiURL:  apiURL,
		ApiVer:  apiVer,
		GroupID: os.Getenv("GROUP_ID"),
		Wait:    wait,
	})

	bot.SetDebugMode(true)
	bot.Start()
}
