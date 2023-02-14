package aggregate

import (
	"errors"

	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
	apientity "github.com/LucasMateus-eng/simple-bank/dto/primary/entity"
	apivalueobject "github.com/LucasMateus-eng/simple-bank/dto/primary/value_object"
)

type WalletAPI struct {
	Person    apientity.PersonAPI          `json:"person"`
	Account   apientity.AccountAPI         `json:"account"`
	Entries   []apientity.EntryAPI         `json:"entries"`
	Transfers []apivalueobject.TransferAPI `json:"transfers"`
}

func (wa *WalletAPI) ToAggregate() (*aggregate.Wallet, error) {
	if wa == nil {
		return nil, errors.New("o dto do agregado Wallet não pode ser vazio")
	}

	wallet, err := aggregate.NewWallet(wa.Person.Name, wa.Person.PersonalIdentification, wa.Person.Email, wa.Person.Password, wa.Person.IsAShopkeeper)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (wa *WalletAPI) FromAggregate(wallet aggregate.Wallet) {
	if wa == nil {
		wa = &WalletAPI{}
	}

	var person apientity.PersonAPI

	wp := wallet.GetPerson()
	if !wp.IsEmpty() {
		person = apientity.PersonAPI{
			UUID:                   wp.UUID.String(),
			Name:                   wp.Name,
			PersonalIdentification: wp.PersonalIdentification,
			Email:                  wp.Email,
			Password:               wp.Password,
			IsAShopkeeper:          wp.IsAShopkeeper,
		}
	}

	var account apientity.AccountAPI

	wac := wallet.GetAccount()
	if !wac.IsEmpty() {
		account = apientity.AccountAPI{
			Owner:     wac.Owner.String(),
			Balance:   wac.Balance,
			CreatedAt: wac.CreatedAt,
		}
	}

	entries := make([]apientity.EntryAPI, len(wallet.GetEntries()))
	for i, et := range wallet.GetEntries() {
		var eta apientity.EntryAPI
		eta.FromEntity(*et)
		entries[i] = eta
	}

	transfers := make([]apivalueobject.TransferAPI, len(wallet.GetTransfers()))
	for i, tr := range wallet.GetTransfers() {
		var tra apivalueobject.TransferAPI
		tra.FromValueObject(tr)
		transfers[i] = tra
	}

	wa.Person = person
	wa.Account = account
	wa.Entries = entries
	wa.Transfers = transfers
}
