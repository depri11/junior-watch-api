package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/depri11/junior-watch-api/pkg/constants"
	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "API Gateway microservice config path")
}

type Config struct {
	ServiceName string         `mapstructure:"serviceName"`
	Logger      *logger.Config `mapstructure:"logger"`
	Http        Http           `mapstructure:"http"`
	Grpc        Grpc           `mapstructure:"grpc"`
}

type Http struct {
	Port                string   `mapstructure:"port"`
	Development         bool     `mapstructure:"development"`
	BasePath            string   `mapstructure:"basePath"`
	UsersPath           string   `mapstructure:"usersPath"`
	DebugHeaders        bool     `mapstructure:"debugHeaders"`
	HttpClientDebug     bool     `mapstructure:"httpClientDebug"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
}

type Grpc struct {
	UserServicePort string `mapstructure:"userServicePort"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/api_gateway_service/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	httpPort := os.Getenv(constants.HttpPort)
	if httpPort != "" {
		cfg.Http.Port = httpPort
	}

	userServicePort := os.Getenv(constants.UserServicePort)
	if userServicePort != "" {
		cfg.Grpc.UserServicePort = userServicePort
	}

	return cfg, nil

}
