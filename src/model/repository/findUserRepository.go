package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/restErrors"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model/repository/entity"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *restErrors.RestErr) {
	logger.Info(
		"Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"),
	)

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.dbConn.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))

			return nil, restErrors.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))

		return nil, restErrors.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserByEmail repository executed successfully.",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *restErrors.RestErr) {
	logger.Info(
		"Init findUserById repository",
		zap.String("journey", "findUserById"),
	)

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.dbConn.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "findUserById"))

			return nil, restErrors.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage, err, zap.String("journey", "findUserById"))

		return nil, restErrors.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserById repository executed successfully.",
		zap.String("journey", "findUserById"),
		zap.String("userId", userEntity.ID.Hex()),
	)

	return converter.ConvertEntityToDomain(*userEntity), nil
}
