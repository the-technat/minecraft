package config

import (
	"log"
	"slices"

	"github.com/spf13/viper"
)

// Config represents the config of the bot
type Config struct {
	Debug        bool     `mapstructure:"DEBUG"`
	Token        string   `mapstructure:"TOKEN"`
	Admins       []string `mapstructure:"ADMINS"`
	Port         int      `mapstructure:"PORT"`
	WebhookURL   string   `mapstructure:"WEBHOOK_URL"`
	WebhookToken string   `mapstructure:"WEBHOOK_TOKEN"`
}

func LoadConfig() *Config {
	viper.SetConfigName("")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("tmm")

	c := Config{}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("Error reading config: %q", err)
	}

	return &c

}

// IsAdmin returns true if the given username is in the allowed admins list
func (c *Config) IsAdmin(user string) bool {
	if slices.Contains(c.Admins, user) {
		return true
	} else {
		return false
	}
}
