package http

import (
	"chat_app/chat_service/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type ChatHandler struct {
	ChatUsecase *usecase.ChatUsecase
}

type sendMessageRequest struct {
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Content    string `json:"content"`
}

func (h *ChatHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var req sendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := h.ChatUsecase.SendMessage(req.SenderID, req.ReceiverID, req.Content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Message sent successfully"))
}

func (h *ChatHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	senderIDStr := r.URL.Query().Get("sender_id")
	receiverIDStr := r.URL.Query().Get("receiver_id")
	if senderIDStr == "" || receiverIDStr == "" {
		http.Error(w, "sender_id and receiver_id are required", http.StatusBadRequest)
		return
	}
	senderID, err := strconv.Atoi(senderIDStr)
	if err != nil {
		http.Error(w, "Invalid sender_id", http.StatusBadRequest)
		return
	}
	receiverID, err := strconv.Atoi(receiverIDStr)
	if err != nil {
		http.Error(w, "Invalid receiver_id", http.StatusBadRequest)
		return
	}
	messages, err := h.ChatUsecase.GetMessages(senderID, receiverID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
