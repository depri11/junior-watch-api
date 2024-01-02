package config

import (
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/depri11/junior-watch-api/pkg/constants"
	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "User service config path")
}

type Config struct {
	ServiceName string         `mapstructure:"service_name"`
	Logger      *logger.Config `mapstructure:"logger"`
	GRPCServer  GRPCServer     `mapstructure:"grpc_server"`
	Database    Database       `mapstructure:"database"`
}

type GRPCServer struct {
	Port              string        `mapstructure:"port"`
	Development       bool          `mapstructure:"development"`
	Timeout           time.Duration `mapstructure:"timeout"`
	ReadTimeout       time.Duration `mapstructure:"read_timeout"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout"`
	MaxConnectionIdle time.Duration `mapstructure:"max_connection_idle"`
	MaxConnectionAge  time.Duration `mapstructure:"max_connection_age"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl_mode"`
	DBDriver string `mapstructure:"db_driver"`
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
			fileName := "config.yaml"
			if flag.Lookup("test.v") != nil {
				getwd = filepath.Join(getwd, "..", "..")
				fileName = "config_test.yaml"
			}
			configPath = filepath.Join(getwd, "user_service", "config", fileName)
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

	grpcPort := os.Getenv(constants.GrpcPort)
	if grpcPort != "" {
		cfg.GRPCServer.Port = grpcPort
	}

	return cfg, nil
}
