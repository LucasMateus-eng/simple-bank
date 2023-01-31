package entity

import (
	"errors"
	"time"

	"github.com/LucasMateus-eng/simple-bank/application/entity"
	"github.com/google/uuid"
)

type EntryGorm struct {
	ID          uint         `gorm:"primaryKey;column:id"`
	UUID        string       `gorm:"uniqueIndex;column:uuid"`
	AccountUUID string       `gorm:"not null;column:account_uuid"`
	Account     *AccountGorm `gorm:"foreignKey:account_uuid;references:UUID;constraint:OnDelete:CASCADE"`
	Amount      float64      `gorm:"not null;column:amount"`
	CreatedAt   time.Time    `gorm:"not null;column:created_at"`
}

func (eg *EntryGorm) TableName() string {
	return "entries"
}

func (eg *EntryGorm) ToEntity() (*entity.Entry, error) {
	if eg == nil {
		return nil, errors.New("o dto da entidade Entry não pode ser vazio")
	}

	parsed, err := uuid.Parse(eg.UUID)
	if err != nil {
		return nil, err
	}

	parsedAccount, err := uuid.Parse(eg.AccountUUID)
	if err != nil {
		return nil, err
	}

	return &entity.Entry{
		UUID:        parsed,
		AccountUUID: parsedAccount,
		Amount:      eg.Amount,
		CreatedAt:   eg.CreatedAt,
	}, nil
}

func (eg *EntryGorm) FromEntity(entry entity.Entry) {
	if eg == nil {
		eg = &EntryGorm{}
	}

	eg.UUID = entry.UUID.String()
	eg.AccountUUID = entry.AccountUUID.String()
	eg.Amount = entry.Amount
	eg.CreatedAt = entry.CreatedAt
}
