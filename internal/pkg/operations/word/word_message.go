package word

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
	GetAWord  = "📖Получить слово"
	Noun      = "Существительное"
	Adjective = "Прилагательное"
	Animal    = "Животное"
)

func SendWordMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{Noun, Adjective, Animal}
	keyboardJSON := keyboard.GenerateKeyboard3x1(buttonsText)

	text := "Выбери, какое слово хочешь получить?"

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

	buttonsText := []string{Noun, Adjective, Animal}
	keyboardJSON := keyboard.GenerateKeyboard3x1(buttonsText)

	text := "Выбери одно из ниже предложенных."

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendResultMessage(cfg *config.Config, peerID int, word string) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	text := "Твое слово: " + word + "!"

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer)
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}
