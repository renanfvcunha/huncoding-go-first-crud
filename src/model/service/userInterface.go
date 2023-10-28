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
	UpdateUserService(string, model.UserDomainInterface) *restErrors.RestErr
	FindAllUsers() ([]model.UserDomainInterface, *restErrors.RestErr)
	FindUserByIDService(
		id string,
	) (model.UserDomainInterface, *restErrors.RestErr)
	FindUserByEmailService(
		email string,
	) (model.UserDomainInterface, *restErrors.RestErr)
	DeleteUserService(string) *restErrors.RestErr
	LoginUserService(
		model.UserDomainInterface,
	) (model.UserDomainInterface, *restErrors.RestErr)
}

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}
