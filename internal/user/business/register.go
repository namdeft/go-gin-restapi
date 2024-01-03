package business

import (
	"context"
	"gin-restapi/internal/user/model"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type RegisterStorage interface {
	Register(ctx context.Context, data *model.RegisterInput) error
}

type registerBusiness struct {
	store RegisterStorage
}

func RegisterBusiness(store RegisterStorage) *registerBusiness {
	return &registerBusiness{store: store}
}

func BeforeRegister(u *model.RegisterInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))

	return nil
}

func (business *registerBusiness) Register(ctx context.Context, data *model.RegisterInput) error {
	if err := BeforeRegister(data); err != nil {
		return err
	}

	if err := business.store.Register(ctx, data); err != nil {
		return err
	}

	return nil
}
