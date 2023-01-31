package valueobject

import (
	"time"

	valueobject "github.com/LucasMateus-eng/simple-bank/application/value_object"
	"github.com/google/uuid"
)

type TransferAPI struct {
	FromAccountUUID string    `json:"from_account_uuid"`
	ToAccountUUID   string    `json:"to_account_uuid"`
	Amount          float64   `json:"amount"`
	CreatedAt       time.Time `json:"created_at"`
}

func (ta *TransferAPI) ToValueObject() (*valueobject.Transfer, error) {
	if ta == nil {
		ta = &TransferAPI{}
	}

	parseFromAccount, err := uuid.Parse(ta.FromAccountUUID)
	if err != nil {
		return nil, err
	}

	parseToAccount, err := uuid.Parse(ta.ToAccountUUID)
	if err != nil {
		return nil, err
	}

	return &valueobject.Transfer{
		FromAccountUUID: parseFromAccount,
		ToAccountUUID:   parseToAccount,
		Amount:          ta.Amount,
		CreatedAt:       ta.CreatedAt,
	}, nil
}

func (ta *TransferAPI) FromValueObject(transfer valueobject.Transfer) {
	if ta == nil {
		ta = &TransferAPI{}
	}

	ta.FromAccountUUID = transfer.FromAccountUUID.String()
	ta.ToAccountUUID = transfer.ToAccountUUID.String()
	ta.Amount = transfer.Amount
	ta.CreatedAt = transfer.CreatedAt
}
