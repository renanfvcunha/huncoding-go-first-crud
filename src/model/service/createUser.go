package service

import (
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"go.uber.org/zap"
)

func (uds *userDomainService) CreateUserService(
	ud model.UserDomainInterface,
) (model.UserDomainInterface, *restErrors.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	user, _ := uds.FindUserByEmailService(ud.GetEmail())

	if user != nil {
		return nil, restErrors.NewUnprocessableEntityError(
			"Email informed already exists",
		)
	}

	ud.EncryptPassword()

	userDomainRepository, err := uds.userRepository.CreateUser(ud)

	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
