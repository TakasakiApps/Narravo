package config

import (
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"os"
)

var mConfig = &Config{
	Crypto: struct {
		TokenAesKey string `json:"tokenAesKey"`
	}{
		TokenAesKey: utils.GenerateAESKey(),
	},
	TencentCloud: struct {
		SecretId  string `json:"secretId"`
		SecretKey string `json:"secretKey"`
	}{
		SecretId:  os.Getenv("TENCENT_CLOUD_SECRET_ID"),
		SecretKey: os.Getenv("TENCENT_CLOUD_SECRET_KEY"),
	},
}

type Config struct {
	Crypto struct {
		TokenAesKey string `json:"tokenAesKey"`
	} `json:"crypto"`
	TencentCloud struct {
		SecretId  string `json:"secretId"`
		SecretKey string `json:"secretKey"`
	} `json:"tencentCloud"`
}

func GetInstance() *Config {
	return mConfig
}
