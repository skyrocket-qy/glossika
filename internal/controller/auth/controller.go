package authcontroller

import (
	"context"

	authucase "glossika/internal/usecase/auth"
)

type Usecase interface {
	Login(c context.Context, in authucase.LoginIn) (*authucase.LoginOut, error)
	SignUp(c context.Context, in authucase.SignUpIn) error
	ConfirmSignUp(c context.Context, in authucase.ConfirmSignUpIn) error
}

type Handler struct {
	Usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
