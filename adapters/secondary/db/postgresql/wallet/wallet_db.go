package wallet

import (
	"fmt"

	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
	"github.com/LucasMateus-eng/simple-bank/application/entity"
	secondaryport "github.com/LucasMateus-eng/simple-bank/application/ports/secondary/wallet"
	valueobject "github.com/LucasMateus-eng/simple-bank/application/value_object"
	gormaggregate "github.com/LucasMateus-eng/simple-bank/dto/secondary/aggregate"
	gormentity "github.com/LucasMateus-eng/simple-bank/dto/secondary/entity"
	gormvalueobject "github.com/LucasMateus-eng/simple-bank/dto/secondary/value_object"
	"github.com/LucasMateus-eng/simple-bank/utils/logging"
	"github.com/LucasMateus-eng/simple-bank/utils/parse"
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
	tx := wr.db.Begin()
	defer func() {
		if wr := recover(); wr != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return aggregate.Wallet{}, err
	}

	var walletGorm gormaggregate.WalletGorm
	if err := tx.First(&walletGorm, "uuid = ?", id).Error; err != nil {
		log.Errorf("Erro ao obter carteira %s no repositório PostgreSQL. Detalhes: %s", id.String(), err.Error())
		tx.Rollback()
		return aggregate.Wallet{}, err
	}

	wallet, err := walletGorm.ToAggregate()
	if err != nil {
		log.Errorf("Erro ao converter o dto da carteira %s para o seu agregado. Detalhes: %s", id.String(), err.Error())
		tx.Rollback()
		return aggregate.Wallet{}, err
	}

	if wallet.IsEmpty() {
		err := fmt.Errorf("não foi possível consultar a carteira %s no banco de dados", id)
		log.Errorf("Erro ao obter carteira %s no repositório PostgreSQL. Detalhes: %s", id.String(), err.Error())
		tx.Rollback()
		return aggregate.Wallet{}, err
	}

	entriesGorm := make([]gormentity.EntryGorm, 0)
	if err := tx.Where("wallet_uuid = ?", walletGorm.UUID).Find(&entriesGorm).Error; err != nil {
		log.Errorf("Erro ao obter as entradas da carteira %s no repositório PostgreSQL. Detalhes: %s", walletGorm.UUID, err.Error())
		tx.Rollback()
		return aggregate.Wallet{}, err
	}

	transfersGorm := make([]gormvalueobject.TransferGorm, 0)
	if err := tx.Where("from_wallet_uuid = ? OR to_wallet_uuid = ?", walletGorm.UUID, walletGorm.UUID).Find(&transfersGorm).Error; err != nil {
		log.Errorf("Erro ao obter as transferências da carteira %s no repositório PostgreSQL. Detalhes: %s", walletGorm.UUID, err.Error())
		tx.Rollback()
		return aggregate.Wallet{}, err
	}

	if len(entriesGorm) > 0 {
		for _, entryGorm := range entriesGorm {
			entry, err := entryGorm.ToEntity()
			if err != nil {
				log.Errorf("Erro ao converter o dto da entrada %s para a sua entidade. Detalhes: %s", entryGorm.UUID, err.Error())
				tx.Rollback()
				return aggregate.Wallet{}, err
			}

			wallet.SetEntries(entry)
		}
	}

	if len(transfersGorm) > 0 {
		for _, transferGorm := range transfersGorm {
			transfer, err := transferGorm.ToValueObject()
			if err != nil {
				log.Errorf("Erro ao converter o dto da transferência %d para o seu objeto de valor. Detalhes: %s", transferGorm.ID, err.Error())
				tx.Rollback()
				return aggregate.Wallet{}, err
			}

			wallet.SetTransfers(*transfer)
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return aggregate.Wallet{}, err
	}

	return *wallet, nil
}

func (wr *walletPostgreSQLRepo) Add(person entity.Person) error {
	aggregate, err := aggregate.NewWallet(person.Name, person.PersonalIdentification, person.Email, person.Password, person.IsAShopkeeper)
	if err != nil {
		return err
	}

	wg := gormaggregate.WalletGorm{}
	wg.FromAggregate(aggregate)

	if err := wr.db.Model(&gormaggregate.WalletGorm{}).Create(&wg).Error; err != nil {
		log.Error("Erro ao criar carteira no repositório PostgreSQL: ", err.Error())
		return err
	}

	return nil
}

func (wr *walletPostgreSQLRepo) Update(person entity.Person) error {
	wg := gormaggregate.WalletGorm{
		Name:                   person.Name,
		PersonalIdentification: person.PersonalIdentification,
		Email:                  person.Email,
		Password:               person.Password,
		IsAShopkeeper:          person.IsAShopkeeper,
	}

	if err := wr.db.Model(gormaggregate.WalletGorm{}).Where("uuid = ?", person.UUID.String()).Updates(wg).Error; err != nil {
		log.Errorf("Erro ao atualizar a carteira %s no repositório PostgreSQL. Detalhes: %s", person.UUID.String(), err.Error())
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

func (wr *walletPostgreSQLRepo) Transfer(transfer valueobject.Transfer) error {
	tx := wr.db.Begin()
	defer func() {
		if wr := recover(); wr != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	fromWalletUUID := transfer.FromWalletUUID.String()
	fromWallet := &gormaggregate.WalletGorm{}
	if err := tx.First(fromWallet, "uuid = ?", fromWalletUUID).Error; err != nil {
		log.Errorf("Erro ao obter carteira %s no repositório PostgreSQL. Detalhes: %s", fromWalletUUID, err.Error())
		tx.Rollback()
		return err
	}

	toWalletUUID := transfer.ToWalletUUID.String()
	toWallet := &gormaggregate.WalletGorm{}
	if err := tx.First(toWallet, "uuid = ?", toWalletUUID).Error; err != nil {
		log.Errorf("Erro ao obter carteira %s no repositório PostgreSQL. Detalhes: %s ", toWalletUUID, err.Error())
		tx.Rollback()
		return err
	}

	createdAt := *parse.SetTime()

	entryFromWallet := gormentity.EntryGorm{
		UUID:      uuid.New().String(),
		Owner:     fromWalletUUID,
		Amount:    -transfer.Amount,
		CreatedAt: createdAt,
	}

	entryToWallet := gormentity.EntryGorm{
		UUID:      uuid.New().String(),
		Owner:     toWalletUUID,
		Amount:    transfer.Amount,
		CreatedAt: createdAt,
	}

	fromWallet.Balance -= transfer.Amount
	toWallet.Balance += transfer.Amount

	if err := tx.Where("uuid = ?", fromWalletUUID).Updates(fromWallet).Error; err != nil {
		log.Errorf("Erro ao atualizar a carteira %s no repositório PostgreSQL. Detalhes: %s", fromWalletUUID, err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Where("uuid = ?", toWalletUUID).Updates(toWallet).Error; err != nil {
		log.Errorf("Erro ao atualizar a carteira %s no repositório PostgreSQL. Detalhes: %s", toWalletUUID, err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Create(&entryFromWallet).Error; err != nil {
		log.Errorf("Erro ao criar entrada %s no repositório PostgreSQL. Detalhes: ", entryFromWallet.UUID, err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Create(&entryToWallet).Error; err != nil {
		log.Errorf("Erro ao criar entrada %s no repositório PostgreSQL. Detalhes: ", entryToWallet.UUID, err.Error())
		tx.Rollback()
		return err
	}

	trg := &gormvalueobject.TransferGorm{}
	trg.FromValueObject(transfer)

	if err := tx.Create(trg).Error; err != nil {
		log.Errorf("Erro ao criar transferência %s no repositório PostgreSQL. Detalhes: ", fmt.Sprintf("%s|%s", fromWalletUUID, toWalletUUID), err.Error())
		tx.Rollback()
		return err
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (wr *walletPostgreSQLRepo) Deposit(deposit valueobject.Transfer) error {
	tx := wr.db.Begin()
	defer func() {
		if wr := recover(); wr != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	walletUUID := deposit.FromWalletUUID.String()
	wallet := &gormaggregate.WalletGorm{}
	if err := tx.First(wallet, "uuid = ?", walletUUID).Error; err != nil {
		log.Errorf("Erro ao obter carteira %s no repositório PostgreSQL. Detalhes: %s", walletUUID, err.Error())
		tx.Rollback()
		return err
	}

	createdAt := *parse.SetTime()

	entryWallet := gormentity.EntryGorm{
		UUID:      uuid.New().String(),
		Owner:     walletUUID,
		Amount:    deposit.Amount,
		CreatedAt: createdAt,
	}

	wallet.Balance += deposit.Amount

	if err := tx.Where("uuid = ?", walletUUID).Updates(wallet).Error; err != nil {
		log.Errorf("Erro ao atualizar a carteira %s no repositório PostgreSQL. Detalhes: %s", walletUUID, err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Create(&entryWallet).Error; err != nil {
		log.Errorf("Erro ao criar entrada %s no repositório PostgreSQL. Detalhes: ", entryWallet.UUID, err.Error())
		tx.Rollback()
		return err
	}

	trg := &gormvalueobject.TransferGorm{}
	trg.FromValueObject(deposit)

	if err := tx.Create(trg).Error; err != nil {
		log.Errorf("Erro ao criar transferência %s no repositório PostgreSQL. Detalhes: ", fmt.Sprintf("%s|%s", walletUUID, walletUUID), err.Error())
		tx.Rollback()
		return err
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
