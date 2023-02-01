package aggregate

import (
	"errors"
	"fmt"
	"time"

	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
)

type WalletGorm struct {
	ID                     uint      `gorm:"primaryKey;column:id"`
	UUID                   string    `gorm:"uniqueIndex;column:uuid"`
	Name                   string    `gorm:"not null;column:name"`
	PersonalIdentification string    `gorm:"uniqueIndex;column:personal_id"`
	Email                  string    `gorm:"uniqueIndex;column:email"`
	Password               string    `gorm:"not null;column:password"`
	IsAShopkeeper          bool      `gorm:"not null;column:is_a_shopkeeper"`
	Balance                float64   `gorm:"not null;column:balance"`
	CreatedAt              time.Time `gorm:"not null;column:created_at"`
	Entries                []string  `gorm:"column:entries"`
	Transfers              []string  `gorm:"column:transfers"`
}

func (wg *WalletGorm) TableName() string {
	return "wallets"
}

func (wg *WalletGorm) ToAggregate() (*aggregate.Wallet, error) {
	if wg == nil {
		return nil, errors.New("o dto do agregado Wallet não pode ser vazio")
	}

	wallet, err := aggregate.NewWallet(wg.Name, wg.PersonalIdentification, wg.Email, wg.Password, wg.IsAShopkeeper)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func NewRow(wallet aggregate.Wallet) WalletGorm {
	wp := wallet.GetPerson()

	wa := wallet.GetAccount()

	entries := make([]string, len(wallet.GetEntries()))
	for _, entry := range wallet.GetEntries() {
		entries = append(entries, entry.UUID.String())
	}

	transfers := make([]string, len(wallet.GetTransfers()))
	for _, transfer := range wallet.GetTransfers() {
		users := fmt.Sprintf("%s|%s", transfer.FromWalletUUID, transfer.ToWalletUUID)
		transfers = append(transfers, users)
	}

	return WalletGorm{
		UUID:                   wp.UUID.String(),
		Name:                   wp.Name,
		PersonalIdentification: wp.PersonalIdentification,
		Email:                  wp.Email,
		Password:               wp.Password,
		IsAShopkeeper:          wp.IsAShopkeeper,
		Balance:                wa.Balance,
		CreatedAt:              wa.CreatedAt,
		Entries:                entries,
		Transfers:              transfers,
	}
}
