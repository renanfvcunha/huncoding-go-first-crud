package service

import (
	"fmt"

	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"go.uber.org/zap"
)

func (uds *userDomainService) CreateUser(ud model.UserDomainInterface) *restErrors.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
