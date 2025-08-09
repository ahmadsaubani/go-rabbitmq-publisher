package main

import (
	"log"
	"net/http"
	"publisher-topic/src/routes"
	"publisher-topic/src/utils/rabbitmqs"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	err := rabbitmqs.InitRabbitMQ()
	if err != nil {
		log.Fatalf("RabbitMQ init error: %v", err)
	}

	router := routes.API(gin.Default())

	s := &http.Server{
		Addr:           ":9001",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
