package app

import (
	coinOperations "VK-bot/internal/pkg/operations/coin"
	commonOperations "VK-bot/internal/pkg/operations/common"
	diceOperations "VK-bot/internal/pkg/operations/dice"
	numberOperations "VK-bot/internal/pkg/operations/number"
	wordOperations "VK-bot/internal/pkg/operations/word"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (b *Bot) coinHandler(forID int) {
	ch := b.openedChannels[forID]
	defer func() {
		close(ch)
		delete(b.openedChannels, forID)
	}()

	input := ""

Loop:
	for {
		select {
		case input = <-ch:
			if input != coinOperations.Heads && input != coinOperations.Tails {
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

	if input == randomSide {
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

	input := ""

Loop:
	for {
		select {
		case input = <-ch:
			if input != diceOperations.One && input != diceOperations.Two && input != diceOperations.Three {
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

	if input == diceOperations.One {
		rand.Seed(time.Now().UnixNano())
		random := dice[rand.Intn(len(dice))]
		err := diceOperations.SendResultMessage(&b.cfg, forID, random)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send result-dice message: %s", err))
			return
		}
	}
	if input == diceOperations.Two {
		rand.Seed(time.Now().UnixNano())
		random1 := dice[rand.Intn(len(dice))]
		random2 := dice[rand.Intn(len(dice))]
		err := diceOperations.SendResultMessage(&b.cfg, forID, random1, random2)
		if err != nil {
			b.log(fmt.Sprintf("Failed to send result-dice message: %s", err))
			return
		}
	}
	if input == diceOperations.Three {
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

	input := ""

Loop:
	for {
		select {
		case input = <-ch:
			if input != wordOperations.Noun && input != wordOperations.Adjective && input != wordOperations.Animal {
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
	switch input {
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

func (b *Bot) numberHandler(forID int) {
	ch := b.openedChannels[forID]
	defer func() {
		close(ch)
		delete(b.openedChannels, forID)
	}()

	input := ""
	lowBound, topBound := 0, 10

Loop:
	for {
		select {
		case input = <-ch:
			if input == numberOperations.SetInterval {
				err := numberOperations.SendPromptMessage(&b.cfg, forID)
				if err != nil {
					b.log(fmt.Sprintf("Failed to send number-prompt message: %s", err))
					return
				}
				select {
				case input = <-ch:
					inputArray := strings.Split(input, " ")
					if len(inputArray) < 2 || len(inputArray[0]) > 10 || len(inputArray[1]) > 10 {
						err := numberOperations.SendNumberOutOfBoundsMessage(&b.cfg, forID)
						if err != nil {
							b.log(fmt.Sprintf("Failed to send number-prompt-OOB message: %s", err))
							return
						}
						break
					} else {
						num1, err1 := strconv.Atoi(inputArray[0])
						num2, err2 := strconv.Atoi(inputArray[1])
						if err1 == nil && err2 == nil {
							if num1 > num2 {
								lowBound, topBound = num2, num1
							} else {
								lowBound, topBound = num1, num2
							}
							break
						} else {
							err := numberOperations.SendIsNotANumberMessage(&b.cfg, forID)
							if err != nil {
								b.log(fmt.Sprintf("Failed to send number-prompt-INAN message: %s", err))
								return
							}
						}
					}
				}
				err = numberOperations.SendIntervalMessage(&b.cfg, forID, lowBound, topBound)
				if err != nil {
					b.log(fmt.Sprintf("Failed to send number-interval message: %s", err))
					return
				}
			} else if input == numberOperations.Confirm {
				break Loop
			} else {
				err := numberOperations.SendWrongMessage(&b.cfg, forID)
				if err != nil {
					b.log(fmt.Sprintf("Failed to send number-wrong message: %s", err))
					return
				}
			}
		}
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := lowBound + rand.Intn(topBound-lowBound+1)

	err := numberOperations.SendResultMessage(&b.cfg, forID, randomNumber)
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
