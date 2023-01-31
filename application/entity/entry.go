package entity

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	UUID        uuid.UUID
	AccountUUID uuid.UUID
	Amount      float64
	CreatedAt   time.Time
}

func (e *Entry) IsEmpty() bool {
	if e == nil {
		return true
	}

	if reflect.DeepEqual(e, &Entry{}) {
		return true
	}

	return false
}
