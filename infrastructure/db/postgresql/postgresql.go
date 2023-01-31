package postgresql

import (
	"fmt"
	"os"
	"time"

	walletpostgres "github.com/LucasMateus-eng/simple-bank/adapters/secondary/db/postgresql/wallet"
	secondaryport "github.com/LucasMateus-eng/simple-bank/application/ports/secondary/wallet"
	"github.com/LucasMateus-eng/simple-bank/utils/logging"
	"github.com/LucasMateus-eng/simple-bank/utils/parse"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	log = logging.NewLogger()
)

func buildPostgreSQLConnDSN() string {
	host := os.Getenv("PG_HOST")
	username := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DATABASE")
	port := os.Getenv("PG_PORT")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, username, password, dbName, port)
}

func initPostgreSQL() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(buildPostgreSQLConnDSN()), &gorm.Config{
		NowFunc: func() time.Time {
			return *parse.SetTime()
		},
	})
	if err != nil {
		log.Error("Erro ao inicializar banco de dados postgres: ", err.Error())
		return nil, err
	}

	return db, nil
}

func InitPostgreSQLRepo() (secondaryport.WalletRepository, error) {
	repo, err := initPostgreSQL()
	if err != nil {
		log.Error("Erro ao inicializar repositório postgres: ", err.Error())
		return nil, err
	}

	return walletpostgres.NewWalletPostgreSQLRepo(repo), nil
}
