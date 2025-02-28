package repository

import (
	"blogs_service/domain"
	"database/sql"
	"errors"
)

type PostgresBlogRepo struct {
	DB *sql.DB
}

func NewPostgresBlogRepo(db *sql.DB) *PostgresBlogRepo {
	return &PostgresBlogRepo{DB: db}
}

func (r *PostgresBlogRepo) CreateBlog(blog *domain.Blog) error {
	query := `
		INSERT INTO blogs (title, content, author_id, author_name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`
	err := r.DB.QueryRow(query, blog.Title, blog.Content, blog.AuthorID, blog.CreatedAt, blog.UpdatedAt).
		Scan(&blog.ID)
	return err
}

func (r *PostgresBlogRepo) GetBlogByID(id int) (*domain.Blog, error) {
	query := `
		SELECT id, title, content, author_id, author_name, created_at, updated_at 
		FROM blogs WHERE id = $1
	`
	blog := &domain.Blog{}
	err := r.DB.QueryRow(query, id).
		Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID, &blog.AuthorName, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("blog not found")
		}
		return nil, err
	}
	return blog, nil
}

func (r *PostgresBlogRepo) UpdateBlog(blog *domain.Blog) error {
	query := `
		UPDATE blogs SET title = $1, content = $2, updated_at = $3 WHERE id = $4
	`
	_, err := r.DB.Exec(query, blog.Title, blog.Content, blog.UpdatedAt, blog.ID)
	return err
}

func (r *PostgresBlogRepo) DeleteBlog(id int) error {
	query := `DELETE FROM blogs WHERE id = $1`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("blog not found")
	}
	return nil
}

func (r *PostgresBlogRepo) ListBlogs() ([]*domain.Blog, error) {
	query := `
		SELECT id, title, content, author_id, author_name, created_at, updated_at FROM blogs
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []*domain.Blog
	for rows.Next() {
		blog := &domain.Blog{}
		err = rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID, &blog.AuthorName, &blog.CreatedAt, &blog.UpdatedAt)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}
