package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/validation"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/controller/model/request"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	domainResult, err := uc.service.CreateUser(domain)

	if err != nil {
		c.JSON(err.Code, err)

		return
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.JSON(201, view.ConvertDomainToResponse(domainResult))
}
