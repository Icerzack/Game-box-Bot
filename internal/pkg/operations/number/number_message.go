package number

import (
	"VK-bot/internal/config"
	"VK-bot/internal/pkg/keyboard"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	GetANumber  = "🔢Получить число"
	SetInterval = "Установить интервал"
	Confirm     = "✅Подтвердить"
)

func SendNumberMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{SetInterval, Confirm}
	keyboardJSON := keyboard.GenerateKeyboard3x1(buttonsText)

	text := `С помощью кнопки ниже установи интервал. По умолчанию 0-10. 

Как закончишь, жми "✅Подтвердить".`

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendPromptMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	text := "Введи два числа, через пробел:"

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer)
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendIntervalMessage(cfg *config.Config, peerID, low, top int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{SetInterval, Confirm}
	keyboardJSON := keyboard.GenerateKeyboard3x1(buttonsText)

	text := "Текущий интервал: " + strconv.Itoa(low) + "..." + strconv.Itoa(top)

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendNumberOutOfBoundsMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{SetInterval, Confirm}
	keyboardJSON := keyboard.GenerateKeyboard3x1(buttonsText)

	text := "Твое число слишком большое, либо ты его не ввел, попробуй заново."

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendIsNotANumberMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{SetInterval, Confirm}
	keyboardJSON := keyboard.GenerateKeyboard3x1(buttonsText)

	text := "Кажется, ты ввел не число :/"

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

	buttonsText := []string{SetInterval, Confirm}
	keyboardJSON := keyboard.GenerateKeyboard3x1(buttonsText)

	text := "Доступны только опции из ниже предложенных."

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendResultMessage(cfg *config.Config, peerID int, number int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	text := "Твое число: " + strconv.Itoa(number) + "!"

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer)
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}
