package service

import (
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"go.uber.org/zap"
)

func (uds *userDomainService) UpdateUserService(userId string, ud model.UserDomainInterface) *restErrors.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	err := uds.userRepository.UpdateUser(userId, ud)

	if err != nil {
		return err
	}

	return nil
}
