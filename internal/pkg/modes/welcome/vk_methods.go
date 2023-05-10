package welcome

import (
	"VK-bot/internal/config"
	"VK-bot/internal/pkg/keyboard"
	"VK-bot/internal/pkg/modes/coin"
	"VK-bot/internal/pkg/modes/dice"
	"VK-bot/internal/pkg/modes/number"
	"VK-bot/internal/pkg/modes/room"
	"VK-bot/internal/pkg/modes/word"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func SendWelcomeMessage(cfg *config.Config, peerID int) error {
	rand.Seed(time.Now().UnixNano())
	max := int32(math.MaxInt32)
	randomInt := rand.Int31n(max)

	buttonsText := []string{coin.TossACoin, dice.TossADice, word.GetAWord, number.GetANumber, room.EnterRoom}
	keyboardJSON := keyboard.GenerateKeyboard2xn(buttonsText)

	text := `
Привет!

Если у тебя возникла необходимость придумать что-то случайное, или вести счет некоторой игры, то данный бот может тебе в этом помочь.

Ты можешь:
1. Подбросить монетку (Орел\Решка)
2. Подбросить кубик(-и) (От 1 до 3)
3. Получить случайное английское слово (Сущ.\Прил.\Животное)
4. Получить случайное число в диапазоне
5. Создать или присоединиться к существующей комнате и вести в ней счет игры
	`

	urlToSend := fmt.Sprintf("%smessages.send?peer_id=%d&message=%s&group_id=%s&random_id=%d&access_token=%s&v=%s&keyboard=%s", cfg.ApiURL, peerID, url.QueryEscape(text), cfg.GroupID, randomInt, cfg.Token, cfg.ApiVer, url.QueryEscape(string(keyboardJSON)))
	_, err := http.Get(urlToSend)
	if err != nil {
		return err
	}
	return nil
}
