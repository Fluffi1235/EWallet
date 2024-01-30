package handler

import (
	"encoding/json"
	"github.com/pkg/errors"
	"infotecs/internal/logger"
	"net/http"
)

const (
	ErrInRequest      = "Ошибка в запросе"
	ErrWalletNotFound = "Указанный кошелек не найден"
	ErrSenderWallet   = "Исходящий кошелек не найден"
	ErrBodyTx         = "Ошибка в пользовательском запросе или ошибка перевода"
	ErrServer         = "Ошибка сервера"
)

func HandlerErrors(w http.ResponseWriter, handle string, err error, resp interface{}, statusCode int) {
	if err != nil {
		newErr := errors.Wrap(err, handle)
		logger.Logger.Error(newErr)
	}

	SendResponses(w, statusCode, resp)
}

func SendResponses(w http.ResponseWriter, statusCode int, resp interface{}) {
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
	}

	if resp == nil {
		return
	}

	respJSN, err := json.Marshal(&resp)
	if err != nil {
		logger.Logger.Error("Error marshal: ", err.Error())
		return
	}

	_, err = w.Write(respJSN)
	if err != nil {
		logger.Logger.Error("Error writes the data to the connection as part of an HTTP reply: ", err.Error())
	}
}
