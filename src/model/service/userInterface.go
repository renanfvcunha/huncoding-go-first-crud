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
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *restErrors.RestErr)
	UpdateUser(string, model.UserDomainInterface) *restErrors.RestErr
	FindUser(string) (*model.UserDomainInterface, *restErrors.RestErr)
	DeleteUser(string) *restErrors.RestErr
}

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}
