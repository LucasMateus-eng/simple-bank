package wallet

import (
	"fmt"

	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
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

	fromWallet.Entries = append(fromWallet.Entries, entryFromWallet.UUID)
	toWallet.Entries = append(toWallet.Entries, entryToWallet.UUID)

	fromWallet.Transfers = append(fromWallet.Transfers, fmt.Sprintf("%s|%s", fromWalletUUID, toWalletUUID))
	toWallet.Transfers = append(toWallet.Transfers, fmt.Sprintf("%s|%s", fromWalletUUID, toWalletUUID))

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
