package usecase

import (
	"blogs_service/domain"
	"errors"
	"time"
)

type BlogUsecase struct {
	Repo domain.BlogRepository
}

func (u *BlogUsecase) CreateBlog(title, content, authorName string, authorID int) error {
	if title == "" || content == "" {
		return errors.New("title and content cannot be empty")
	}
	blog := &domain.Blog{
		Title:      title,
		Content:    content,
		AuthorName: authorName,
		AuthorID:   authorID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return u.Repo.CreateBlog(blog)
}

func (u *BlogUsecase) GetBlogById(id int) (*domain.Blog, error) {
	return u.Repo.GetBlogByID(id)
}

func (u *BlogUsecase) DeleteBlog(id int) error {
	return u.Repo.DeleteBlog(id)
}

func (u *BlogUsecase) ListBlogs() ([]*domain.Blog, error) {
	return u.Repo.ListBlogs()
}

func (u *BlogUsecase) UpdateBlog(id int, title, content string) error {
	blog, err := u.Repo.GetBlogByID(id)
	if err != nil {
		return err
	}
	if title != "" {
		blog.Title = title
	}
	if content != "" {
		blog.Content = content
	}
	blog.UpdatedAt = time.Now()
	return u.Repo.UpdateBlog(blog)

}
