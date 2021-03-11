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
		
		session.EndSession(ctx)
		return err
	})

	
	if res != nil {
		return true,err
	}
	return false,err
}
