package valueobject

import (
	"errors"
	"time"

	valueobject "github.com/LucasMateus-eng/simple-bank/application/value_object"
	gormaggregate "github.com/LucasMateus-eng/simple-bank/dto/secondary/aggregate"
	"github.com/google/uuid"
)

type TransferGorm struct {
	ID             uint                      `gorm:"primaryKey;column:id"`
	FromWalletUUID string                    `gorm:"not null;column:from_account_uuid"`
	FromWallet     *gormaggregate.WalletGorm `gorm:"foreignKey:from_account_uuid;references:UUID;constraint:OnDelete:CASCADE"`
	ToWalletUUID   string                    `gorm:"not null;column:to_account_uuid"`
	ToWallet       *gormaggregate.WalletGorm `gorm:"foreignKey:to_account_uuid;references:UUID;constraint:OnDelete:CASCADE"`
	Amount         float64                   `gorm:"not null;column:amount"`
	CreatedAt      time.Time                 `gorm:"not null;column:created_at"`
}

func (tg *TransferGorm) TableName() string {
	return "transfers"
}

func (tg *TransferGorm) ToValueObject() (*valueobject.Transfer, error) {
	if tg == nil {
		return nil, errors.New("o dto do objeto de valor Transfer não pode ser vazio")
	}

	parsedFromWallet, err := uuid.Parse(tg.FromWalletUUID)
	if err != nil {
		return nil, err
	}

	parsedToWallet, err := uuid.Parse(tg.ToWalletUUID)
	if err != nil {
		return nil, err
	}

	return &valueobject.Transfer{
		FromWalletUUID: parsedFromWallet,
		ToWalletUUID:   parsedToWallet,
		Amount:         tg.Amount,
		CreatedAt:      tg.CreatedAt,
	}, nil
}

func (tg *TransferGorm) FromValueObject(transfer valueobject.Transfer) {
	if tg == nil {
		tg = &TransferGorm{}
	}

	tg.FromWalletUUID = transfer.FromWalletUUID.String()
	tg.ToWalletUUID = transfer.ToWalletUUID.String()
	tg.Amount = transfer.Amount
	tg.CreatedAt = transfer.CreatedAt
}
