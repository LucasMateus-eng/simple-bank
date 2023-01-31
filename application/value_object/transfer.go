package valueobject

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	FromWalletUUID uuid.UUID
	ToWalletUUID   uuid.UUID
	Amount         float64
	CreatedAt      time.Time
}

func (t *Transfer) IsEmpty() bool {
	if t == nil {
		return true
	}

	if reflect.DeepEqual(t, &Transfer{}) {
		return true
	}

	return false
}
