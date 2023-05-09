package objects

const NewMessage = "message_new"

type MessageNewObject struct {
	Message Message `json:"message"`
}

type Message struct {
	FromID int    `json:"from_id"`
	Text   string `json:"text"`
}
