package app

import (
	coinOperations "VK-bot/internal/pkg/operations/coin"
	commonOperations "VK-bot/internal/pkg/operations/common"
	diceOperations "VK-bot/internal/pkg/operations/dice"
	wordOperations "VK-bot/internal/pkg/operations/word"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

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
	err := commonOperations.SendDefaultMessage(&b.cfg, forID)
	if err != nil {
		b.log(fmt.Sprintf("Failed to send default message: %s", err))
		return
	}
}

func (b *Bot) diceHandler(forID int) {
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
			if inp != diceOperations.One && inp != diceOperations.Two && inp != diceOperations.Three {
				err := diceOperations.SendWrongMessage(&b.cfg, forID)
				if err != nil {
					b.log(fmt.Sprintf("Failed to send wrong-dice message: %s", err))
					return
				}
			} else {
				break Loop
			}
		}
	}

	dice := [...]int{1, 2, 3, 4, 5, 6}

	if inp == diceOperations.One {
		rand.Seed(time.Now().UnixNano())
		random := dice[rand.Intn(len(dice))]
		err := diceOperations.SendResultMessage(&b.cfg, forID, random)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send result-dice message: %s", err))
			return
		}
	}
	if inp == diceOperations.Two {
		rand.Seed(time.Now().UnixNano())
		random1 := dice[rand.Intn(len(dice))]
		random2 := dice[rand.Intn(len(dice))]
		err := diceOperations.SendResultMessage(&b.cfg, forID, random1, random2)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send result-dice message: %s", err))
			return
		}
	}
	if inp == diceOperations.Three {
		rand.Seed(time.Now().UnixNano())
		random1 := dice[rand.Intn(len(dice))]
		random2 := dice[rand.Intn(len(dice))]
		random3 := dice[rand.Intn(len(dice))]
		err := diceOperations.SendResultMessage(&b.cfg, forID, random1, random2, random3)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send result-dice message: %s", err))
			return
		}
	}
	err := commonOperations.SendDefaultMessage(&b.cfg, forID)
	if err != nil {
		b.log(fmt.Sprintf("Failed to send default message: %s", err))
		return
	}
}

func (b *Bot) wordHandler(forID int) {
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
			if inp != wordOperations.Noun && inp != wordOperations.Adjective && inp != wordOperations.Animal {
				err := wordOperations.SendWrongMessage(&b.cfg, forID)
				if err != nil {
					b.log(fmt.Sprintf("Failed to send wrong-word message: %s", err))
					return
				}
			} else {
				break Loop
			}
		}
	}
	urlToSend := ""
	switch inp {
	case wordOperations.Noun:
		urlToSend = "https://random-word-form.herokuapp.com/random/noun"
	case wordOperations.Adjective:
		urlToSend = "https://random-word-form.herokuapp.com/random/adjective"
	case wordOperations.Animal:
		urlToSend = "https://random-word-form.herokuapp.com/random/animal"
	}

	resp, err := http.Get(urlToSend)
	if err != nil {
		b.log(fmt.Sprintf("Failed to send word request: %s", err))
		return
	}
	defer resp.Body.Close()

	var word wordOperations.Word
	err = json.NewDecoder(resp.Body).Decode(&word)
	if err != nil {
		b.log(fmt.Sprintf("Failed to decode word: %s", err))
		return
	}

	err = wordOperations.SendResultMessage(&b.cfg, forID, word[0])
	if err != nil {
		b.log(fmt.Sprintf("Failed to send word-result message: %s", err))
		return
	}

	err = commonOperations.SendDefaultMessage(&b.cfg, forID)
	if err != nil {
		b.log(fmt.Sprintf("Failed to send default message: %s", err))
		return
	}
}

func (b *Bot) numberHandler(forID int) {}
