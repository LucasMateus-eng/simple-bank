package aggregate

import (
	"errors"
	"time"

	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
	"github.com/LucasMateus-eng/simple-bank/application/entity"
	"github.com/google/uuid"
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

	parsed, err := uuid.Parse(wg.UUID)
	if err != nil {
		return nil, err
	}

	wallet.SetID(parsed)

	account := entity.Account{
		Owner:     parsed,
		Balance:   wg.Balance,
		CreatedAt: wg.CreatedAt,
	}

	wallet.SetAccount(&account)

	return &wallet, nil
}

func (wg *WalletGorm) FromAggregate(wallet aggregate.Wallet) {
	if wg == nil {
		wg = &WalletGorm{}
	}

	wp := wallet.GetPerson()

	wa := wallet.GetAccount()

	wg.UUID = wp.UUID.String()
	wg.Name = wp.Name
	wg.PersonalIdentification = wp.PersonalIdentification
	wg.Email = wp.Email
	wg.Password = wp.Password
	wg.IsAShopkeeper = wp.IsAShopkeeper
	wg.Balance = wa.Balance
	wg.CreatedAt = wa.CreatedAt
}
