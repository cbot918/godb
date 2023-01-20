package config

import (
	"fmt"
	u "godb/util"

	"github.com/spf13/viper"
)

type Config struct {
	// Server struct {
	// 	Name string
	// 	Port string
	// 	Host string
	// }
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

func New() (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("test")
	v.AddConfigPath("./test/service/configs")
	// v.AutomaticEnv()
	// v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
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
