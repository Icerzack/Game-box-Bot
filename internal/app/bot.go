package app

import (
	botConfig "VK-bot/internal/config"
	vkAPI "VK-bot/internal/pkg/api"
	updateObjects "VK-bot/internal/pkg/api/objects"
	roomOperations "VK-bot/internal/pkg/modes/room"
	"encoding/json"
	"fmt"
	"net/http"
)

type Bot struct {
	cfg            botConfig.Config
	debugMode      bool
	openedChannels map[int]chan string
	roomsHub       roomOperations.RoomsHub
}

func NewBot(cfg botConfig.Config) *Bot {
	return &Bot{
		cfg:            cfg,
		debugMode:      false,
		openedChannels: make(map[int]chan string),
		roomsHub:       roomOperations.RoomsHub{Rooms: make(map[int32]*roomOperations.Room)},
	}
}

func (b *Bot) SetDebugMode(mode bool) {
	b.debugMode = mode
}

func (b *Bot) Start() {
	longPollServerResponse, err := b.getLongPollServer()
	if err != nil {
		b.log(fmt.Sprintf("%s", err))
		return
	}
	b.log("Bot successfully launched up!\nLongPoll server parameters:\nSERVER: " + longPollServerResponse.Server + "\nKEY: " + longPollServerResponse.Key)
	ts := longPollServerResponse.Ts
	for {
		resp, err := b.getUpdates(longPollServerResponse.Server, longPollServerResponse.Key, ts, b.cfg.Wait)
		if err != nil {
			b.log(fmt.Sprintf("%s", err))
			return
		}
		for _, upd := range resp.Updates {
			go b.updateHandler(&upd)
		}
		ts = resp.Ts
	}
}

func (b *Bot) updateHandler(upd *vkAPI.Update) {
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
		b.messageHandler(senderID, receivedText)
	}
}

func (b *Bot) log(message string) {
	if !b.debugMode {
		return
	}
	fmt.Println(message)
}

func (b *Bot) getLongPollServer() (*vkAPI.LongPollServerResponse, error) {
	urlToSend := fmt.Sprintf("%sgroups.getLongPollServer?group_id=%s&access_token=%s&v=%s", b.cfg.ApiURL, b.cfg.GroupID, b.cfg.Token, b.cfg.ApiVer)
	resp, err := http.Get(urlToSend)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var longPollServer vkAPI.LongPollServer

	err = json.NewDecoder(resp.Body).Decode(&longPollServer)
	if err != nil {
		return nil, err
	}
	return &longPollServer.Response, nil
}

func (b *Bot) getUpdates(server, key, ts, wait string) (*vkAPI.LongPollUpdateResponse, error) {
	urlToSend := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%s", server, key, ts, wait)
	resp, err := http.Get(urlToSend)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lpResp vkAPI.LongPollUpdateResponse
	err = json.NewDecoder(resp.Body).Decode(&lpResp)
	if err != nil {
		return nil, err
	}
	return &lpResp, nil
}
