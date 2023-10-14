package model

import (
	"fmt"

	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *restErrors.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
