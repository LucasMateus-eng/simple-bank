package entity

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Owner     uuid.UUID
	Balance   float64
	CreatedAt time.Time
}

func NewAccount(owner uuid.UUID, balance float64) *Account {
	return &Account{
		Owner:     owner,
		Balance:   balance,
		CreatedAt: time.Now(),
	}
}

func (a *Account) IsEmpty() bool {
	if a == nil {
		return true
	}

	if reflect.DeepEqual(a, &Account{}) {
		return true
	}

	return false
}
