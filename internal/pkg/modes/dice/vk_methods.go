package dice

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
	TossADice = "🎲Подбросить кубик"
	One       = "Один"
	Two       = "Два"
	Three     = "Три"
)

func SendDiceMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{One, Two, Three}
	keyboardJSON := keyboard.GenerateKeyboard3x1(buttonsText)

	text := "Сколько кубиков бросаем?"

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

	buttonsText := []string{One, Two, Three}
	keyboardJSON := keyboard.GenerateKeyboard2x1(buttonsText)

	text := "Выбирай значения из ниже предложенных."

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendResultMessage(cfg *config.Config, peerID int, values ...int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	valuesString := ""
	for _, val := range values {
		valuesString += " " + fmt.Sprintf("%d", val)
	}

	text := "Выпало следующее:" + valuesString + "."

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer)
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}
