package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/controller/model/response"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info(
		"Init findUserByID controller",
		zap.String("journey", "findUserByID"),
	)

	uuid, err := primitive.ObjectIDFromHex(c.Param("userId"))

	if err != nil {
		logger.Error(
			"Error trying to validate userId",
			err,
			zap.String("journey", "findUserById"),
		)
		errorMessage := restErrors.NewBadRequestError("UserID is not a valid uuid")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userId := uuid.Hex()

	userDomain, restErr := uc.service.FindUserByIDService(userId)

	if restErr != nil {
		logger.Error(
			"Error trying to call findUserById Service",
			restErr,
			zap.String("journey", "findUserById"),
		)
		c.JSON(restErr.Code, restErr)

		return
	}

	logger.Info(
		"findUserById controller executed successfully",
		zap.String("journey", "findUserByID"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {

	email, err := mail.ParseAddress(c.Param("userEmail"))

	if err != nil {
		logger.Error(
			"Error trying to validate userEmail",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		errorMessage := restErrors.NewBadRequestError(
			"UserEmail is not a valid email",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userEmail := email.Address

	userDomain, restErr := uc.service.FindUserByEmailService(userEmail)

	if restErr != nil {
		logger.Error(
			"Error trying to call findUserByEmail Service",
			restErr,
			zap.String("journey", "findUserByEmail"),
		)
		c.JSON(restErr.Code, restErr)

		return
	}

	logger.Info(
		"findUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindAllUsers(c *gin.Context) {
	logger.Info(
		"Init findAllUsers controller",
		zap.String("journey", "findAllUsers"),
	)

	userDomain, restErr := uc.service.FindAllUsers()

	if restErr != nil {
		logger.Error(
			"Error trying to call findUserByEmail Service",
			restErr,
			zap.String("journey", "findUserByEmail"),
		)
		c.JSON(restErr.Code, restErr)

		return
	}

	var users []response.UserResponse

	for _, user := range userDomain {
		users = append(users, view.ConvertDomainToResponse(user))
	}

	c.JSON(http.StatusOK, users)
}
