package config

import (
	"fmt"
	u "godb/util"

	"github.com/spf13/viper"
)

type Config struct {
	Connections struct {
		Postgresql struct {
			DB       string
			Host     string
			Port     int
			User     string
			Password string
			MaxOpen  int
			MaxIdle  int
		}
	}
}

func New(folderPath string, fileName string, fileType string) (*Config, error) {
	v := viper.New()
	v.AddConfigPath(folderPath)
	v.SetConfigName(fileName)
	v.SetConfigType(fileType)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config error: %s", err)
	}

	u.Logg(v.Get("connection"))

	cfg := &Config{}
	err := v.Unmarshal(&cfg)

	if err != nil {
		return nil, fmt.Errorf("marshal error: %s", err)
	}

	return cfg, nil
}
