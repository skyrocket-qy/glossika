package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

var validatr *validator.Validate

func Get() *validator.Validate {
	return validatr
}

func New() {
	log.Info().Msg("InitValidator")
	validatr = validator.New(validator.WithRequiredStructEnabled())
}
