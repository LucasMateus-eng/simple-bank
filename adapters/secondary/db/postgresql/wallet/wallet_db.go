package wallet

import (
	"fmt"

	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
	secondaryport "github.com/LucasMateus-eng/simple-bank/application/ports/secondary/wallet"
	gormaggregate "github.com/LucasMateus-eng/simple-bank/dto/secondary/aggregate"
	"github.com/LucasMateus-eng/simple-bank/utils/logging"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	log = logging.NewLogger()
)

type walletPostgreSQLRepo struct {
	db *gorm.DB
}

func NewWalletPostgreSQLRepo(db *gorm.DB) secondaryport.WalletRepository {
	return &walletPostgreSQLRepo{
		db: db,
	}
}

func (wr *walletPostgreSQLRepo) Get(id uuid.UUID) (aggregate.Wallet, error) {
	var wg gormaggregate.WalletGorm
	if err := wr.db.First(&wg, "person_uuid = ?", id).Error; err != nil {
		log.Error("Erro ao obter carteira no repositório PostgreSQL: ", err.Error())
		return aggregate.Wallet{}, err
	}

	wallet, err := wg.ToAggregate()
	if err != nil {
		log.Error("Erro ao converter o dto da carteira para o seu agregado: ", err.Error())
		return aggregate.Wallet{}, err
	}

	if wallet.IsEmpty() {
		err := fmt.Errorf("não foi possível consultar a carteira %s no banco de dados", id)
		log.Error("Erro ao obter carteira no repositório PostgreSQL: ", err.Error())
		return aggregate.Wallet{}, err
	}

	return *wallet, nil
}

func (wr *walletPostgreSQLRepo) Add(wallet aggregate.Wallet) error {
	wg := gormaggregate.NewRow(wallet)

	if err := wr.db.Create(&wg).Error; err != nil {
		log.Error("Erro ao criar carteira no repositório PostgreSQL: ", err.Error())
		return err
	}

	return nil
}

func (wr *walletPostgreSQLRepo) Update(wallet aggregate.Wallet) error {
	wg := gormaggregate.NewRow(wallet)

	if err := wr.db.Where("person_uuid = ?", wallet.GetID()).Updates(&wg).Error; err != nil {
		log.Error("Erro ao atualizar a carteira no repositório PostgreSQL: ", err.Error())
		return err
	}

	return nil
}

func (wr *walletPostgreSQLRepo) Delete(id uuid.UUID) error {
	if err := wr.db.Delete(&gormaggregate.WalletGorm{}, id).Error; err != nil {
		log.Errorf("Erro ao excluir a carteria com o id %s no repositório PostgreSQL: %s", id, err.Error())
		return err
	}

	return nil
}
