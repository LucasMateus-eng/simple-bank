package entity

import (
	"errors"
	"time"

	"github.com/LucasMateus-eng/simple-bank/application/entity"
	"github.com/google/uuid"
)

type AccountGorm struct {
	ID        uint        `gorm:"primaryKey;column:id"`
	UUID      string      `gorm:"uniqueIndex;column:uuid"`
	Owner     string      `gorm:"not null;column:owner"`
	Person    *PersonGorm `gorm:"foreignKey:Owner;references:UUID;constraint:OnDelete:CASCADE"`
	Balance   float64     `gorm:"not null;column:balance"`
	CreatedAt time.Time   `gorm:"not null;column:created_at"`
}

func (ag *AccountGorm) TableName() string {
	return "accounts"
}

func (ag *AccountGorm) ToEntity() (*entity.Account, error) {
	if ag == nil {
		return nil, errors.New("o dto da entidade Account não pode ser vazio")
	}

	parsed, err := uuid.Parse(ag.UUID)
	if err != nil {
		return nil, err
	}

	parsedOwner, err := uuid.Parse(ag.Owner)
	if err != nil {
		return nil, err
	}

	return &entity.Account{
		UUID:      parsed,
		Owner:     parsedOwner,
		Balance:   ag.Balance,
		CreatedAt: ag.CreatedAt,
	}, nil
}

func (ag *AccountGorm) FromEntity(account entity.Account) {
	if ag == nil {
		ag = &AccountGorm{}
	}

	ag.UUID = account.UUID.String()
	ag.Owner = account.Owner.String()
	ag.Balance = account.Balance
	ag.CreatedAt = account.CreatedAt
}
