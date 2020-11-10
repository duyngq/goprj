package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"paymenthub/models"
	"paymenthub/repositories"
	"time"
)

type AccountService struct{}

type AccountServiceInterface interface {
	StoreAccount()
}

func (accountService AccountService) StoreAccount(c *gin.Context) {
	var repo repositories.AccountRepository
	var body models.Account
	err := c.BindJSON(&body)
	if err != nil {
		log.Fatal(err)
	}
	body.CreatedAt = time.Now()
	body.UpdatedAt = time.Now()
	body.Birthday = time.Now()
	status, id, errorrRepo := repo.StoreAccount(body)
	if errorrRepo != nil {
		log.Fatal(errorrRepo)
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "id": id})
}
