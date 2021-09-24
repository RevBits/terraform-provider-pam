package pam

import (
	"github.com/RevBits/pam-api-go/client"
)

type Config struct {
	Domain string
	SecretID string
	ApiKey string
}

func (config *Config) GetPamSecretValue() (interface{}, error) {
	secretValue, err := client.PamSecretValue(config.Domain, config.SecretID, config.ApiKey)
	return secretValue, err
}

