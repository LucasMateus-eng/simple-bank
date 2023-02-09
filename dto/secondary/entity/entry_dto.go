package entity

import (
	"errors"
	"time"

	"github.com/LucasMateus-eng/simple-bank/application/entity"
	gormaggregate "github.com/LucasMateus-eng/simple-bank/dto/secondary/aggregate"
	"github.com/google/uuid"
)

type EntryGorm struct {
	ID        uint                     `gorm:"primaryKey;column:id"`
	UUID      string                   `gorm:"uniqueIndex;column:uuid"`
	Owner     string                   `gorm:"not null;column:wallet_uuid"`
	Wallet    gormaggregate.WalletGorm `gorm:"foreignKey:wallet_uuid;references:UUID;constraint:OnDelete:CASCADE"`
	Amount    float64                  `gorm:"not null;column:amount"`
	CreatedAt time.Time                `gorm:"not null;column:created_at"`
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

	parsedOwner, err := uuid.Parse(eg.Owner)
	if err != nil {
		return nil, err
	}

	return &entity.Entry{
		UUID:      parsed,
		Owner:     parsedOwner,
		Amount:    eg.Amount,
		CreatedAt: eg.CreatedAt,
	}, nil
}

func (eg *EntryGorm) FromEntity(entry entity.Entry) {
	if eg == nil {
		eg = &EntryGorm{}
	}

	eg.UUID = entry.UUID.String()
	eg.Owner = entry.Owner.String()
	eg.Amount = entry.Amount
	eg.CreatedAt = entry.CreatedAt
}
