package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/validation"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/controller/model/request"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser Controller", zap.String("journey", "updateUser"))
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "updateUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := restErrors.NewBadRequestError("Invalid userId, must be a hex value")

		c.JSON(errRest.Code, errRest)
	}

	domain := model.UpdateUserDomain(userRequest.Name, userRequest.Age)

	err := uc.service.UpdateUserService(userId, domain)

	if err != nil {
		logger.Error(
			"Error trying to call UpdateUser service",
			err,
			zap.String("journey", "updateUser"),
		)
		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"User created successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)
}
