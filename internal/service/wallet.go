package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"infotecs/internal/handler"
	"infotecs/internal/model"
	"infotecs/internal/repo"
	"time"
)

const (
	salt              = "rtgyumtnt"
	ResponseCreateOk  = "Кошелек создан"
	responseSendOk    = "Перевод успешно проведен"
	responseHistoryOk = "История транзакций получена"
	responseInfoOk    = "OK"
)

type WalletService struct {
	repo repo.WalletRepo
}

func NewWalletService(repo repo.WalletRepo) *WalletService {
	return &WalletService{
		repo: repo,
	}
}

func (s *WalletService) CreateWallet() (*model.ResponseWallet, error) {
	response := &model.ResponseWallet{Info: model.InfoResponse{handler.ErrInRequest}}
	id := generateClientID()

	wallet, err := s.repo.Create(id)
	if err != nil {
		return response, err
	}

	response.Info.Message = ResponseCreateOk
	response.Wallet = wallet

	return response, nil
}

func (s *WalletService) ChangeBalance(parameters *model.ParametersTransaction) (model.InfoResponse, error) {
	response := model.InfoResponse{Message: responseSendOk}

	if parameters.FromId == parameters.ToId {
		response.Message = handler.ErrBodyTx
		return response, errors.New("Id's are the same")
	}

	err := s.CheckValidId(parameters.FromId)
	if err != nil {
		response.Message = handler.ErrSenderWallet
		return response, err
	}

	err = s.CheckValidId(parameters.ToId)
	if err != nil {
		response.Message = handler.ErrBodyTx
		return response, err
	}

	if parameters.Amount <= 0.0 {
		response.Message = handler.ErrBodyTx
		return response, errors.New("Incorrect amount")
	}

	senderInfo, err := s.repo.GetWalletInfo(parameters.FromId)
	if err != nil {
		response.Message = handler.ErrServer
		return response, err
	}

	if senderInfo.Balance-parameters.Amount < 0 {
		response.Message = handler.ErrBodyTx
		return response, errors.New("Not enough money")
	}

	err = s.repo.ChangeBalance(parameters)
	if err != nil {
		response.Message = handler.ErrServer
		return response, err
	}

	return response, nil
}

func (s *WalletService) GetWalletInfo(id string) (*model.ResponseWallet, error) {
	response := &model.ResponseWallet{Info: model.InfoResponse{handler.ErrWalletNotFound}}
	err := s.CheckValidId(id)
	if err != nil {
		return response, err
	}

	wallet, err := s.repo.GetWalletInfo(id)
	if err != nil {
		return response, err
	}

	response.Info.Message = responseInfoOk
	response.Wallet = wallet

	return response, nil
}

func (s *WalletService) GetHistoryWallet(id string) (*model.ResponseGetHistory, error) {
	response := &model.ResponseGetHistory{Info: model.InfoResponse{handler.ErrWalletNotFound}}
	err := s.CheckValidId(id)
	if err != nil {
		return response, err
	}

	historyWallet, err := s.repo.GetHistoryWallet(id)
	if err != nil {
		return response, err
	}

	response.Info.Message = responseHistoryOk
	response.History = historyWallet

	return response, nil
}

func (s *WalletService) CheckValidId(id string) error {
	err := s.repo.CheckValidId(id)
	if err != nil {
		return err
	}

	return nil
}

func generateClientID() string {
	seed := time.Now().UnixNano()

	seedBytes := []byte(fmt.Sprintf("%d", seed))
	saltBytes := []byte(salt)
	data := append(seedBytes, saltBytes...)

	hasher := sha1.New()
	hasher.Write(data)
	hashBytes := hasher.Sum(nil)

	clientID := hex.EncodeToString(hashBytes)

	return clientID
}
