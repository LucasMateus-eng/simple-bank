package wallet

import (
	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
	valueobject "github.com/LucasMateus-eng/simple-bank/application/value_object"
	"github.com/google/uuid"
)

type WalletRepository interface {
	Get(uuid.UUID) (aggregate.Wallet, error)
	Add(aggregate.Wallet) error
	Update(aggregate.Wallet) error
	Delete(uuid.UUID) error
	Transfer(valueobject.Transfer) error
}
