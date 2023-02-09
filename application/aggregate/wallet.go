package aggregate

import (
	"errors"
	"reflect"
	"time"

	"github.com/LucasMateus-eng/simple-bank/application/entity"
	valueobject "github.com/LucasMateus-eng/simple-bank/application/value_object"
	"github.com/LucasMateus-eng/simple-bank/utils/validations"
	"github.com/google/uuid"
)

type Wallet struct {
	person    *entity.Person
	account   *entity.Account
	entries   []*entity.Entry
	transfers []valueobject.Transfer
}

func (w *Wallet) IsEmpty() bool {
	if w == nil {
		return true
	}

	if reflect.DeepEqual(w, &Wallet{}) {
		return true
	}

	return false
}

func NewWallet(name, personalID, email, password string, isAShopKeeper bool) (Wallet, error) {
	if validations.StringIsEmptyOrWhiteSpace(name) {
		return Wallet{}, errors.New("o nome do usuário não pode ser vazio")
	}

	if validations.StringIsEmptyOrWhiteSpace(personalID) {
		return Wallet{}, errors.New("a identificação pessoal do usuário não pode ser vazia")
	}

	if validations.StringIsEmptyOrWhiteSpace(email) {
		return Wallet{}, errors.New("o email do usuário não pode ser vazio")
	}

	if validations.StringIsEmptyOrWhiteSpace(password) {
		return Wallet{}, errors.New("a senha do usuário não pode ser vazia")
	}

	personUUID := uuid.New()
	person := &entity.Person{
		UUID:                   personUUID,
		Name:                   name,
		PersonalIdentification: personalID,
		Email:                  email,
		Password:               password,
		IsAShopkeeper:          isAShopKeeper,
	}

	account := &entity.Account{
		Owner:     personUUID,
		Balance:   0,
		CreatedAt: time.Now(),
	}

	return Wallet{
		person:    person,
		account:   account,
		entries:   make([]*entity.Entry, 0),
		transfers: make([]valueobject.Transfer, 0),
	}, nil
}

func (w *Wallet) GetID() uuid.UUID {
	return w.person.UUID
}

func (w *Wallet) SetID(id uuid.UUID) {
	if w.person.IsEmpty() {
		w.person = &entity.Person{}
	}
	w.person.UUID = id
}

func (w *Wallet) GetPerson() entity.Person {
	return *w.person
}

func (w *Wallet) SetPerson(person *entity.Person) {
	if w.account.IsEmpty() || person.IsEmpty() {
		w.person = &entity.Person{}
	}
	w.person = person
}

func (w *Wallet) GetAccount() *entity.Account {
	return w.account
}

func (w *Wallet) SetAccount(acc *entity.Account) {
	if w.account.IsEmpty() || acc.IsEmpty() {
		w.account = &entity.Account{}
	}
	w.account = acc
}

func (w *Wallet) GetEntries() []*entity.Entry {
	return w.entries
}

func (w *Wallet) SetEntries(entries ...*entity.Entry) {
	if w.entries == nil {
		w.entries = make([]*entity.Entry, len(entries))
	}
	for _, entryInput := range entries {
		for i := range w.entries {
			w.entries[i] = entryInput
		}
	}
}

func (w *Wallet) GetTransfers() []valueobject.Transfer {
	return w.transfers
}

func (w *Wallet) SetTransfers(transfers ...valueobject.Transfer) {
	if w.transfers == nil {
		w.transfers = make([]valueobject.Transfer, len(transfers))
	}
	for _, transferInput := range transfers {
		for i := range w.entries {
			w.transfers[i] = transferInput
		}
	}
}
