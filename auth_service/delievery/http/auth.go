package delivery

import (
	"chat_app/auth_service/usecase"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	AuthUsecase *usecase.AuthUsecase
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "неверные данные", http.StatusBadRequest)
		return
	}

	err = h.AuthUsecase.Register(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("пользователь зарегистрирован"))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "неверные данные", http.StatusBadRequest)
		return
	}

	token, err := h.AuthUsecase.Login(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
