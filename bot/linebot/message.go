package linebot

type MessageInterface interface {
	GetType() string
}

func (e Message) GetType() string {
	return e.Type
}

type Message struct {
	Type string `json:"type"`
}

type TextMessage struct {
	Message
	Text string `json:"text"`
}

type SentMessage struct {
	Id         string `json:"id"`
	QuoteToken string `json:"quoteToken,omitempty"`
}

func NewTextMessage(text string) *TextMessage {
	return &TextMessage{
		Message: Message{
			Type: "text",
		},
		Text: text,
	}
}
