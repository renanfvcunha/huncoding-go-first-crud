package service

import (
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model/repository"
)

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *restErrors.RestErr)
	UpdateUser(string, model.UserDomainInterface) *restErrors.RestErr
	FindUserByIDService(
		id string,
	) (model.UserDomainInterface, *restErrors.RestErr)
	FindUserByEmailService(
		email string,
	) (model.UserDomainInterface, *restErrors.RestErr)
	DeleteUser(string) *restErrors.RestErr
}

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}
