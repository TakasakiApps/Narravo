package config

import "github.com/TakasakiApps/Narravo/backend/internal/utils"

var mConfig = &Config{
	Crypto: struct {
		AesKey string `json:"aesKey"`
	}{
		AesKey: utils.GenerateAESKey(),
	},
}

type Config struct {
	Crypto struct {
		AesKey string `json:"aesKey"`
	} `json:"crypto"`
}

func GetInstance() *Config {
	return mConfig
}
