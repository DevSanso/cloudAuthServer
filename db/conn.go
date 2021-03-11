package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _client *mongo.Client = nil
var _databaseName = ""

func Connect(ctx context.Context,dbName string,opt ...*options.ClientOptions) error {
	if _client != nil {
		return errors.New("already connect database")
	}
	var err error
	_client,err = mongo.Connect(ctx,opt...)
	_databaseName = dbName
	return err
}

