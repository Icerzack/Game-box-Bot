package common

import (
	"VK-bot/internal/config"
	"VK-bot/internal/pkg/keyboard"
	"VK-bot/internal/pkg/operations/coin"
	"VK-bot/internal/pkg/operations/dice"
	"VK-bot/internal/pkg/operations/number"
	"VK-bot/internal/pkg/operations/word"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func SendDefaultMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{coin.TossACoin, dice.TossADice, word.GetAWord, number.GetANumber}
	keyboardJSON := keyboard.GenerateKeyboard2x2(buttonsText)

	text := "Что делаем дальше?"

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendNoOpMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{coin.TossACoin, dice.TossADice, word.GetAWord, number.GetANumber}
	keyboardJSON := keyboard.GenerateKeyboard2x2(buttonsText)

	text := "Такой команды я не знаю - выбирай что-нибудь из предложенного!"

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}
