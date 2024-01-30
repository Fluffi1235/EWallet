package wallet_methods

import (
	"fmt"
	"github.com/go-chi/chi"
	"infotecs/internal/handler"
	"infotecs/internal/logger"
	"infotecs/internal/service"
	"net/http"
)

const (
	errCreate = "Error create wallet_methods: "
)

func CreateWallet(r chi.Router, service *service.WalletService) {
	r.Post("/wallet", func(w http.ResponseWriter, r *http.Request) {
		response, err := service.CreateWallet()
		if err != nil {
			handler.HandlerErrors(w, errCreate, err, response, http.StatusBadRequest)
			return
		}

		handler.HandlerErrors(w, errCreate, err, response, http.StatusOK)
		logger.Logger.Info(fmt.Sprintf("Create new wallet_methods id: %s", response.Wallet.Id))
	})
}
