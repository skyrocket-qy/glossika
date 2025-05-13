package boot

import (
	"github.com/rs/zerolog/log"
	validate "glossika/internal/service/validator"
)

func NewService() error {
	log.Info().Msg("InitService")

	validate.New()

	return nil
}
