package api

import "encoding/json"

type LongPollServer struct {
	Response struct {
		Server string `json:"server"`
		Key    string `json:"key"`
		Ts     string `json:"ts"`
	} `json:"response"`
}

type LongPollResponse struct {
	Ts      string   `json:"ts"`
	Updates []Update `json:"updates"`
}
type Update struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
}
