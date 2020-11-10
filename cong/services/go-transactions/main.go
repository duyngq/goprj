package main

import (
	router "reward-point/http"
	"reward-point/services/go-transactions/handlers"
	"reward-point/services/go-transactions/repositories"
	"reward-point/services/go-transactions/services"
)

var (
	transactionRepository = repositories.NewTransactionRepository()
	transactionService    = services.NewTransactionService(transactionRepository)
	transactionHandler    = handlers.NewTransactionHandler(transactionService)
	httpRouter            = router.NewMuxRouter()
)
func main() {

	// /rp/api/v1/
	httpRouter.POST("/createtransaction", transactionHandler.CreateTransaction )
	httpRouter.SERVE(":8090")
}