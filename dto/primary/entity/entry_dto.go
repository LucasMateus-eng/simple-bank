package entity

import (
	"time"

	"github.com/LucasMateus-eng/simple-bank/application/entity"
	"github.com/google/uuid"
)

type EntryAPI struct {
	UUID        string    `json:"uuid"`
	AccountUUID string    `json:"account_uuid"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}

func (ea *EntryAPI) ToEntity() (*entity.Entry, error) {
	if ea == nil {
		ea = &EntryAPI{}
	}

	parse, err := uuid.Parse(ea.UUID)
	if err != nil {
		return nil, err
	}

	parseAccount, err := uuid.Parse(ea.AccountUUID)
	if err != nil {
		return nil, err
	}

	return &entity.Entry{
		UUID:        parse,
		AccountUUID: parseAccount,
		Amount:      ea.Amount,
		CreatedAt:   ea.CreatedAt,
	}, nil
}

func (ea *EntryAPI) FromEntity(entry entity.Entry) {
	if ea == nil {
		ea = &EntryAPI{}
	}

	ea.UUID = entry.UUID.String()
	ea.AccountUUID = entry.AccountUUID.String()
	ea.Amount = entry.Amount
	ea.CreatedAt = entry.CreatedAt
}
