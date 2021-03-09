package init
import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cloudAuthServer/db/document"
	
)

func InitDb(client *mongo.Client,dbName string) {
	var db = client.Database(dbName)
	createUser(db)
	createPermission(db)
	createMapping(db)
	
}

func createUser(db *mongo.Database) {
	db.CreateCollection(context.Background(),document.UserName)
}

func createMapping(db *mongo.Database) {
	db.CreateCollection(context.Background(),document.MappingName)
}

func createPermission(db *mongo.Database) {
	db.CreateCollection(context.Background(),document.PermissionName)
}