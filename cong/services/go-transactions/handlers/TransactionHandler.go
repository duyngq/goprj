package handlers

import (
	"encoding/json"
	"net/http"
	"reward-point/models"
	"reward-point/services/go-transactions/services"
)

var(
	transService services.TransactionService
)
type TransactionHandler interface {
	CreateTransaction(w http.ResponseWriter, r *http.Request)
}

type handler struct {}

func NewTransactionHandler(service services.TransactionService) TransactionHandler {
	transService = service
	return &handler{}
}


func (*handler) CreateTransaction (w http.ResponseWriter, r *http.Request) {

	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err  != nil{
		response := models.TransactionResponse{
			ErrorCode:    "02",
			ErrorMessage: "Internal Err",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	id,err := transService.CreateTransaction(&transaction)
	if err != nil {
		response := models.TransactionResponse{
			ErrorCode:    "01",
			ErrorMessage: "Database Err",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"transaction_id": id})
}



