package main

import (
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"infotecs/internal/config"
	"infotecs/internal/controller"
	"infotecs/internal/logger"
	"infotecs/internal/repo"
	"infotecs/internal/service"
	"log"
	"net/http"
)

func main() {
	err := logger.NewLogger()
	if err != nil {
		log.Fatal("Can't initialize logger: ", err.Error())
	}

	cfg, err := config.LoadConfigFromYaml()
	if err != nil {
		logger.Logger.Fatal("Error read config: ", err.Error())
	}

	db, err := repo.NewPostgresDB(cfg)
	if err != nil {
		logger.Logger.Fatal("Error connect DB: ", err.Error())
	}

	repository := repo.NewRepo(db)
	walletService := service.NewWalletService(repository)
	defer db.Close()

	router := chi.NewRouter()
	controller.WalletController(router, walletService)

	logger.Logger.Info("Start server")

	err = http.ListenAndServe(cfg.ServicePort, router)
	if err != nil {
		log.Fatal(err)
	}
}
