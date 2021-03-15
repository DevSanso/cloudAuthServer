package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cloudAuthServer/db/document"
)




func RegisterUser(ctx context.Context,user *document.User) error {
	return _client.UseSession(ctx,func(session mongo.SessionContext) error {
		info := user.ToBsonAndHashingPasswd()

		err := session.StartTransaction(options.Transaction())
		if err != nil {return err}
		
		_,err = session.Client().Database(_databaseName).
			Collection(document.UserName).
			InsertOne(ctx,info,&options.InsertOneOptions{})
		
		if err != nil {
			session.AbortTransaction(ctx)
			return err
		}	
		err = session.CommitTransaction(ctx)
		
		return err
	})

}
func GetUserId(ctx context.Context,user *document.User) (int,error) {
	session,err := _client.StartSession(options.Session())
	if err != nil {return 0,err}
	var res bson.M
	err = session.Client().Database(_databaseName).
	Collection(document.MappingName).
	FindOne(ctx,user.ToBsonSubVailEmailAndHashing()).
	Decode(&res)
	if err != nil {return 0,err}
	return res["_id"].(int),nil
}

func ExistUser(ctx context.Context,user *document.User) (bool,error) {
	var res bson.M
	var err = _client.UseSession(ctx,func(session mongo.SessionContext) error {
		info := user.ToBsonSubVailEmailAndHashing()

		err := session.StartTransaction(options.Transaction())
		if err != nil {return err}

		
		err = session.Client().Database(_databaseName).
		Collection(document.UserName).
		FindOne(ctx,info).
		Decode(&res)
		if err != nil {session.AbortTransaction(ctx);return err}
		return session.CommitTransaction(ctx)
	})

	
	if res != nil {
		return true,err
	}
	return false,err
}

