package view

import (
	"github.com/renanfvcunha/huncoding-go-first-crud/src/controller/model/response"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "0",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
