package valueobject

import (
	"time"

	valueobject "github.com/LucasMateus-eng/simple-bank/application/value_object"
	"github.com/google/uuid"
)

type TransferAPI struct {
	FromWalletUUID string    `json:"from_wallet_uuid"`
	ToWalletUUID   string    `json:"to_wallet_uuid"`
	Amount         float64   `json:"amount"`
	CreatedAt      time.Time `json:"created_at"`
}

func (ta *TransferAPI) ToValueObject() (*valueobject.Transfer, error) {
	if ta == nil {
		ta = &TransferAPI{}
	}

	parseFromWallet, err := uuid.Parse(ta.FromWalletUUID)
	if err != nil {
		return nil, err
	}

	parseToWallet, err := uuid.Parse(ta.ToWalletUUID)
	if err != nil {
		return nil, err
	}

	return &valueobject.Transfer{
		FromWalletUUID: parseFromWallet,
		ToWalletUUID:   parseToWallet,
		Amount:         ta.Amount,
		CreatedAt:      ta.CreatedAt,
	}, nil
}

func (ta *TransferAPI) FromValueObject(transfer valueobject.Transfer) {
	if ta == nil {
		ta = &TransferAPI{}
	}

	ta.FromWalletUUID = transfer.FromWalletUUID.String()
	ta.ToWalletUUID = transfer.ToWalletUUID.String()
	ta.Amount = transfer.Amount
	ta.CreatedAt = transfer.CreatedAt
}
