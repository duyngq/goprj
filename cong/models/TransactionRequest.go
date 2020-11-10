package models

type TransactionRequest struct {
	CisId string `json:"cis_id" `
	RequestId string `json:"request_id" `
	Point int `json:"point" `
}