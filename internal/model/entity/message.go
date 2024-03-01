package entity

import (
	"time"
)

type Message struct {
	ConversationId string    `bson:"conversationId"`
	Sender         string    `bson:"sender"`
	SentAt         time.Time `bson:"sentAt"`
	Text           string    `bson:"text"`
}
