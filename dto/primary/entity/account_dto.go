package entity

import (
	"errors"
	"time"

	"github.com/LucasMateus-eng/simple-bank/application/entity"
	"github.com/google/uuid"
)

type AccountAPI struct {
	Owner     string    `json:"owner"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func (ap *AccountAPI) ToEntity() (*entity.Account, error) {
	if ap == nil {
		return nil, errors.New("o dto da antidade Account não pode ser vazio")
	}

	parseOwner, err := uuid.Parse(ap.Owner)
	if err != nil {
		return nil, err
	}

	return &entity.Account{
		Owner:     parseOwner,
		Balance:   ap.Balance,
		CreatedAt: ap.CreatedAt,
	}, nil
}

func (ap *AccountAPI) FromEntity(account entity.Account) {
	if ap == nil {
		ap = &AccountAPI{}
	}

	ap.Owner = account.Owner.String()
	ap.Balance = account.Balance
	ap.CreatedAt = account.CreatedAt
}
