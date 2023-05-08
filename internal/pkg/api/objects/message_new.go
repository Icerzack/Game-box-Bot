package objects

const NewMessage = "message_new"

type MessageNewObject struct {
	Message struct {
		FromID int    `json:"from_id"`
		Text   string `json:"text"`
	} `json:"message"`
}
