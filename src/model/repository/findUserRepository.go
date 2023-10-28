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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
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

func (ur *userRepository) FindAllUsers() ([]model.UserDomainInterface, *restErrors.RestErr) {
	logger.Info(
		"Init findAllUsers repository",
		zap.String("journey", "findAllUsers"),
	)

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.dbConn.Collection(collectionName)

	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		errorMessage := "Error on finding users"
		logger.Error(
			errorMessage,
			err,
			zap.String("journey", "findAllUsers"),
		)

		return nil, restErrors.NewInternalServerError(errorMessage)
	}

	var queryResults []entity.UserEntity

	if err = cur.All(context.Background(), &queryResults); err != nil {
		errorMessage := "Error on decoding users"
		logger.Error(
			errorMessage,
			err,
			zap.String("journey", "findAllUsers"),
		)

		return nil, restErrors.NewInternalServerError(errorMessage)
	}

	var results []model.UserDomainInterface

	for _, result := range queryResults {
		results = append(results, converter.ConvertEntityToDomain(result))
	}

	return results, nil
}

func (ur *userRepository) FindUserByEmailAndPassword(
	email string,
	password string,
) (model.UserDomainInterface, *restErrors.RestErr) {
	logger.Info(
		"Init findUserByEmailAndPassword repository",
		zap.String("journey", "findUserByEmailAndPassword"),
	)

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.dbConn.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "Invalid credentials"
			logger.Error(
				errorMessage,
				err,
				zap.String("journey", "findUserByEmailAndPassword"),
			)

			return nil, restErrors.NewUnauthorizedError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(
			errorMessage,
			err,
			zap.String("journey", "findUserByEmailAndPassword"),
		)

		return nil, restErrors.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserByEmail repository executed successfully.",
		zap.String("journey", "findUserByEmailAndPassword"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)

	return converter.ConvertEntityToDomain(*userEntity), nil
}
