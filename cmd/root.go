package cmd

import (
	"github.com/rs/zerolog/log"

	"recsvc/cmd/server"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "A brief description of your application",
	Long:  `The longer description`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(server.Cmd)
}
