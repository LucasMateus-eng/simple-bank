package wallet

import (
	"errors"

	"github.com/LucasMateus-eng/simple-bank/application/aggregate"
	primaryport "github.com/LucasMateus-eng/simple-bank/application/ports/primary/wallet"
	secondaryport "github.com/LucasMateus-eng/simple-bank/application/ports/secondary/wallet"
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
