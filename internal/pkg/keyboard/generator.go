package keyboard

import (
	"encoding/json"
	"fmt"
)

// GenerateKeyboard2x2 Генерирует клавиатуру размером 2x2 (две строки по две кнопки в каждой)
// и возвращает JSON объект клавиатуры с переданными параметрами
//
// Все кнопки имеют type:"text"
func GenerateKeyboard2x2(buttonsText []string) []byte {
	var buttons [][]button

	for i := 0; i < len(buttonsText); i += 2 {
		var row []button

		row = append(row, button{
			Action: buttonAction{
				Type:    "text",
				Label:   buttonsText[i],
				Payload: fmt.Sprintf("{\"button\":%d}", i),
			},
			Color: "primary",
		})

		if i+1 < len(buttonsText) {
			row = append(row, button{
				Action: buttonAction{
					Type:    "text",
					Label:   buttonsText[i+1],
					Payload: fmt.Sprintf("{\"button\":%d}", i+1),
				},
				Color: "primary",
			})
		}

		buttons = append(buttons, row)
	}

	keyboard := keyboard{
		OneTime: true,
		Buttons: buttons,
	}

	keyboardJSON, _ := json.Marshal(keyboard)

	return keyboardJSON
}

// GenerateKeyboard2x1 Генерирует клавиатуру размером 2x1 (одна строка, две кнопки)
// и возвращает JSON объект клавиатуры с переданными параметрами
//
// Все кнопки имеют type:"text"
func GenerateKeyboard2x1(buttonsText []string) []byte {
	var buttons [][]button

	row := make([]button, 0, 2)

	row = append(row, button{
		Action: buttonAction{
			Type:    "text",
			Label:   buttonsText[0],
			Payload: fmt.Sprintf("{\"button\":%d}", 0),
		},
		Color: "primary",
	})

	row = append(row, button{
		Action: buttonAction{
			Type:    "text",
			Label:   buttonsText[1],
			Payload: fmt.Sprintf("{\"button\":%d}", 1),
		},
		Color: "primary",
	})

	buttons = append(buttons, row)

	keyboard := keyboard{
		OneTime: true,
		Buttons: buttons,
	}

	keyboardJSON, _ := json.Marshal(keyboard)

	return keyboardJSON
}

// GenerateKeyboard3x2 Генерирует клавиатуру размером 3x2 (две строки по три кнопки в каждой)
// и возвращает JSON объект клавиатуры с переданными параметрами
//
// Все кнопки имеют type:"text"
func GenerateKeyboard3x1(buttonsText []string) []byte {
	var buttons [][]button

	row := make([]button, 0, 3)

	for id := range buttonsText {
		row = append(row, button{
			Action: buttonAction{
				Type:    "text",
				Label:   buttonsText[id],
				Payload: fmt.Sprintf("{\"button\":%d}", id),
			},
			Color: "primary",
		})
	}

	buttons = append(buttons, row)

	keyboard := keyboard{
		OneTime: true,
		Buttons: buttons,
	}

	keyboardJSON, _ := json.Marshal(keyboard)

	return keyboardJSON
}
