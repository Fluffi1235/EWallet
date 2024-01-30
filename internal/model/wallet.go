package model

type Wallet struct {
	Id      string `json:"id" db:"id"`
	Balance int    `json:"balance" db:"balance"`
}

type ParametersTransaction struct {
	Time   string `json:"time"`
	FromId string `json:"from"`
	ToId   string `json:"to"`
	Amount int    `json:"amount"`
}

type InfoResponse struct {
	Message string `json:"status"`
}

type ResponseWallet struct {
	Info   InfoResponse
	Wallet *Wallet `json:"wallet"`
}

type ResponseGetHistory struct {
	Info    InfoResponse
	History []*ParametersTransaction `json:"history"`
}
