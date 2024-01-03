package business

import (
	"context"
	"gin-restapi/internal/token"
	"gin-restapi/internal/user/model"

	"golang.org/x/crypto/bcrypt"
)

type LoginStorage interface {
	Login(ctx context.Context, data *model.LoginInput) (*model.User, error)
}

type loginBusiness struct {
	store LoginStorage
}

func LoginBusiness(store LoginStorage) *loginBusiness {
	return &loginBusiness{store: store}
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(data *model.LoginInput, u *model.User) (string, error) {
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

func (business *loginBusiness) Login(ctx context.Context, data *model.LoginInput) (string, error) {
	u, err := business.store.Login(ctx, data)
	if err != nil {
		return "", err
	}

	tokenKey, err := LoginCheck(data, u)
	if err != nil {
		return "", err
	}

	return tokenKey, nil
}
