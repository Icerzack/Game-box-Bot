package app

import (
	botConfig "VK-bot/internal/config"
	vkAPI "VK-bot/internal/pkg/api"
	updateObjects "VK-bot/internal/pkg/api/objects"
	coinOperations "VK-bot/internal/pkg/operations/coin"
	welcomeOperations "VK-bot/internal/pkg/operations/welcome"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Bot struct {
	cfg            botConfig.Config
	debugMode      bool
	openedChannels map[int]chan string
}

func NewBot(cfg botConfig.Config) *Bot {
	return &Bot{
		cfg:            cfg,
		debugMode:      false,
		openedChannels: make(map[int]chan string),
	}
}

func (b *Bot) Start() {
	server, key, ts, err := b.getLongPollServer()
	if err != nil {
		b.log(fmt.Sprintf("%s", err))
		return
	}
	for {
		resp, err := b.getUpdates(server, key, ts, b.cfg.Wait)
		if err != nil {
			b.log(fmt.Sprintf("%s", err))
			return
		}
		for _, upd := range resp.Updates {
			go b.updateHandler(upd)
		}
		ts = resp.Ts
	}
}

func (b *Bot) updateHandler(upd vkAPI.Update) {
	switch upd.Type {
	case updateObjects.NewMessage:
		var messageObject updateObjects.MessageNewObject
		err := json.Unmarshal(upd.Object, &messageObject)
		if err != nil {
			b.log(fmt.Sprintf("Failed to unmarshal message_new object: %s", err))
			return
		}
		receivedText := messageObject.Message.Text
		senderID := messageObject.Message.FromID
		b.log(fmt.Sprintf("Received text: %s\nFrom id: %d\n", receivedText, senderID))

		if _, ok := b.openedChannels[senderID]; ok {
			b.openedChannels[senderID] <- receivedText
			return
		}

		if receivedText == "Начать" {
			err := welcomeOperations.SendWelcomeMessage(&b.cfg, senderID)
			if err != nil {
				b.log(fmt.Sprintf("Failed to send welcome message: %s", err))
				return
			}
		}
		if receivedText == coinOperations.TossACoin {
			if _, ok := b.openedChannels[senderID]; ok {
				return
			}
			err := coinOperations.SendCoinMessage(&b.cfg, senderID)
			if err != nil {
				b.log(fmt.Sprintf("Failed to send coin message: %s", err))
				return
			}
			messageChan := make(chan string)
			b.openedChannels[senderID] = messageChan
			go b.coinHandler(senderID)
		}

	}
}

func (b *Bot) coinHandler(forID int) {
	ch := b.openedChannels[forID]
	defer func() {
		close(ch)
		delete(b.openedChannels, forID)
	}()

	inp := ""

Loop:
	for {
		select {
		case inp = <-ch:
			if inp != coinOperations.Heads && inp != coinOperations.Tails {
				err := coinOperations.SendWrongMessage(&b.cfg, forID)
				if err != nil {
					b.log(fmt.Sprintf("Failed to send wrong-coin message: %s", err))
					return
				}
			} else {
				break Loop
			}
		}
	}

	coin := [...]string{coinOperations.Heads, coinOperations.Tails}

	rand.Seed(time.Now().UnixNano())
	randomSide := coin[rand.Intn(len(coin))]

	if inp == randomSide {
		err := coinOperations.SendResultMessage(&b.cfg, forID, true, randomSide)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send result-coin message: %s", err))
			return
		}
	} else {
		err := coinOperations.SendResultMessage(&b.cfg, forID, false, randomSide)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send result-coin message: %s", err))
			return
		}
	}

}

func (b *Bot) SetDebugMode(mode bool) {
	b.debugMode = mode
}

func (b *Bot) log(message string) {
	if !b.debugMode {
		return
	}
	fmt.Println(message)
}

func (b *Bot) getLongPollServer() (string, string, string, error) {
	urlToSend := fmt.Sprintf("%sgroups.getLongPollServer?group_id=%s&access_token=%s&v=%s", b.cfg.ApiURL, b.cfg.GroupID, b.cfg.Token, b.cfg.ApiVer)
	resp, err := http.Get(urlToSend)
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	var longPollServer vkAPI.LongPollServer

	err = json.NewDecoder(resp.Body).Decode(&longPollServer)
	if err != nil {
		return "", "", "", err
	}
	return longPollServer.Response.Server, longPollServer.Response.Key, longPollServer.Response.Ts, nil
}

func (b *Bot) getUpdates(server, key, ts, wait string) (*vkAPI.LongPollResponse, error) {
	urlToSend := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%s", server, key, ts, wait)
	resp, err := http.Get(urlToSend)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lpResp vkAPI.LongPollResponse
	err = json.NewDecoder(resp.Body).Decode(&lpResp)
	if err != nil {
		return nil, err
	}
	return &lpResp, nil
}
