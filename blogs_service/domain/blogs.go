package domain

import "time"

type Blog struct {
	ID         int
	Title      string
	Content    string
	AuthorID   int
	AuthorName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type BlogRepository interface {
	CreateBlog(blog *Blog) error
	GetBlogByID(id int) (*Blog, error)
	UpdateBlog(blog *Blog) error
	DeleteBlog(id int) error
	ListBlogs() ([]*Blog, error)
}
