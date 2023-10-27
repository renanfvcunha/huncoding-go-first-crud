package service

import (
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"go.uber.org/zap"
)

func (uds *userDomainService) DeleteUserService(
	userId string,
) *restErrors.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))

	err := uds.userRepository.DeleteUser(userId)

	if err != nil {
		logger.Error(
			"Error trying to call repository",
			err,
			zap.String("journey", "deleteUser"),
		)

		return err
	}

	return nil
}
