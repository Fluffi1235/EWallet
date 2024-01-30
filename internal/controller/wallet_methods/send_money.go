package wallet_methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"infotecs/internal/handler"
	"infotecs/internal/logger"
	"infotecs/internal/model"
	"infotecs/internal/service"
	"io"
	"net/http"
)

const (
	errSendMoney = "Error send money: "
	badRequest   = "Ошибка в пользовательском запросе"
)

func SendMoney(r chi.Router, service *service.WalletService) {
	r.Post("/wallet/{walletId}/send", func(w http.ResponseWriter, r *http.Request) {
		fromId := chi.URLParam(r, "walletId")
		response := model.InfoResponse{badRequest}
		parameters, err := io.ReadAll(r.Body)
		if err != nil {
			handler.HandlerErrors(w, errSendMoney, err, response, http.StatusBadRequest)
			return
		}

		decoder := json.NewDecoder(bytes.NewReader(parameters))
		parametersTransaction := &model.ParametersTransaction{FromId: fromId}
		err = decoder.Decode(parametersTransaction)
		if err != nil {
			handler.HandlerErrors(w, errSendMoney, err, response, http.StatusBadRequest)
			return
		}

		response, err = service.ChangeBalance(parametersTransaction)
		if err != nil {
			if response.Message == handler.ErrSenderWallet {
				handler.HandlerErrors(w, errSendMoney, err, response, http.StatusNotFound)
			} else {
				handler.HandlerErrors(w, errSendMoney, err, response, http.StatusBadRequest)
			}
			return
		}

		handler.HandlerErrors(w, errSendMoney, err, response, http.StatusOK)
		logger.Logger.Info(fmt.Sprintf("Transaction completed fromId: %s, toId: %s, amount: %d",
			parametersTransaction.FromId, parametersTransaction.ToId, parametersTransaction.Amount))
	})
}
