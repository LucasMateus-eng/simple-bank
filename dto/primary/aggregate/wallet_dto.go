package aggregate

import (
	"errors"

	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
	"github.com/LucasMateus-eng/simple-bank/application/entity"
	valueobject "github.com/LucasMateus-eng/simple-bank/application/value_object"
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

	person, err := wa.Person.ToEntity()
	if err != nil {
		return nil, err
	}

	account, err := wa.Account.ToEntity()
	if err != nil {
		return nil, err
	}

	wallet, err := aggregate.NewWallet(person.Name, person.PersonalIdentification, person.Email, person.Password, person.IsAShopkeeper)
	if err != nil {
		return nil, err
	}

	entries := make([]*entity.Entry, len(wa.Entries))
	for _, eta := range wa.Entries {
		et, err := eta.ToEntity()
		if err != nil {
			return nil, err
		}
		entries = append(entries, et)
	}

	transfers := make([]valueobject.Transfer, len(wa.Transfers))
	for _, tra := range wa.Transfers {
		tr, err := tra.ToValueObject()
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
	for _, et := range wallet.GetEntries() {
		var eta apientity.EntryAPI
		eta.FromEntity(*et)
		entries = append(entries, eta)
	}

	transfers := make([]apivalueobject.TransferAPI, len(wallet.GetTransfers()))
	for _, tr := range wallet.GetTransfers() {
		var tra apivalueobject.TransferAPI
		tra.FromValueObject(tr)
		transfers = append(transfers, tra)
	}

	wa.Person = person
	wa.Account = account
	wa.Entries = entries
	wa.Transfers = transfers
}

type WalletForUpdateAPI struct {
	Person    apientity.PersonAPI          `json:"person"`
	Account   apientity.AccountAPI         `json:"account"`
	Entries   []apientity.EntryAPI         `json:"entries"`
	Transfers []apivalueobject.TransferAPI `json:"transfers"`
}

func (wua *WalletForUpdateAPI) ToAggregate() (*aggregate.Wallet, error) {
	if wua == nil {
		return nil, errors.New("o dto do agregado Wallet não pode ser vazio")
	}

	person, err := wua.Person.ToEntity()
	if err != nil {
		return nil, err
	}

	account, err := wua.Account.ToEntity()
	if err != nil {
		return nil, err
	}

	wallet, err := aggregate.NewWallet(person.Name, person.PersonalIdentification, person.Email, person.Password, person.IsAShopkeeper)
	if err != nil {
		return nil, err
	}

	entries := make([]*entity.Entry, len(wua.Entries))
	for _, eta := range wua.Entries {
		et, err := eta.ToEntity()
		if err != nil {
			return nil, err
		}
		entries = append(entries, et)
	}

	transfers := make([]valueobject.Transfer, len(wua.Transfers))
	for _, tra := range wua.Transfers {
		tr, err := tra.ToValueObject()
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, *tr)
	}

	wallet.SetID(person.UUID)
	wallet.SetAccount(account)
	wallet.SetEntries(entries...)
	wallet.SetTransfers(transfers...)

	return &wallet, nil
}
