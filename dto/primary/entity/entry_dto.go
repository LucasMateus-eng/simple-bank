package entity

import (
	"errors"
	"time"

	"github.com/LucasMateus-eng/simple-bank/application/entity"
	"github.com/google/uuid"
)

type EntryAPI struct {
	UUID      string    `json:"uuid"`
	Owner     string    `json:"account_uuid"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func (ea *EntryAPI) ToEntity() (*entity.Entry, error) {
	if ea == nil {
		return nil, errors.New("o dto da entidade Entry não pode ser vazio")
	}

	parse, err := uuid.Parse(ea.UUID)
	if err != nil {
		return nil, err
	}

	parseOwner, err := uuid.Parse(ea.Owner)
	if err != nil {
		return nil, err
	}

	return &entity.Entry{
		UUID:      parse,
		Owner:     parseOwner,
		Amount:    ea.Amount,
		CreatedAt: ea.CreatedAt,
	}, nil
}

func (ea *EntryAPI) FromEntity(entry entity.Entry) {
	if ea == nil {
		ea = &EntryAPI{}
	}

	ea.UUID = entry.UUID.String()
	ea.Owner = entry.Owner.String()
	ea.Amount = entry.Amount
	ea.CreatedAt = entry.CreatedAt
}
