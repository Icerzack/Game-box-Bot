package keyboard

type button struct {
	Action buttonAction `json:"action"`
	Color  string       `json:"color,omitempty"`
}

type buttonAction struct {
	Type    string `json:"type"`
	Label   string `json:"label"`
	Payload string `json:"payload"`
}

type keyboard struct {
	OneTime bool       `json:"one_time"`
	Buttons [][]button `json:"buttons"`
}
