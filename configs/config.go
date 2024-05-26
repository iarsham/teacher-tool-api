package configs

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
	"strings"
	"time"
)

//go:embed config.yaml
var Configurations []byte

type Postgres struct {
	Host            string
	Port            int
	Username        string
	Password        string
	DB              string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
}

type App struct {
	Host    string
	Port    int
	Addr    string
	Debug   bool
	BaseAPI string
}

type Config struct {
	Postgres *Postgres
	App      *App
}

func NewConfig() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ENV")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(Configurations)); err != nil {
		return nil, err
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
