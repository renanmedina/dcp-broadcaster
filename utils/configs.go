package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	LOG_FORMAT_TEXT   = "text"
	LOG_FORMAT_JSON   = "json"
	DEFAULT_IMAP_PORT = 993
)

type Configs struct {
	DB_URL            string
	DISCORD_BOT_TOKEN string
	LOG_FORMAT        string
	imapConfigs       *ImapConfigs
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

func GetImapConfigs() *ImapConfigs {
	return GetConfigs().imapConfigs
}

func loadConfigs() *Configs {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	receiverServerPort, err := strconv.Atoi(os.Getenv("RECEIVER_SERVER_PORT"))

	if err != nil {
		receiverServerPort = DEFAULT_IMAP_PORT // default
	}

	return &Configs{
		DB_URL:            os.Getenv("DB_URL"),
		DISCORD_BOT_TOKEN: os.Getenv("DISCORD_BOT_TOKEN"),
		LOG_FORMAT:        os.Getenv("LOG_FORMAT"),
		imapConfigs: &ImapConfigs{
			ServerUrl:  os.Getenv("RECEIVER_SERVER"),
			ServerPort: receiverServerPort,
			Username:   os.Getenv("RECEIVER_USERNAME"),
			Password:   os.Getenv("RECEIVER_PASSWORD"),
		},
	}
}
