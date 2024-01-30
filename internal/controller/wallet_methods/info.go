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
	errGetInfo = "Error get info wallet_methods: "
)

func GetWalletInfo(r chi.Router, service *service.WalletService) {
	r.Get("/wallet/{walletId}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "walletId")

		response, err := service.GetWalletInfo(id)
		if err != nil {
			handler.HandlerErrors(w, errGetInfo, err, response, http.StatusNotFound)
			return
		}

		handler.HandlerErrors(w, errGetInfo, err, response, http.StatusOK)
		logger.Logger.Info(fmt.Sprintf("Return info wallet_methods with id: %s", id))
	})
}
