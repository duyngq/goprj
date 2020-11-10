package models

type BodyRequest struct {
	CustomerId string `json:"customer_id"`
	PaymentId  string `json:"payment_id"`
	MerchantId string `json:"merchant_id"`
	Amount float64 `json:"amount"`
	Description string `json:"description"`
}