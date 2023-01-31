package wallet

import (
	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
	"github.com/google/uuid"
)

type WalletService interface {
	Get(uuid.UUID) (aggregate.Wallet, error)
	Add(aggregate.Wallet) error
	Update(aggregate.Wallet) error
	Delete(uuid.UUID) error
}
