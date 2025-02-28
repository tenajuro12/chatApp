// main.go
package main

import (
	http2 "blogs_service/delievery/http"
	"blogs_service/infrastructure/repository"
	"blogs_service/usecase"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:murderpe@localhost:5432/users?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()
	// ... additional code
	// Verify the connection.
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// Initialize repository, use case, and HTTP handler.
	blogRepo := repository.NewPostgresBlogRepo(db)
	blogUsecase := &usecase.BlogUsecase{
		Repo: blogRepo,
	}
	blogHandler := http2.BlogHandler{
		BlogUsecase: blogUsecase,
	}

	// Register HTTP endpoints.
	http.HandleFunc("/blogs/create", blogHandler.CreateBlog)
	http.HandleFunc("/blogs/get", blogHandler.GetBlog)
	http.HandleFunc("/blogs/update", blogHandler.UpdateBlog)
	http.HandleFunc("/blogs/delete", blogHandler.DeleteBlog)
	http.HandleFunc("/blogs/list", blogHandler.ListBlogs)

	log.Println("Blog service is running on port 8080")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
