package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

const envPrefix = "sample_server"

type SampleServerConfig struct {
	JsonLog    bool         `default:"true" split_words:"true"`
	LogLevel   logrus.Level `default:"info" split_words:"true"`
	ServerAddr string       `default:":4000" split_words:"true"`
}

func ReadConfig() (*SampleServerConfig, error) {
	var s SampleServerConfig
	err := envconfig.Process(envPrefix, &s)
	if err != nil {
		return nil, fmt.Errorf("read env: %w", err)
	}
	return &s, nil
}

func PrintHelp() error {
	err := envconfig.Usage(envPrefix, &SampleServerConfig{})
	return err
}
