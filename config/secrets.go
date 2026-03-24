package config

import (
	"errors"
	"os"
	"strings"

	"github.com/zalando/go-keyring"
)

const (
	service = "aj-tool"
	user    = "default"
	envVar  = "AJ_API_KEY"
)

func SetAPIKey(key string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return errors.New("empty API key")
	}
	return keyring.Set(service, user, key)
}

func GetAPIKey() (string, error) {
	// 1. env override
	if v := strings.TrimSpace(os.Getenv(envVar)); v != "" {
		return v, nil
	}

	// 2. system keyring
	v, err := keyring.Get(service, user)
	if err == nil {
		return v, nil
	}

	return "", err
}

func DeleteAPIKey() error {
	return keyring.Delete(service, user)
}
