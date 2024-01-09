package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tommynurwantoro/kafkid/config"
	"github.com/tommynurwantoro/kafkid/internal/bootstrap"
	"github.com/tommynurwantoro/kafkid/internal/pkg/logger"
)

func RunPubisher() *cobra.Command {
	var configFile string
	command := &cobra.Command{
		Use:     "publisher",
		Aliases: []string{"pub"},
		Short:   "Run Publisher Service",
		Run: func(c *cobra.Command, args []string) {
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

			bootstrap := bootstrap.NewBootstrap()
			bootstrap.RunPubisher(conf)
		},
	}
	command.Flags().StringVar(&configFile, "config", "./config.yaml", "Set config file path")

	return command
}

func RunConsumer() *cobra.Command {
	var configFile string
	var topics []string
	command := &cobra.Command{
		Use:     "consumer",
		Aliases: []string{"con"},
		Short:   "Run Consumer Service",
		Run: func(c *cobra.Command, args []string) {
			conf := &config.Configuration{}
			conf.LoadConfig(configFile)
			conf.Consumer.Topics = topics

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

			bootstrap := bootstrap.NewBootstrap()
			bootstrap.RunConsumer(conf)
		},
	}
	command.Flags().StringVar(&configFile, "config", "./config.yaml", "Set config file path")
	command.Flags().StringSliceVar(&topics, "topics", []string{}, "Set kafka topics. Example: --topics=topic1,topic2")
	command.MarkFlagRequired("topics")

	return command
}
