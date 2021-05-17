package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Config *options.ClientOptions

func init() {
	Config = options.Client()

	Config.SetAuth(options.Credential{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})

	hosts := []string{
		os.Getenv("MONGO_HOST"),
	}

	Config.SetHosts(hosts)
}

func GetClient(ctx context.Context) (*mongo.Client, error) {
	return mongo.Connect(
		ctx,
		Config,
	)
}
