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

func (a *Account) IsEmpty() bool {
	if a == nil {
		return true
	}

	if reflect.DeepEqual(a, &Account{}) {
		return true
	}

	return false
}
