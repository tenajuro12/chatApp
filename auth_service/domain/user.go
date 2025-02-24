package domain

type User struct {
	ID       int
	Username string
	Password string
}

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
}
