package aggregate

import (
	"errors"

	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
	"github.com/LucasMateus-eng/simple-bank/application/entity"
	valueobject "github.com/LucasMateus-eng/simple-bank/application/value_object"
	gormentity "github.com/LucasMateus-eng/simple-bank/dto/secondary/entity"
	gormvalueobject "github.com/LucasMateus-eng/simple-bank/dto/secondary/value_object"
)

type WalletGorm struct {
	ID               uint                           `gorm:"primaryKey;auto_increment;column:id"`
	PersonUUID       string                         `gorm:"uniqueIndex;column:person_uuid"`
	Person           *gormentity.PersonGorm         `gorm:"foreignKey:PersonUUID;references:UUID;constraint:OnDelete:CASCADE"`
	AccountUUID      string                         `gorm:"uniqueIndex;column:account_uuid"`
	Account          *gormentity.AccountGorm        `gorm:"foreignKey:AccountUUID;references:UUID;constraint:OnDelete:CASCADE"`
	EntryAccountUUID string                         `gorm:"column:entry_account_uuid"`
	Entries          []gormentity.EntryGorm         `gorm:"foreignKey:UUID;references:EntryAccountUUID"`
	TransferGormID   uint                           `gorm:"column:transfer_gorm_id"`
	Transfers        []gormvalueobject.TransferGorm `gorm:"foreignKey:ID;references:TransferGormID"`
}

func (wg *WalletGorm) TableName() string {
	return "wallets"
}

func (wg *WalletGorm) ToAggregate() (*aggregate.Wallet, error) {
	if wg == nil {
		return nil, errors.New("o dto do agregado Wallet não pode ser vazio")
	}

	person, err := wg.Person.ToEntity()
	if err != nil {
		return nil, err
	}

	account, err := wg.Account.ToEntity()
	if err != nil {
		return nil, err
	}

	wallet, err := aggregate.NewWallet(person.Name, person.PersonalIdentification, person.Email, person.Password, person.IsAShopkeeper)
	if err != nil {
		return nil, err
	}

	entries := make([]*entity.Entry, len(wg.Entries))
	for _, etg := range wg.Entries {
		et, err := etg.ToEntity()
		if err != nil {
			return nil, err
		}
		entries = append(entries, et)
	}

	transfers := make([]valueobject.Transfer, len(wg.Transfers))
	for _, trg := range wg.Transfers {
		tr, err := trg.ToValueObject()
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, *tr)
	}

	wallet.SetAccount(account)
	wallet.SetEntries(entries...)
	wallet.SetTransfers(transfers...)

	return &wallet, nil
}

func NewRow(wallet aggregate.Wallet) WalletGorm {
	var person *gormentity.PersonGorm

	wp := wallet.GetPerson()
	if !wp.IsEmpty() {
		person = &gormentity.PersonGorm{
			UUID:                   wp.UUID.String(),
			Name:                   wp.Name,
			PersonalIdentification: wp.PersonalIdentification,
			Email:                  wp.Email,
			Password:               wp.Password,
			IsAShopkeeper:          wp.IsAShopkeeper,
		}
	}

	var account *gormentity.AccountGorm

	wa := wallet.GetAccount()
	if !wa.IsEmpty() {
		account = &gormentity.AccountGorm{
			UUID:      wa.UUID.String(),
			Person:    person,
			Balance:   wa.Balance,
			CreatedAt: wa.CreatedAt,
		}
	}

	entries := make([]gormentity.EntryGorm, len(wallet.GetEntries()))
	for _, et := range wallet.GetEntries() {
		var etg gormentity.EntryGorm
		etg.FromEntity(*et)
		entries = append(entries, etg)
	}

	transfers := make([]gormvalueobject.TransferGorm, len(wallet.GetTransfers()))
	for _, tr := range wallet.GetTransfers() {
		var trg gormvalueobject.TransferGorm
		trg.FromValueObject(tr)
		transfers = append(transfers, trg)
	}

	return WalletGorm{
		Person:    person,
		Account:   account,
		Entries:   entries,
		Transfers: transfers,
	}
}
