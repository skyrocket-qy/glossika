package boot

import (
	validate "recsvc/internal/service/validator"

	"github.com/rs/zerolog/log"
)

func NewService() error {
	log.Info().Msg("InitService")

	validate.New()

	return nil
}
