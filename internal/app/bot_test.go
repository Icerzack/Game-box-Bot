package app

import (
	botConfig "VK-bot/internal/config"
	"VK-bot/internal/pkg/api"
	"VK-bot/tools"
	"fmt"
	"os"
	"testing"
)

func TestNewBot(t *testing.T) {
	cfg := botConfig.Config{}

	b := NewBot(cfg)
	t.Run("config", func(t *testing.T) {
		if b.cfg != cfg {
			t.Errorf("Expected cfg to be %v, but got %v", cfg, b.cfg)
		}
	})

	t.Run("debugMode", func(t *testing.T) {
		if b.debugMode != false {
			t.Errorf("Expected debugMode to be false, but got %v", b.debugMode)
		}
	})

	t.Run("openedChannels", func(t *testing.T) {
		if b.openedChannels == nil {
			t.Errorf("Expected openedChannels to be initialized, but got nil")
		}
	})

}

func TestBot_SetDebugMode(t *testing.T) {
	b := &Bot{}

	t.Run("isTrue", func(t *testing.T) {
		b.SetDebugMode(true)
		if b.debugMode != true {
			t.Errorf("Expected debugMode to be true, but got %v", b.debugMode)
		}
	})

	t.Run("isFalse", func(t *testing.T) {
		b.SetDebugMode(false)
		if b.debugMode != false {
			t.Errorf("Expected debugMode to be false, but got %v", b.debugMode)
		}
	})
}

func TestBot_Start(t *testing.T) {
	tools.LoadEnv("../../.env")

	const (
		apiURL = "https://api.vk.com/method/"
		apiVer = "5.131"
		wait   = "25"
	)

	cfg := botConfig.Config{
		Token:   os.Getenv("TOKEN"),
		ApiURL:  apiURL,
		ApiVer:  apiVer,
		GroupID: os.Getenv("GROUP_ID"),
		Wait:    wait,
	}

	b := NewBot(cfg)

	var longPollServerResponse *api.LongPollServerResponse
	t.Run("getLongPollServer", func(t *testing.T) {
		longPollServerResponse, _ = b.getLongPollServer()

		if longPollServerResponse.Server != fmt.Sprintf("https://lp.vk.com/wh%s", cfg.GroupID) {
			t.Errorf("Expected Server to be https://lp.vk.com/wh%s, but got https://lp.vk.com/wh%s", cfg.GroupID, longPollServerResponse.Server)
		}
		if longPollServerResponse.Key == "" {
			t.Errorf("Null Key returned")
		}
		if longPollServerResponse.Ts == "" {
			t.Errorf("Null Ts returned")
		}
	})

	t.Run("getUpdates", func(t *testing.T) {
		longPollUpdateResponse, _ := b.getUpdates(longPollServerResponse.Server, longPollServerResponse.Key, longPollServerResponse.Ts, "1")

		if longPollUpdateResponse.Ts == "" {
			t.Errorf("Null Ts returned")
		}
	})
}
