package services

import (
	"context"
	"gin-restapi/internal/token"
	"gin-restapi/internal/user/dto"
	"gin-restapi/internal/user/model"
	"gin-restapi/internal/user/repository"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, data *dto.RegisterInput) error
	Login(ctx context.Context, data *dto.LoginInput) (string, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *authService {
	return &authService{
		userRepository: userRepo,
	}
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
	u, err := service.userRepository.Login(ctx, data)
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

	if err := service.userRepository.Register(ctx, data); err != nil {
		return err
	}

	return nil
}
