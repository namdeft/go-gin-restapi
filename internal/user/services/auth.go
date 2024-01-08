package services

import (
	"context"
	"gin-restapi/internal/token"
	"gin-restapi/internal/user/dto"
	"gin-restapi/internal/user/model"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthStorage interface {
	Login(ctx context.Context, data *dto.LoginInput) (*model.User, error)
	Register(ctx context.Context, data *dto.RegisterInput) error
}

type authService struct {
	store AuthStorage
}

func AuthService(store AuthStorage) *authService {
	return &authService{store: store}
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(data *dto.LoginInput, u *model.User) (string, error) {
	err := VerifyPassword(data.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", nil
	}

	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", nil
	}

	return token, nil
}

func (service *authService) Login(ctx context.Context, data *dto.LoginInput) (string, error) {
	u, err := service.store.Login(ctx, data)
	if err != nil {
		return "", err
	}

	tokenKey, err := LoginCheck(data, u)
	if err != nil {
		return "", err
	}

	return tokenKey, nil
}

func BeforeRegister(u *dto.RegisterInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))

	return nil
}

func (service *authService) Register(ctx context.Context, data *dto.RegisterInput) error {
	if err := BeforeRegister(data); err != nil {
		return err
	}

	if err := service.store.Register(ctx, data); err != nil {
		return err
	}

	return nil
}
