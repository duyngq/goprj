package models

type TransactionResponse struct {
	ErrorCode string 		`json:"error_code"`
	ErrorMessage string		`json:"error_message"`
}
