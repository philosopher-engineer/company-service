package main

import (
	"company-service/api/v1"
	"company-service/config"
	"company-service/internal/company"
	"company-service/internal/db"
	"company-service/internal/kafka"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Load config
	config.LoadConfig()

	// Initialize DB
	db := db.InitDB()
	defer db.Close()

	// Initialize Kafka Producer
	producer := kafka.InitKafkaProducer([]string{config.Conf.KafkaBroker})

	// Initialize repository, service, and handler
	repo := company.NewRepository(db)
	service := company.NewService(repo, producer)
	handler := company.NewHandler(service)

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	v1.AddRoutes(router, handler)

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("runtime error: %v", err)
	}
}
