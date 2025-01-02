package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

// Config represents the config of the bot
type Config struct {
	Token        string   `env:"TELEGRAM_TOKEN"`
	Admins       []string `env:"TELEGRAM_ADMINS"`
	FlyOrgToken  string   `env:"FLY_ORG_TOKEN"`
	WebhookURL   string   `env:"WEBHOOK_URL"`
	Port         int      `env:"WEBHOOK_PORT" envDefault:"8443"`
	Debug        bool     `env:"DEBUG" envDefault:"false"`
	WebhookToken string   `env:"WEBHOOK_TOKEN" envDefault:"changeme"`
}

func LoadConfig() *Config {
	c := Config{}
	err := env.ParseWithOptions(&c, env.Options{
		// https://github.com/caarlos0/env?tab=readme-ov-file#parse-options
		Prefix:                "TMM_",
		RequiredIfNoDef:       true,
		UseFieldNameByDefault: true,
	})
	if err != nil {
		log.Fatalf("Error reading config: %q", err)
	}
	return &c
}
