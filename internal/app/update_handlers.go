package app

import (
	coinOperations "VK-bot/internal/pkg/modes/coin"
	commonOperations "VK-bot/internal/pkg/modes/common"
	diceOperations "VK-bot/internal/pkg/modes/dice"
	numberOperations "VK-bot/internal/pkg/modes/number"
	roomOperations "VK-bot/internal/pkg/modes/room"
	welcomeOperations "VK-bot/internal/pkg/modes/welcome"
	wordOperations "VK-bot/internal/pkg/modes/word"
	"fmt"
)

func (b *Bot) messageHandler(senderID int, text string) {
	if _, ok := b.openedChannels[senderID]; ok {
		b.openedChannels[senderID] <- text
		return
	}
	switch text {
	case "Начать", "Start":
		err := welcomeOperations.SendWelcomeMessage(&b.cfg, senderID)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send welcome message: %s", err))
			return
		}
	case coinOperations.TossACoin:
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
	case diceOperations.TossADice:
		if _, ok := b.openedChannels[senderID]; ok {
			return
		}
		err := diceOperations.SendDiceMessage(&b.cfg, senderID)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send dice message: %s", err))
			return
		}
		messageChan := make(chan string)
		b.openedChannels[senderID] = messageChan
		go b.diceHandler(senderID)
	case wordOperations.GetAWord:
		if _, ok := b.openedChannels[senderID]; ok {
			return
		}
		err := wordOperations.SendWordMessage(&b.cfg, senderID)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send word message: %s", err))
			return
		}
		messageChan := make(chan string)
		b.openedChannels[senderID] = messageChan
		go b.wordHandler(senderID)
	case numberOperations.GetANumber:
		if _, ok := b.openedChannels[senderID]; ok {
			return
		}
		err := numberOperations.SendNumberMessage(&b.cfg, senderID)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send number message: %s", err))
			return
		}
		messageChan := make(chan string)
		b.openedChannels[senderID] = messageChan
		go b.numberHandler(senderID)
	case roomOperations.EnterRoom:
		if _, ok := b.openedChannels[senderID]; ok {
			return
		}
		err := roomOperations.SendRoomMessage(&b.cfg, senderID)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send room message: %s", err))
			return
		}
		messageChan := make(chan string)
		b.openedChannels[senderID] = messageChan
		go b.roomHandler(senderID)
	default:
		err := commonOperations.SendNoOpMessage(&b.cfg, senderID)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send no-op message: %s", err))
			return
		}
	}
}
