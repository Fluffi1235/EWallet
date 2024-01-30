package controller

import (
	"github.com/go-chi/chi"
	"infotecs/internal/controller/wallet_methods"
	"infotecs/internal/service"
)

func WalletController(r *chi.Mux, service *service.WalletService) {
	r.Route("/api/v1", func(r chi.Router) {
		wallet_methods.CreateWallet(r, service)
		wallet_methods.SendMoney(r, service)
		wallet_methods.GetWalletInfo(r, service)
		wallet_methods.WalledHistory(r, service)
	})
}
