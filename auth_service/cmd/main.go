// main.go
package main

import (
	delivery "chat_app/auth_service/delievery/http"
	"chat_app/auth_service/infrastructure/repository"
	"chat_app/auth_service/usecase"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq" //
)

func main() {
	connStr := "postgres://username:password@localhost:5432/mydatabase?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Не удалось установить соединение с базой данных:", err)
	}

	userRepo := repository.NewPostgresUserRepo(db)

	authUsecase := &usecase.AuthUsecase{
		UserRepo:     userRepo,
		JwtSecretKey: "your-secret-key",
	}

	authHandler := &delivery.AuthHandler{
		AuthUsecase: authUsecase,
	}

	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)

	log.Println("Сервис запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
