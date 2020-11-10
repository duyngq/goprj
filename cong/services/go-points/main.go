package main

import (
	router "reward-point/http"
	"reward-point/services/go-points/handlers"
	"reward-point/services/go-points/repositories"
	"reward-point/services/go-points/services"
)



var (
		pointRepository = repositories.NewPointRepository()
		pointServ  = services.NewPointService(pointRepository)
		pointHandler = handlers.NewPointHandler(pointServ)
		httpRouter  = router.NewMuxRouter()
)
func main() {
		// TestAPI
		httpRouter.GET("/customers", pointHandler.GetAllCustomers )
		httpRouter.GET("/customers/{id}", pointHandler.GetCustomer )
		httpRouter.POST("/customers", pointHandler.InsertCustomer )
		httpRouter.PUT("/customers/{id}", pointHandler.UpdateCustomer )
		httpRouter.DELETE("/customers/{id}", pointHandler.DeleteCustomer )

		// /rp/api/v1/
		httpRouter.POST("/inquiry", pointHandler.Inquiry )
		httpRouter.POST("/burn", pointHandler.Burn )
		httpRouter.POST("/earn", pointHandler.Earn )

		httpRouter.SERVE(":8080")
}

