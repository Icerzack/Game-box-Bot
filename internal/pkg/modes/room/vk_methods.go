package room

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
	EnterRoom  = "👥Создать/Найти комнату"
	CreateRoom = "🔨Создать комнату"
	JoinRoom   = "🔍Найти комнату"
	Exit       = "🚪Выйти"
)

func SendRoomMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{CreateRoom, JoinRoom, Exit}
	keyboardJSON := keyboard.GenerateKeyboard1x3(buttonsText)

	text := "Если ты знаешь ID комнаты, то выбирай \"" + JoinRoom + "\".\n\nЧтобы создать -> выбери \"" + CreateRoom + "\".\n\nДля изменения своего состояния, просто отправь новое сообщение в чат комнаты!"

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

	buttonsText := []string{CreateRoom, JoinRoom, Exit}
	keyboardJSON := keyboard.GenerateKeyboard1x3(buttonsText)

	text := "Выбери одно из ниже предложенных."

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendEnterUsernameMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	text := "Введи свое имя, которое будет отображаться в комнате:"

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer)
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendEnterRoomCodeMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	text := "Введи код комнаты:"

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer)
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendNoRoomFoundMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	text := "Комнаты с таким кодом не найдено :("

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer)
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendGeneratedRoomCodeMessage(cfg *config.Config, peerID int, code string) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	text := "Код созданной комнаты:\n\n " + code + "\n\n Запиши его и передай другим, чтоб они могли к тебе подключиться."

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer)
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}

func SendRoomStatusMessage(cfg *config.Config, room *Room) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{Exit}
	keyboardJSON := keyboard.GenerateKeyboard1x1(buttonsText)

	var status string

	for _, user := range room.Users {
		status += user.UserName + ": " + user.Value + "\n"
	}

	for _, user := range room.Users {
		text := "Текущий статус комнаты:\n\n" + status

		urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, user.UserID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
		_, err := http.Get(urlToSend)
		if err != nil {
			return err
		}
	}

	return nil
}
