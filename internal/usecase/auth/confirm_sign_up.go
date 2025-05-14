package authucase

import (
	"context"

	"recsvc/internal/domain/er"
	"recsvc/internal/model"
	validate "recsvc/internal/service/validator"

	"github.com/redis/go-redis/v9"
)

type ConfirmSignUpIn struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code"  validate:"required"`
}

func (u *Usecase) ConfirmSignUp(c context.Context, in ConfirmSignUpIn) error {
	if err := validate.Get().Struct(in); err != nil {
		return er.W(err, er.ValidateInput)
	}

	val, err := u.redisSvc.Cli.Get(c, in.Email).Result()
	if err != nil {
		if err == redis.Nil {
			return er.NewAppErr(er.NotFound)
		}
		return er.W(err)
	}

	if val != in.Code {
		return er.NewAppErr(er.InvalidCode)
	}

	if err := u.db.WithContext(c).
		Model(model.User{}).
		Where("email = ?", in.Email).
		Update("confirmed", true).Error; err != nil {
		return er.W(err)
	}

	if err := u.redisSvc.Cli.Del(c, in.Email).Err(); err != nil {
		return er.W(err)
	}

	return nil
}
