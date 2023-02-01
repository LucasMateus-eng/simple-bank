package wallet

import (
	"errors"
	"fmt"

	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
	primaryport "github.com/LucasMateus-eng/simple-bank/application/ports/primary/wallet"
	secondaryport "github.com/LucasMateus-eng/simple-bank/application/ports/secondary/wallet"
	valueobject "github.com/LucasMateus-eng/simple-bank/application/value_object"
	"github.com/LucasMateus-eng/simple-bank/utils/logging"
	"github.com/google/uuid"
)

var (
	log = logging.NewLogger()
)

type walletService struct {
	walletRepo secondaryport.WalletRepository
}

func NewWalletService(walletRepo secondaryport.WalletRepository) primaryport.WalletService {
	return &walletService{
		walletRepo: walletRepo,
	}
}

func (ws *walletService) Get(id uuid.UUID) (aggregate.Wallet, error) {
	wallet, err := ws.walletRepo.Get(id)
	if err != nil {
		log.Errorf("Erro ao consultar carteira %s no walletService: %s", id, err.Error())
		return aggregate.Wallet{}, err
	}

	if wallet.IsEmpty() {
		err := errors.New("não foi possível consultar a carteria no walletService")
		log.Errorf("Erro ao consultar carteira %s no walletService: %s", id, err.Error())
		return aggregate.Wallet{}, err
	}

	return wallet, nil
}

func (ws *walletService) Add(wallet aggregate.Wallet) error {
	err := ws.walletRepo.Add(wallet)
	if err != nil {
		log.Error("Erro ao salvar uma carteira no walletService: ", err.Error())
		return err
	}

	return nil
}

func (ws *walletService) Update(wallet aggregate.Wallet) error {
	err := ws.walletRepo.Update(wallet)
	if err != nil {
		log.Error("Erro ao atualizar uma carteira no walletService: ", err.Error())
		return err
	}

	return nil
}

func (ws *walletService) Delete(id uuid.UUID) error {
	err := ws.walletRepo.Delete(id)
	if err != nil {
		log.Error("Erro ao deletar uma carteira no walletService: ", err.Error())
		return err
	}

	return nil
}

func (ws *walletService) Transfer(transfer valueobject.Transfer) error {
	if transfer.Amount <= 0 {
		err := errors.New("o valor transferido deve ser positivo e diferente de zero")
		log.Error("Erro ao realizar uma transferência no walletService: ", err.Error())
		return err
	}

	walletThatTransfers, err := ws.Get(transfer.FromWalletUUID)
	if err != nil {
		err := fmt.Errorf("não foi possível consultar os dados da carteira %s que está transferindo. Detalhes: %s", transfer.FromWalletUUID, err.Error())
		log.Error("Erro ao realizar uma transferência no walletService: ", err.Error())
		return err
	}

	walletHolder := walletThatTransfers.GetPerson()
	if walletHolder.IsAShopkeeper {
		err := fmt.Errorf("o titular %s da carteira que está transferindo não pode ser logista", walletHolder.UUID)
		log.Error("Erro ao realizar uma transferência no walletService: ", err.Error())
		return err
	}

	walletAccount := walletThatTransfers.GetAccount()
	if walletAccount.Balance-transfer.Amount < 0 {
		err := fmt.Errorf("a carteira %s que está transferindo deve possuir saldo suficiente. Saldo atual: %v", walletAccount.Owner, walletAccount.Balance)
		log.Error("Erro ao realizar uma transferência no walletService: ", err.Error())
		return err
	}

	err = ws.walletRepo.Transfer(transfer)
	if err != nil {
		log.Error("Erro ao realizar uma transferência no walletService: ", err.Error())
		return err
	}

	return nil
}
