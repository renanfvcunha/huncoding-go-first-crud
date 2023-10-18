package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/database/mongodb"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/controller"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/controller/routes"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model/repository"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model/service"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDbConnection(context.Background())

	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
	}

	// Init Deps
	repository := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repository)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
