package repository

import (
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

type userRepository struct {
	dbConn *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		db,
	}
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *restErrors.RestErr)
	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *restErrors.RestErr
	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *restErrors.RestErr)
	FindUserByID(
		id string,
	) (model.UserDomainInterface, *restErrors.RestErr)
	FindAllUsers() ([]model.UserDomainInterface, *restErrors.RestErr)
	DeleteUser(
		userId string,
	) *restErrors.RestErr
}
