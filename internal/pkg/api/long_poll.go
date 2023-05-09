package api

import "encoding/json"

type LongPollServer struct {
	Response LongPollServerResponse `json:"response"`
}

type LongPollServerResponse struct {
	Server string `json:"server"`
	Key    string `json:"key"`
	Ts     string `json:"ts"`
}

type LongPollUpdateResponse struct {
	Ts      string   `json:"ts"`
	Updates []Update `json:"updates"`
}

type Update struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
}
