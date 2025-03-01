package domain

import "time"

type ChatMessage struct {
	ID         int       `json:"id"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

type ChatRepository interface {
	SendMessage(msg *ChatMessage) error

	GetMessages(senderID, receiverID int) ([]*ChatMessage, error)
}
