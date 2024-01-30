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
	errGetHistory = "Error get history wallet_methods: "
)

func WalledHistory(r chi.Router, service *service.WalletService) {
	r.Get("/wallet/{walletId}/history", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "walletId")

		response, err := service.GetHistoryWallet(id)
		if err != nil {
			handler.HandlerErrors(w, errGetHistory, err, response, http.StatusNotFound)
			return
		}

		handler.HandlerErrors(w, errGetHistory, err, response, http.StatusOK)
		logger.Logger.Info(fmt.Sprintf("Get history wallet_methods: %s", id))
	})
}
