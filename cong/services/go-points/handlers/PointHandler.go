package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"reward-point/dto"

	"reward-point/models"
	"reward-point/services/go-points/services"
	"time"
)

var(
	pointService services.PointService
)
type PointsHandler interface {
	Inquiry(w http.ResponseWriter, r *http.Request)
	Earn(w http.ResponseWriter, r *http.Request)
	Burn(w http.ResponseWriter, r *http.Request)

	GetCustomer(w http.ResponseWriter, r *http.Request)
	GetAllCustomers(w http.ResponseWriter, r *http.Request)
	InsertCustomer(w http.ResponseWriter, r *http.Request)
	UpdateCustomer(w http.ResponseWriter, r *http.Request)
	DeleteCustomer(w http.ResponseWriter, r *http.Request)
}

type pointHandler struct {}

func NewPointHandler(service services.PointService) PointsHandler {
	pointService = service
	return &pointHandler{}
}




func (*pointHandler) Inquiry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inquiryRequestDto *dto.InquiryRequestDto
	err := json.NewDecoder(r.Body).Decode(&inquiryRequestDto)
	// Checking decode error
	if err  != nil {
		response := models.TransactionResponse{
			ErrorCode:    "02",
			ErrorMessage: "Internal Err",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	// Get customer points
	customerPoint, err := pointService.GetPoint(inquiryRequestDto.CisId)
	//Checking service error
	if err != nil {
		response := models.TransactionResponse{
			ErrorCode:    "02",
			ErrorMessage: "Internal Err",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	//Return to Client
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customerPoint)
}



func (*pointHandler) Earn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transactionRequestDto *dto.TransactionRequestDto
	err := json.NewDecoder(r.Body).Decode(&transactionRequestDto)
	//Checking decode error
	if err  != nil{
		response := models.TransactionResponse{
			ErrorCode:    "02",
			ErrorMessage: "Internal Err",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	//Get customer info
	customer, err := pointService.GetCustomer(transactionRequestDto.CisId)
	//Checking decode error
	if err != nil{
		response := models.TransactionResponse{
			ErrorCode:    "01",
			ErrorMessage: "Database Err",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	//Create request body
	timeNow := time.Now()
	earnTransaction := models.Transaction{
		TransactionType: 0,
		CreatedAt:       timeNow,
		UpdatedAt:       timeNow,
		Point:           transactionRequestDto.Point,
	}
	reqBody, _ := json.Marshal(earnTransaction)
	//Sending request to transaction service
	transactionResponse,err := http.Post(
		"http://localhost:8090/rp/api/v1/createtransaction",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	//Checking request error
	if err != nil{
		response := models.TransactionResponse{
			ErrorCode:    "02",
			ErrorMessage: "Internal Err",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	//Get data from request body
	transactionId,_ := ioutil.ReadAll(transactionResponse.Body)
	var transactionResponseDto = dto.TransactionResponseDto{}
	json.NewDecoder(bytes.NewReader([]byte(transactionId))).Decode(&transactionResponseDto)
	//Create point
	var point = models.Point{
		Amount:        transactionRequestDto.Point,
		ExpiredAt:     earnTransaction.CreatedAt.AddDate(0,2,0),
		CreatedAt:     earnTransaction.CreatedAt,
		UpdatedAt:     earnTransaction.CreatedAt,
		TransactionId: transactionResponseDto.TransactionId,
	}
	customer.AvailablePoints = append(customer.AvailablePoints, point)
	pointService.UpdateCustomer(&customer)

	response := models.TransactionResponse{
		ErrorCode:    "00",
		ErrorMessage: "Success",
	}
	json.NewEncoder(w).Encode(response)



}

func (*pointHandler) Burn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transactionRequestDto *dto.TransactionRequestDto
	err := json.NewDecoder(r.Body).Decode(&transactionRequestDto)
	//Checking decode error
	if err  != nil{
		response := models.TransactionResponse{
			ErrorCode:    "02",
			ErrorMessage: "Internal Err",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	//Get customer info
	customer, err := pointService.GetCustomer(transactionRequestDto.CisId)
	if err != nil{
		response := models.TransactionResponse{
			ErrorCode:    "01",
			ErrorMessage: "Database Err",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	//Create request body
	timeNow := time.Now()
	burnTransaction := models.Transaction{
		TransactionType: 1,
		CreatedAt:       timeNow,
		UpdatedAt:       timeNow,
		Point:           transactionRequestDto.Point,
	}
	reqBody, _ := json.Marshal(burnTransaction)

	//Sending request to transaction service
	transactionResponse,err := http.Post(
		"http://localhost:8090/rp/api/v1/createtransaction",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	//Checking request error
	if err != nil{
		response := models.TransactionResponse{
			ErrorCode:    "02",
			ErrorMessage: "Internal Err",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	//Get data from request body
	transactionId,_ := ioutil.ReadAll(transactionResponse.Body)
	var model = models.Transaction{}
	json.NewDecoder(bytes.NewReader([]byte(transactionId))).Decode(&model)

	//Update customer point
	pointService.UpdatePoint(&customer,transactionRequestDto.Point,1)
	response := models.TransactionResponse{
		ErrorCode:    "00",
		ErrorMessage: "Success",
	}
	json.NewEncoder(w).Encode(response)
}


func (*pointHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var idParam = mux.Vars(r)["id"]
	customer,err := pointService.GetCustomer(idParam)
	if err != nil{
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(customer)
}


func (*pointHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers  []models.Customer
	defer r.Body.Close()
	customers,err := pointService.GetAllCustomers()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	json.NewEncoder(w).Encode(customers)
}

func (*pointHandler) InsertCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCustomer *models.Customer
	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = pointService.InsertCustomer(newCustomer)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newCustomer)
}
func (*pointHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCustomer models.Customer
	json.NewDecoder(r.Body).Decode(&newCustomer)

	err :=  pointService.UpdateCustomer(&newCustomer)
	if err != nil{
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(http.StatusOK)

}
func (*pointHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	var idParam  = mux.Vars(r)["id"]
	err :=  pointService.DeleteCustomer(idParam)
	if err != nil{
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)

}







