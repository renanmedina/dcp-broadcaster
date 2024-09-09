package utils

import (
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
)

const (
	LOG_FORMAT_TEXT = "text"
	LOG_FORMAT_JSON = "json"
)

type Configs struct {
	DB_URL            string
	DISCORD_BOT_TOKEN string
	LOG_FORMAT        string
}

var loadedConfigs *Configs

func init() {
	loadedConfigs = loadConfigs()
}

func (c *Configs) DbConnectionInfo() string {
	return c.DB_URL
}

func GetConfigs() *Configs {
	return loadedConfigs
}

func loadConfigs() *Configs {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return &Configs{
		DB_URL:            os.Getenv("DB_URL"),
		DISCORD_BOT_TOKEN: os.Getenv("DISCORD_BOT_TOKEN"),
		LOG_FORMAT:        os.Getenv("LOG_FORMAT"),
	}
}

func loadB3TokenCached() (string, error) {
	fileContent, err := os.ReadFile("./b3_token_cached.json")

	if err != nil {
		return "", err
	}

	var tokenInfo B3Token
	json.Unmarshal(fileContent, &tokenInfo)
	return "Bearer " + tokenInfo.AccessToken, nil
}

type B3Token struct {
	AccessToken string `json:"access_token"`
}
