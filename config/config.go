package config

import (
	"github.com/LucasMateus-eng/simple-bank/utils/logging"
	"github.com/joho/godotenv"
)

var (
	log = logging.NewLogger()
)

func GetConfig() error {
	if err := godotenv.Load("./.env"); err != nil {
		log.Error("config error: ", err.Error())
		return err
	}
	return nil
}
