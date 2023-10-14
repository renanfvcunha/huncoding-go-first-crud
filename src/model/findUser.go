package model

import "github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"

func (ud *UserDomain) FindUser(string) (*UserDomain, *restErrors.RestErr) {
	return &UserDomain{
		Email:    ud.Email,
		Name:     ud.Name,
		Password: ud.Password,
		Age:      ud.Age,
	}, nil
}
