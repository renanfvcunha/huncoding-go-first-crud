package service

import (
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"go.uber.org/zap"
)

func (uds *userDomainService) FindUserByIDService(
	id string,
) (model.UserDomainInterface, *restErrors.RestErr) {
	logger.Info("Init findUserByID services", zap.String("journey", "findUserByID"))

	return uds.userRepository.FindUserByID(id)
}

func (uds *userDomainService) FindUserByEmailService(
	email string,
) (model.UserDomainInterface, *restErrors.RestErr) {
	logger.Info("Init findUserByEmail services", zap.String("journey", "findUserByEmail"))

	return uds.userRepository.FindUserByEmail(email)
}
