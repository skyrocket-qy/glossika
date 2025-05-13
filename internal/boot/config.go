package boot

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func InitConfig() (err error) {
	viper.SetConfigName(".env") // name of the file (no extension)
	viper.SetConfigType("yaml") // required if extension not in name
	viper.AddConfigPath(".")    // or wherever your file is located

	err = viper.ReadInConfig()
	if err != nil {
		log.Error().Err(err).Msg("error reading config file")
	}

	return nil
}
