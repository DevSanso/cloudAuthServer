package init

import (
	"cloudAuthServer/db/document"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)	

func InitDb(client *mongo.Client,dbName string) {
	var db = client.Database(dbName)
	createUser(db)
	createPermission(db)
	createMapping(db)
	
}

func createUser(db *mongo.Database) {
	db.CreateCollection(context.Background(),document.UserName)
	_,err := db.Collection(document.UserName).Indexes().
		CreateOne(context.Background(),mongo.IndexModel{
			Keys: bsonx.Doc {{document.UserCollectionEmailFieldName,bsonx.String("")}},
			Options : options.Index().SetUnique(true),
		});

	if err != nil {log.Panicln(err)}
	
}

func createMapping(db *mongo.Database) {
	db.CreateCollection(context.Background(),document.MappingName)
}

func createPermission(db *mongo.Database) {
	db.CreateCollection(context.Background(),document.PermissionName)
}