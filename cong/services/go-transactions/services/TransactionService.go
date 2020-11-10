package services

import (
	"reward-point/models"
	"reward-point/services/go-transactions/repositories"
	"time"
)

type TransactionService interface {
	GetTransaction(id string)  (models.Transaction,error)
	CreateTransaction(trans *models.Transaction) (string,error)
}
type service struct {}
var(
	repo repositories.TransactionRepository
)

func NewTransactionService(repository repositories.TransactionRepository) TransactionService {
	repo = repository
	return &service{}
}

func (*service) GetTransaction(id string)  (models.Transaction,error){
	transaction,err := repo.GetTransaction(id)
	return transaction,err
}
func (*service) CreateTransaction(trans *models.Transaction) (string, error){
	timeNow := time.Now()

	trans.CreatedAt = timeNow
	trans.UpdatedAt = timeNow
	id,err := repo.InsertTransaction(trans)
	if err != nil {
		return id,err
	}

	 return id,err
}