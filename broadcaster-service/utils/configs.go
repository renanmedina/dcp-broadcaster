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
	ENVIRONMENT                  string
	DB_URL                       string
	DISCORD_BOT_TOKEN            string
	WHATSAPP_API_URL             string
	WHATSAPP_API_TOKEN           string
	WHATSAPP_SESSION_ID          string
	LOG_FORMAT                   string
	MIGRATIONS_PATH              string
	SERVICE_NAME                 string
	OPEN_TELEMETRY_COLLECTOR_URL string
	OLLAMA_SERVICE_API_URL       string
	imapConfigs                  *ImapConfigs
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
	if err != nil && os.Getenv("ENVIRONMENT") == "" {
		panic(err.Error())
	}

	receiverServerPort, err := strconv.Atoi(os.Getenv("RECEIVER_SERVER_PORT"))

	if err != nil {
		receiverServerPort = DEFAULT_IMAP_PORT // default
	}

	return &Configs{
		ENVIRONMENT:                  os.Getenv("ENVIRONMENT"),
		DB_URL:                       os.Getenv("DB_URL"),
		DISCORD_BOT_TOKEN:            os.Getenv("DISCORD_BOT_TOKEN"),
		WHATSAPP_API_URL:             os.Getenv("WHATSAPP_API_URL"),
		WHATSAPP_API_TOKEN:           os.Getenv("WHATSAPP_API_TOKEN"),
		WHATSAPP_SESSION_ID:          os.Getenv("WHATSAPP_SESSION_ID"),
		LOG_FORMAT:                   os.Getenv("LOG_FORMAT"),
		MIGRATIONS_PATH:              os.Getenv("MIGRATIONS_PATH"),
		SERVICE_NAME:                 os.Getenv("SERVICE_NAME"),
		OPEN_TELEMETRY_COLLECTOR_URL: os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"),
		OLLAMA_SERVICE_API_URL:       os.Getenv("OLLAMA_SERVICE_API_URL"),
		imapConfigs: &ImapConfigs{
			ServerUrl:  os.Getenv("RECEIVER_SERVER"),
			ServerPort: receiverServerPort,
			Username:   os.Getenv("RECEIVER_USERNAME"),
			Password:   os.Getenv("RECEIVER_PASSWORD"),
		},
	}
}
