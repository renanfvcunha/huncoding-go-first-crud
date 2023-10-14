package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
)

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

type UserDomainInterface interface {
	CreateUser() *restErrors.RestErr
	UpdateUser(string) *restErrors.RestErr
	FindUser(string) (*UserDomain, *restErrors.RestErr)
	DeleteUser(string) *restErrors.RestErr
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
