package mongodb

import (
	"context"
	"os"

	"github.com/renanfvcunha/huncoding-go-first-crud/src/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL      = "MONGODB_URL"
	MONGODB_DATABASE = "MONGODB_DATABASE"
)

func NewMongoDbConnection(ctx context.Context) (*mongo.Database, error) {
	mongodbUri := os.Getenv(MONGODB_URL)
	mongodbDatabase := os.Getenv(MONGODB_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Database Connected Successfully")

	return client.Database(mongodbDatabase), nil
}
