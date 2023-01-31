package main

import (
	"fmt"
	"os"

	restapiv1 "github.com/LucasMateus-eng/simple-bank/adapters/primary/api/rest/router/v1"
	wallethandler "github.com/LucasMateus-eng/simple-bank/adapters/primary/api/rest/wallet"
	walletservice "github.com/LucasMateus-eng/simple-bank/application/services/wallet"
	"github.com/LucasMateus-eng/simple-bank/config"
	psqldatabase "github.com/LucasMateus-eng/simple-bank/infrastructure/db/postgresql"
	"github.com/LucasMateus-eng/simple-bank/utils/logging"
)

var (
	log = logging.NewLogger()
)

func init() {
	if err := config.GetConfig(); err != nil {
		log.Fatal("Config error : ", err.Error())
	}
}

func main() {
	walletRepo, err := psqldatabase.InitPostgreSQLRepo()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("the database has been connected")

	walletService := walletservice.NewWalletService(walletRepo)
	handler := wallethandler.NewWalletHandler(walletService)

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	app := restapiv1.SetupEchoRouter(handler)
	app.Logger.Fatal(app.Start(port))
}
