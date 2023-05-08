package coin

import (
	"VK-bot/internal/config"
	"VK-bot/internal/pkg/keyboard"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const (
	TossACoin = "🪙Подбросить монетку"
	Heads     = "Орел"
	Tails     = "Решка"
)

func SendCoinMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{Heads, Tails}
	keyboardJSON := keyboard.GenerateKeyboard2x1(buttonsText)

	text := `
Что выберешь - орел или решка?
	`

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendWrongMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{Heads, Tails}
	keyboardJSON := keyboard.GenerateKeyboard2x1(buttonsText)

	text := `
Нажми на одну из кнопок ниже: орел или решка?
	`

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendResultMessage(cfg *config.Config, peerID int, isWin bool, chosenSide string) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{Heads, Tails}
	keyboardJSON := keyboard.GenerateKeyboard2x2(buttonsText)

	text := ""
	if isWin {
		text = "Повезло! " + chosenSide + "!"
	} else {
		text = "Эх, не повезло: " + chosenSide + "."
	}

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}
