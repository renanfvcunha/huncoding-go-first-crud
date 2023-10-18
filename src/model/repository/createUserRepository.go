package repository

import (
	"context"
	"os"

	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *restErrors.RestErr) {
	logger.Info("Init createUser repository")
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.dbConn.Collection(collectionName)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, restErrors.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*value), nil
}
