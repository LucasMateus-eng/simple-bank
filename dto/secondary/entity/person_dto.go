package entity

import (
	"errors"

	"github.com/LucasMateus-eng/simple-bank/application/entity"
	"github.com/google/uuid"
)

type PersonGorm struct {
	ID                     uint   `gorm:"primaryKey;column:id"`
	UUID                   string `gorm:"uniqueIndex;column:uuid"`
	Name                   string `gorm:"not null;column:name"`
	PersonalIdentification string `gorm:"uniqueIndex;column:personal_id"`
	Email                  string `gorm:"uniqueIndex;column:email"`
	Password               string `gorm:"not null;column:password"`
	IsAShopkeeper          bool   `gorm:"not null;column:is_a_shopkeeper"`
}

func (pg *PersonGorm) TableName() string {
	return "persons"
}

func (pg *PersonGorm) ToEntity() (*entity.Person, error) {
	if pg == nil {
		return nil, errors.New("o dto da entidade Person não pode ser vazio")
	}

	parsed, err := uuid.Parse(pg.UUID)
	if err != nil {
		return nil, err
	}

	return &entity.Person{
		UUID:                   parsed,
		Name:                   pg.Name,
		PersonalIdentification: pg.PersonalIdentification,
		Email:                  pg.Email,
		Password:               pg.Password,
		IsAShopkeeper:          pg.IsAShopkeeper,
	}, nil
}

func (pg *PersonGorm) FromEntity(person entity.Person) {
	if pg == nil {
		pg = &PersonGorm{}
	}

	pg.UUID = person.UUID.String()
	pg.Name = person.Name
	pg.PersonalIdentification = person.PersonalIdentification
	pg.Email = person.Email
	pg.Password = person.Password
	pg.IsAShopkeeper = person.IsAShopkeeper
}
