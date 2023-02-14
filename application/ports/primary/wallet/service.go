package wallet

import (
	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
	"github.com/LucasMateus-eng/simple-bank/application/entity"
	valueobject "github.com/LucasMateus-eng/simple-bank/application/value_object"
	"github.com/google/uuid"
)

type WalletService interface {
	Get(uuid.UUID) (aggregate.Wallet, error)
	Add(entity.Person) error
	Update(entity.Person) error
	Delete(uuid.UUID) error
	Transfer(valueobject.Transfer) error
	Deposit(valueobject.Transfer) error
}
