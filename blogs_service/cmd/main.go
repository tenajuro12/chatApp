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
	connStr := "postgres://postgres:murderpe@blog_db:5432/blogdb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	blogRepo := repository.NewPostgresBlogRepo(db)
	blogUsecase := &usecase.BlogUsecase{
		Repo: blogRepo,
	}
	blogHandler := http2.BlogHandler{
		BlogUsecase: blogUsecase,
	}

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
