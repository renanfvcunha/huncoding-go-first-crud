package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser Controller", zap.String("journey", "deleteUser"))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := restErrors.NewBadRequestError("Invalid userId, must be a hex value")

		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUserService(userId)

	if err != nil {
		logger.Error(
			"Error trying to call DeleteUser service",
			err,
			zap.String("journey", "deleteUser"),
		)
		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"User created successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	c.Status(http.StatusOK)
}
