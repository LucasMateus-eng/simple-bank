package entity

import (
	"reflect"

	"github.com/google/uuid"
)

type Person struct {
	UUID                   uuid.UUID
	Name                   string
	PersonalIdentification string
	Email                  string
	Password               string
	IsAShopkeeper          bool
}

func (p *Person) IsEmpty() bool {
	if p == nil {
		return true
	}

	if reflect.DeepEqual(p, &Person{}) {
		return true
	}

	return false
}
