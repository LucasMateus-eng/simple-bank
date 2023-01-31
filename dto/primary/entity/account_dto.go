package entity

import (
	"time"

	"github.com/LucasMateus-eng/simple-bank/application/entity"
	"github.com/google/uuid"
)

type AccountAPI struct {
	UUID      string    `json:"uuid"`
	Owner     string    `json:"owner"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func (ap *AccountAPI) ToEntity() (*entity.Account, error) {
	if ap == nil {
		ap = &AccountAPI{}
	}

	parse, err := uuid.Parse(ap.UUID)
	if err != nil {
		return nil, err
	}

	parseOwner, err := uuid.Parse(ap.Owner)
	if err != nil {
		return nil, err
	}

	return &entity.Account{
		UUID:      parse,
		Owner:     parseOwner,
		Balance:   ap.Balance,
		CreatedAt: ap.CreatedAt,
	}, nil
}

func (ap *AccountAPI) FromEntity(account entity.Account) {
	if ap == nil {
		ap = &AccountAPI{}
	}

	ap.UUID = account.UUID.String()
	ap.Owner = account.Owner.String()
	ap.Balance = account.Balance
	ap.CreatedAt = account.CreatedAt
}
