package keyboard

import (
	"encoding/json"
	"fmt"
)

// GenerateKeyboard2xn Генерирует клавиатуру размером 2xn (n строк по две кнопки в каждой)
// и возвращает JSON объект клавиатуры с переданными параметрами
//
// Все кнопки имеют type:"text"
func GenerateKeyboard2xn(buttonsText []string) []byte {
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

	for i := 0; i < 3; i++ {
		row := make([]button, 0, 1)
		row = append(row, button{
			Action: buttonAction{
				Type:    "text",
				Label:   buttonsText[i],
				Payload: fmt.Sprintf("{\"button\":%d}", i),
			},
			Color: "primary",
		})
		buttons = append(buttons, row)
	}

	keyboard := keyboard{
		OneTime: true,
		Buttons: buttons,
	}

	keyboardJSON, _ := json.Marshal(keyboard)

	return keyboardJSON
}

// GenerateKeyboard1x1 Генерирует клавиатуру размером 1x1 (одна кнопка)
// и возвращает JSON объект клавиатуры с переданными параметрами
//
// Все кнопки имеют type:"text"
func GenerateKeyboard1x1(buttonsText []string) []byte {
	var buttons [][]button

	for i := 0; i < 1; i++ {
		row := make([]button, 0, 1)
		row = append(row, button{
			Action: buttonAction{
				Type:    "text",
				Label:   buttonsText[i],
				Payload: fmt.Sprintf("{\"button\":%d}", i),
			},
			Color: "primary",
		})
		buttons = append(buttons, row)
	}

	keyboard := keyboard{
		OneTime: true,
		Buttons: buttons,
	}

	keyboardJSON, _ := json.Marshal(keyboard)

	return keyboardJSON
}
