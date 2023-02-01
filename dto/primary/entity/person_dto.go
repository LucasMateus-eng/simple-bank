package entity

import (
	"errors"

	"github.com/LucasMateus-eng/simple-bank/application/entity"
	"github.com/google/uuid"
)

type PersonAPI struct {
	UUID                   string `json:"uuid"`
	Name                   string `json:"name"`
	PersonalIdentification string `json:"personal_id"`
	Email                  string `json:"email"`
	Password               string `json:"password"`
	IsAShopkeeper          bool   `json:"is_a_shopkeeper"`
}

func (pa *PersonAPI) ToEntity() (*entity.Person, error) {
	if pa == nil {
		return nil, errors.New("o dto da entidade Person não pode ser vazio")
	}

	parse, err := uuid.Parse(pa.UUID)
	if err != nil {
		return nil, err
	}

	return &entity.Person{
		UUID:                   parse,
		Name:                   pa.Name,
		PersonalIdentification: pa.PersonalIdentification,
		Email:                  pa.Email,
		Password:               pa.Password,
		IsAShopkeeper:          pa.IsAShopkeeper,
	}, nil
}

func (pa *PersonAPI) FromEntity(person entity.Person) {
	if pa == nil {
		pa = &PersonAPI{}
	}

	pa.UUID = person.UUID.String()
	pa.Name = person.Name
	pa.PersonalIdentification = person.PersonalIdentification
	pa.Email = person.Email
	pa.Password = person.Password
	pa.IsAShopkeeper = person.IsAShopkeeper
}
