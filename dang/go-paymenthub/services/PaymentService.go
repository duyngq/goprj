package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"paymenthub/models"
	"paymenthub/repositories"
)
type PaymentService struct{}
type PaymentInterface interface {
	GetInquiry()
	ChargePayment()
	CancelPayment()
	SetSettlement()
}
//Get detail information of payment by payment_id, customer_id, merchant_id
func (paymentService PaymentService) GetInquiry(c *gin.Context) {
	var repo repositories.PaymentRepository
	var body models.BodyRequest
	err := c.BindJSON(&body)
	if err != nil {
		log.Fatal(err)
	}
	payment := repo.GetInquiry(body.PaymentId, body.CustomerId, body.MerchantId)
	c.JSON(http.StatusOK, payment)
}
//Pay for the product by merchant_id
func (paymentService PaymentService) ChargePayment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func (paymentService PaymentService) CancelPayment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func (paymentService PaymentService) SetSettlement(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

