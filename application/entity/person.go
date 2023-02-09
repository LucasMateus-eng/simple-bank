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

func NewPerson(name, personalID, email, password string, isAShopkeeper bool) *Person {
	return &Person{
		UUID:                   uuid.New(),
		Name:                   name,
		PersonalIdentification: personalID,
		Email:                  email,
		Password:               password,
		IsAShopkeeper:          isAShopkeeper,
	}
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
