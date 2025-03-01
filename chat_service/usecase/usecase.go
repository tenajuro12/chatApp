package usecase

import (
	"chat_app/chat_service/domain"
	"time"
)

// ChatUsecase contains the business logic for chat operations.
type ChatUsecase struct {
	Repo domain.ChatRepository
}

func (u *ChatUsecase) SendMessage(senderID, receiverID int, content string) error {
	msg := &domain.ChatMessage{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		CreatedAt:  time.Now(),
	}
	return u.Repo.SendMessage(msg)
}

func (u *ChatUsecase) GetMessages(senderID, receiverID int) ([]*domain.ChatMessage, error) {
	return u.Repo.GetMessages(senderID, receiverID)
}
