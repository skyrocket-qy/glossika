package authucase

import (
	"context"
	"errors"

	"recsvc/internal/domain/er"
	"recsvc/internal/model"
	validate "recsvc/internal/service/validator"
	"recsvc/internal/util"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginIn struct {
	Email    string `json:"email"    validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginOut struct {
	AccessToken string `json:"accessToken"`
}

func (u *Usecase) Login(c context.Context, in LoginIn) (*LoginOut, error) {
	if err := validate.Get().Struct(in); err != nil {
		return nil, er.W(err, er.ValidateInput)
	}

	// Find user by email
	var user model.User
	if err := u.db.WithContext(c).
		Where("email = ?", in.Email).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, er.NewAppErr(er.Unauthorized)
		}
		return nil, er.W(err)
	}

	if !user.Confirmed {
		return nil, er.NewAppErr(er.Unauthorized)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, er.NewAppErr(er.Unauthorized)
	}

	// Generate access token
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		return nil, er.W(err)
	}

	return &LoginOut{AccessToken: token}, nil
}
