package valueobject

import (
	"errors"
	"time"

	valueobject "github.com/LucasMateus-eng/simple-bank/application/value_object"
	gormentity "github.com/LucasMateus-eng/simple-bank/dto/secondary/entity"
	"github.com/google/uuid"
)

type TransferGorm struct {
	ID              uint                    `gorm:"primaryKey;column:id"`
	FromAccountUUID string                  `gorm:"not null;column:from_account_uuid"`
	FromAccount     *gormentity.AccountGorm `gorm:"foreignKey:from_account_uuid;references:UUID;constraint:OnDelete:CASCADE"`
	ToAccountUUID   string                  `gorm:"not null;column:to_account_uuid"`
	ToAccount       *gormentity.AccountGorm `gorm:"foreignKey:to_account_uuid;references:UUID;constraint:OnDelete:CASCADE"`
	Amount          float64                 `gorm:"not null;column:amount"`
	CreatedAt       time.Time               `gorm:"not null;column:created_at"`
}

func (tg *TransferGorm) TableName() string {
	return "transfers"
}

func (tg *TransferGorm) ToValueObject() (*valueobject.Transfer, error) {
	if tg == nil {
		return nil, errors.New("o dto do objeto de valor Transfer não pode ser vazio")
	}

	parsedFromAccount, err := uuid.Parse(tg.FromAccountUUID)
	if err != nil {
		return nil, err
	}

	parsedToAccount, err := uuid.Parse(tg.ToAccountUUID)
	if err != nil {
		return nil, err
	}

	return &valueobject.Transfer{
		FromAccountUUID: parsedFromAccount,
		ToAccountUUID:   parsedToAccount,
		Amount:          tg.Amount,
		CreatedAt:       tg.CreatedAt,
	}, nil
}

func (tg *TransferGorm) FromValueObject(transfer valueobject.Transfer) {
	if tg == nil {
		tg = &TransferGorm{}
	}

	tg.FromAccountUUID = transfer.FromAccountUUID.String()
	tg.ToAccountUUID = transfer.ToAccountUUID.String()
	tg.Amount = transfer.Amount
	tg.CreatedAt = transfer.CreatedAt
}
