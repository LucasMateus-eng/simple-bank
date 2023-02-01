package wallet

import (
	"fmt"
	"net/http"

	primaryport "github.com/LucasMateus-eng/simple-bank/application/ports/primary/wallet"
	apiaggregate "github.com/LucasMateus-eng/simple-bank/dto/primary/aggregate"
	"github.com/LucasMateus-eng/simple-bank/utils/formatter"
	"github.com/LucasMateus-eng/simple-bank/utils/logging"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	log = logging.NewLogger()
)

type WalletHandler struct {
	walletService primaryport.WalletService
}

func NewWalletHandler(walletService primaryport.WalletService) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

// Get handler to get a wallet based on wallet id
//
// @Summary Get a wallet
// @Description Get a wallet based on the wallet id from the database
// @Tags wallet
// @Produce json
// @Param   wallet_id	path string	true "wallet identifier"
// @Success 200 {object} formatter.ResponseOKWithData "Successful get wallet."
// @Failure 400 {object} formatter.ResponseErrorWithData "Error parsing the id parameter."
// @Failure 404 {object} formatter.ResponseErrorWithData "Error: wallet can't find."
// @Failure 500 {object} formatter.ResponseErrorWithData "Error getting wallet."
// @Router /wallet/{wallet_id} [get]
func (wh *WalletHandler) Get(c echo.Context) error {
	walletID := c.Param("wallet_id")

	parsed, err := uuid.Parse(walletID)
	if err != nil {
		log.Error("Erro ao consultar a carteira na API rest: ", err.Error())
		return formatter.ErrorWithDataJSON(c, http.StatusBadRequest, "Erro ao realizar o parse no parâmeto uuid", err.Error())
	}

	result, err := wh.walletService.Get(parsed)
	if err != nil {
		log.Error("Erro ao consultar a carteira na API rest: ", err.Error())
		return formatter.ErrorWithDataJSON(c, http.StatusInternalServerError, "Erro ao consultar carteira", err.Error())
	}

	if result.IsEmpty() {
		err := fmt.Errorf("não foram econtrados dados para a carteira %s", walletID)
		log.Error("Erro ao consultar a carteira na API rest: ", err.Error())
		return formatter.ErrorWithDataJSON(c, http.StatusNotFound, "Erro ao consultar carteira", err.Error())
	}

	var wallet *apiaggregate.WalletAPI
	wallet.FromAggregate(result)

	return formatter.SuccessWithDataJSON(c, http.StatusOK, "Sucesso ao consultar carteira", wallet)
}

// Add handler for creating new wallet
//
// @Summary Create a wallet
// @Description Create a new wallet in the database
// @Tags wallet
// @Accept json
// @Produce json
// @Param wallet body apiaggregate.WalletAPI true "Wallet DTO for create"
// @Success 201 {object} formatter.ResponseOKWithData "Wallet successfully created."
// @Failure 400 {object} formatter.ResponseErrorWithData "Invalid payload."
// @Failure 500 {object} formatter.ResponseErrorWithData "Failed to create wallet."
// @Router /wallet [post]
func (wh *WalletHandler) Add(c echo.Context) error {
	var body = new(apiaggregate.WalletAPI)
	if err := c.Bind(&body); err != nil {
		log.Error("Erro ao criar carteira na API rest: ", err.Error())
		return formatter.ErrorWithDataJSON(c, http.StatusBadRequest, "Payload inválido", err.Error())
	}

	wallet, err := body.ToAggregate()
	if err != nil {
		log.Error("Erro ao criar carteira na API rest: ", err.Error())
		return formatter.ErrorWithDataJSON(c, http.StatusInternalServerError, "Erro ao converter para agregado", err.Error())
	}

	err = wh.walletService.Add(*wallet)
	if err != nil {
		log.Error("Erro ao criar carteira na API rest: ", err.Error())
		return formatter.ErrorWithDataJSON(c, http.StatusInternalServerError, "Erro ao criar carteira", err.Error())
	}

	return formatter.SuccessJSON(c, http.StatusCreated, "Carteira criada com sucesso")
}

// Update handler for updating a wallet
//
// @Summary Update a wallet
// @Description Update a wallet in the database
// @Tags wallet
// @Accept json
// @Produce json
// @Param wallet body apiaggregate.WalletForUpdateAPI true "Wallet DTO for update"
// @Success 204 {object} formatter.ResponseOKWithData "Wallet successfully updated."
// @Failure 400 {object} formatter.ResponseErrorWithData "Invalid payload."
// @Failure 500 {object} formatter.ResponseErrorWithData "Failed to update wallet."
// @Router /wallet [put]
func (wh *WalletHandler) Update(c echo.Context) error {
	var body = new(apiaggregate.WalletForUpdateAPI)
	if err := c.Bind(&body); err != nil {
		log.Error("Erro ao atualizar carteira na API rest: ", err.Error())
		return formatter.ErrorWithDataJSON(c, http.StatusBadRequest, "Payload inválido", err.Error())
	}

	wallet, err := body.ToAggregate()
	if err != nil {
		log.Error("Erro ao atualizar carteira na API rest: ", err.Error())
		return formatter.ErrorWithDataJSON(c, http.StatusInternalServerError, "Erro ao converter para agregado", err.Error())
	}

	err = wh.walletService.Update(*wallet)
	if err != nil {
		log.Error("Erro ao atualizar carteira na API rest: ", err.Error())
		return formatter.ErrorWithDataJSON(c, http.StatusInternalServerError, "Erro ao atualizar carteira", err.Error())
	}

	return formatter.SuccessJSON(c, http.StatusOK, "Carteira atualizada com sucesso")
}

// Delete handler for wallet deletion based on wallet id
//
// @Summary Delete a wallet
// @Description Delete a wallet based on the wallet id from the database
// @Tags wallet
// @Produce json
// @Param   wallet_id	path string	true "wallet identifier"
// @Success 204 {object} formatter.Response "Wallet successfully deleted."
// @Failure 400 {object} formatter.ResponseErrorWithData "Error parsing the id parameter."
// @Failure 500 {object} formatter.ResponseErrorWithData "Failed to delete wallet."
// @Router /wallet/{wallet_id} [delete]
func (wh *WalletHandler) Delete(c echo.Context) error {
	walletID := c.Param("wallet_id")

	parsed, err := uuid.Parse(walletID)
	if err != nil {
		log.Error("Erro ao excluir a carteira na API rest: ", err.Error())
		return formatter.ErrorWithDataJSON(c, http.StatusBadRequest, "Erro ao realizar o parse no parâmeto uuid", err.Error())
	}

	if err := wh.walletService.Delete(parsed); err != nil {
		log.Error("Erro ao excluir a carteira na API rest: ", err.Error())
		return formatter.ErrorWithDataJSON(c, http.StatusInternalServerError, "Erro ao excluir carteira", err.Error())
	}

	return formatter.SuccessJSON(c, http.StatusNoContent, "Carteira excluída com sucesso")
}
