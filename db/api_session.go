package db

import (
	"context"

	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cloudAuthServer/db/document"
)

func SetContainerId(ctx context.Context,userId int,containerId int) error {
	err := _client.UseSession(ctx,func(sess mongo.SessionContext) error {
		err := sess.StartTransaction(options.Transaction())
		if err != nil {return err}
		
		var doc = document.ContainerMapping {
			userId,
			containerId,
		}

		_,err = sess.Client().Database(_databaseName).Collection(document.MappingName).
		InsertOne(ctx,doc.ToBson())
		if err != nil {sess.AbortTransaction(ctx);return err}
		return sess.CommitTransaction(ctx)
	})
	return err
}