package main

import (
	"github.com/gin-contrib/cors"
	"paymenthub/handlers"
)

func main() {
	r := handlers.Routers()
	r.Use(cors.Default())
	r.Run(":8001")
}
