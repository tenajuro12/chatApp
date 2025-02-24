package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"myapp/domain"
)

type AuthUsecase struct {
	UserRepo     domain.UserRepository
	JwtSecretKey string
}

func (a *AuthUsecase) Register(username, password string) error {
	_, err := a.UserRepo.GetUserByUsername(username)
	if err == nil {
		return errors.New("пользователь уже существует")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return a.UserRepo.CreateUser(user)
}

func (a *AuthUsecase) Login(username, password string) (string, error) {
	user, err := a.UserRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("неверное имя пользователя или пароль")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("неверное имя пользователя или пароль")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.JwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
