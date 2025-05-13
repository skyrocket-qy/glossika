package authucase

import (
	"context"
	"regexp"
	"strings"
	"time"

	"glossika/internal/domain/er"
	"glossika/internal/model"
	validate "glossika/internal/service/validator"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SignUpIn struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=16"`
}

func (u *Usecase) SignUp(c context.Context, in SignUpIn) error {
	if err := validate.Get().Struct(in); err != nil {
		return er.W(err, er.ValidateInput)
	}

	if !isValidPassword(in.Password) {
		return er.NewAppErr(er.InvalidPassword)
	}

	// check if user already exists
	tx := u.db.WithContext(c).Where("email = ?", in.Email).Take(&model.User{})
	if err := tx.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return er.NewAppErr(er.Unknown)
		}
	}
	if tx.RowsAffected > 0 {
		return er.NewAppErr(er.AlreadyExists)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return er.W(err)
	}

	user := model.User{
		Email:    in.Email,
		Password: string(hashedPass),
	}
	if err := u.db.WithContext(c).Create(&user).Error; err != nil {
		return er.W(err)
	}

	OTPCode := "1234"
	if err := u.storeCode(in.Email, OTPCode); err != nil {
		return er.W(err)
	}

	if err := sendEmail(in.Email, OTPCode); err != nil {
		return er.W(err)
	}

	return nil
}

func containsSpecial(s string) bool {
	special := `()[]{}<>+\-*/?,.:;"'_\\|~` + "`" + `!@#$%^&=`
	for _, r := range s {
		if strings.ContainsRune(special, r) {
			return true
		}
	}
	return false
}

func isValidPassword(pw string) bool {
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pw)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(pw)

	return hasUpper && hasLower && containsSpecial(pw)
}

func sendEmail(email string, code string) error {
	return nil
}

func (u *Usecase) storeCode(email string, code string) error {
	return u.redisSvc.Cli.Set(context.Background(), email, code, 10*time.Minute).Err()
}
