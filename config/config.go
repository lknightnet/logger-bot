package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App `yaml:"app"`
		TG  `yaml:"tg"`
		API `yaml:"api"`
	}
	App struct {
		Name string `yaml:"name" env:"APP_NAME"`
	}

	TG struct {
		Token string `yaml:"token" env:"TG_TOKEN"`
	}
	API struct {
		UrlAPI string `yaml:"url" env:"API_URL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, err
}
