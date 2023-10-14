package service

import (
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *restErrors.RestErr
	UpdateUser(string, model.UserDomainInterface) *restErrors.RestErr
	FindUser(string) (*model.UserDomainInterface, *restErrors.RestErr)
	DeleteUser(string) *restErrors.RestErr
}
