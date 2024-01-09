package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	App      string
	AppVer   string
	Env      string
	Http     HttpConfig
	Logger   LogConfig
	Kafka    KafkaConfig
	Consumer ConsumerConfig
}

type HttpConfig struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

type KafkaConfig struct {
	Addresses []string
}

type ConsumerConfig struct {
	GroupID          string
	Topics           []string
	OffsetFromNewest bool
	Output           struct {
		Stdout       bool
		FileLocation string
	}
}

type LogConfig struct {
	FileLocation  string
	FileMaxSize   int
	FileMaxBackup int
	FileMaxAge    int
	Stdout        bool
}

func (c *Configuration) LoadConfig(path string) {
	viper.AddConfigPath(filepath.Dir(path))
	viper.SetConfigName(filepath.Base(path))
	viper.SetConfigType(filepath.Ext(path)[1:])

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}

	err = viper.Unmarshal(c)
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}
