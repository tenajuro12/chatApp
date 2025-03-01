package repository

import (
	"chat_app/chat_service/domain"
	"database/sql"
	"errors"
)

type PostgresChatRepo struct {
	DB *sql.DB
}

func NewPostgresChatRepo(db *sql.DB) *PostgresChatRepo {
	return &PostgresChatRepo{DB: db}
}

func (r *PostgresChatRepo) SendMessage(msg *domain.ChatMessage) error {
	query := `INSERT INTO chat_messages (sender_id, receiver_id, content, created_at)
	          VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.DB.QueryRow(query, msg.SenderID, msg.ReceiverID, msg.Content, msg.CreatedAt).
		Scan(&msg.ID)
	return err
}

func (r *PostgresChatRepo) GetMessages(senderID, receiverID int) ([]*domain.ChatMessage, error) {
	query := `SELECT id, sender_id, receiver_id, content, created_at 
	          FROM chat_messages 
	          WHERE sender_id = $1 AND receiver_id = $2
	          ORDER BY created_at ASC`
	rows, err := r.DB.Query(query, senderID, receiverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*domain.ChatMessage
	for rows.Next() {
		msg := &domain.ChatMessage{}
		if err := rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if len(messages) == 0 {
		return nil, errors.New("no messages found")
	}
	return messages, nil
}
