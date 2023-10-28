package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/validation"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/controller/model/request"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser Controller", zap.String("journey", "loginUser"))

	var userLoginRequest request.UserLogin

	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "loginUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userLoginRequest.Email,
		userLoginRequest.Password,
	)

	domainResult, err := uc.service.LoginUserService(domain)

	if err != nil {
		logger.Error(
			"Error trying to call LoginUser service",
			err,
			zap.String("journey", "loginUser"),
		)
		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"User logged successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"),
	)

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
