package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func GetContainer(ctx context.Context,userId int) (int,error) {
	var result int

	err := _client.UseSession(ctx,func(sess mongo.SessionContext) error {
		err := sess.StartTransaction(options.Transaction())
		if err != nil {return err}
		
		var doc = bson.M {"userId":userId}

		res := sess.Client().Database(_databaseName).Collection(document.MappingName).FindOne(ctx,doc)
		err = res.Err()
		if err != nil {sess.AbortTransaction(ctx);return err}
		var temp bson.M
		err = res.Decode(&temp)
		if err != nil {sess.AbortTransaction(ctx);return err}
		result = temp[document.CmCollectionContainerIdFieldName].(int)
		return sess.CommitTransaction(ctx)
	})
	return result,err
}