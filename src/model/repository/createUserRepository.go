package repository

import (
	"context"
	"os"

	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *restErrors.RestErr) {
	logger.Info("Init createUser repository")
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.dbConn.Collection(collectionName)

	jsonValue, err := userDomain.GetJSONValue()

	if err != nil {
		return nil, restErrors.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), jsonValue)

	if err != nil {
		return nil, restErrors.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
