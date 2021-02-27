package document

type User struct {
	Id int `bson:id`
	Email string `bson:email`
	passwd string `bson:passwd`
	IsVailEmail bool `bson:isVailEmail`
}