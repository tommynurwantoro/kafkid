package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/tommynurwantoro/kafkid/config"
	"github.com/tommynurwantoro/kafkid/internal/bootstrap"
	"github.com/tommynurwantoro/kafkid/internal/pkg/logger"
)

var (
	configFile string
	command    = &cobra.Command{
		Use:     "service",
		Aliases: []string{"svc"},
		Short:   "Run service",
		Run: func(c *cobra.Command, args []string) {
			// Load env variable
			err := godotenv.Load(".env")
			if err != nil {
				fmt.Println("Fatal error loading .env file.\n", err)
				os.Exit(1)
			}

			conf := &config.Configuration{}
			conf.LoadConfig(configFile)

			// Load logger
			loggerConfig := logger.Config{
				App:           conf.App,
				AppVer:        conf.AppVer,
				Env:           conf.Env,
				FileLocation:  conf.Logger.FileLocation,
				FileMaxSize:   conf.Logger.FileMaxAge,
				FileMaxBackup: conf.Logger.FileMaxBackup,
				FileMaxAge:    conf.Logger.FileMaxAge,
				Stdout:        conf.Logger.Stdout,
			}

			logger.Load(loggerConfig)
			bootstrap.Run(conf)
		},
	}
)

func GetCommand() *cobra.Command {
	command.Flags().StringVar(&configFile, "config", "./config.sample.yaml", "Set config file path")

	return command
}
