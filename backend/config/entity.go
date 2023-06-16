package config

import "github.com/TakasakiApps/Narravo/backend/internal/utils"

var mConfig = &Config{
	Crypto: struct {
		TokenAesKey string `json:"tokenAesKey"`
	}{
		TokenAesKey: utils.GenerateAESKey(),
	},
}

type Config struct {
	Crypto struct {
		TokenAesKey string `json:"tokenAesKey"`
	} `json:"crypto"`
}

func GetInstance() *Config {
	return mConfig
}
