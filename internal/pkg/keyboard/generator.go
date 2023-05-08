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

// GenerateKeyboard3x1 Генерирует клавиатуру размером 3x1 (одна строка, три кнопки)
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

// GenerateKeyboard1x3 Генерирует клавиатуру размером 1x3 (три строки по одной кнопке в каждой)
// и возвращает JSON объект клавиатуры с переданными параметрами
//
// Все кнопки имеют type:"text"
func GenerateKeyboard1x3(buttonsText []string) []byte {
	var buttons [][]button

	row1 := make([]button, 0, 1)
	row2 := make([]button, 0, 1)
	row3 := make([]button, 0, 1)

	row1 = append(row1, button{
		Action: buttonAction{
			Type:    "text",
			Label:   buttonsText[0],
			Payload: fmt.Sprintf("{\"button\":%d}", 0),
		},
		Color: "primary",
	})

	row2 = append(row2, button{
		Action: buttonAction{
			Type:    "text",
			Label:   buttonsText[1],
			Payload: fmt.Sprintf("{\"button\":%d}", 1),
		},
		Color: "primary",
	})

	row3 = append(row3, button{
		Action: buttonAction{
			Type:    "text",
			Label:   buttonsText[2],
			Payload: fmt.Sprintf("{\"button\":%d}", 2),
		},
		Color: "primary",
	})

	buttons = append(buttons, row1)
	buttons = append(buttons, row2)
	buttons = append(buttons, row3)

	keyboard := keyboard{
		OneTime: true,
		Buttons: buttons,
	}

	keyboardJSON, _ := json.Marshal(keyboard)

	return keyboardJSON
}
