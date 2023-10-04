package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
)

func CreateUser(c *gin.Context) {
	err := restErrors.NewBadRequestError("Você chamou a rota de forma errada")

	c.JSON(err.Code, err)
}
