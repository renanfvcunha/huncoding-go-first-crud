package service

import (
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"go.uber.org/zap"
)

func (uds *userDomainService) LoginUserService(
	ud model.UserDomainInterface,
) (model.UserDomainInterface, *restErrors.RestErr) {
	logger.Info("Init loginUser model", zap.String("journey", "loginUser"))

	ud.EncryptPassword()

	user, err := uds.findUserByEmailAndPasswordService(
		ud.GetEmail(),
		ud.GetPassword(),
	)

	if err != nil {
		return nil, err
	}

	logger.Info(
		"loginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"),
	)

	return user, nil
}
