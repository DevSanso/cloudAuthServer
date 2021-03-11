package document

import (
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"

	"go.mongodb.org/mongo-driver/bson"

)

type User struct {
	Email string `bson:email`
	Passwd string `bson:passwd`
	IsVailEmail bool `bson:isVailEmail`
}

const (
	UserCollectionEmailFieldName = "email"
	UserCollectionPasswdFieldName  = "passwd"
	UserCollectionIsVaildFieldName = "isVaild"
)


func (u *User)ToBson() bson.D {
	return bson.D{
	{UserCollectionEmailFieldName,u.Email},
	{UserCollectionPasswdFieldName,u.Passwd},
	{UserCollectionIsVaildFieldName,false}}
}

func (u *User)ToBsonAndHashingPasswd() bson.D {
	return bson.D{
	{UserCollectionEmailFieldName,u.Email},
	{UserCollectionPasswdFieldName,u.passwordHashing()},
	{UserCollectionIsVaildFieldName,false}}
}

func(u *User)ToBsonSubVailEmailAndHashing() bson.D {
	return bson.D{
		{UserCollectionEmailFieldName,u.Email},
		{UserCollectionPasswdFieldName,u.passwordHashing()},
	}
}
func(u *User)passwordHashing() string {
	return string(pbkdf2.Key([]byte(u.Passwd),[]byte(u.Email),1000,64,sha256.New))
}


const UserName = "user"