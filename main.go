package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/controller"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/controller/routes"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model/service"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init Deps
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
