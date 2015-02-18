package server

import (
	"encoding/json"
	"os"
)

type Config struct {
	HTTPAddress string
	APIKey      string
	SecretKey   string
	SessionKey  string
	TemplateDir string
	SiteURL     string
}

func ReadConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
