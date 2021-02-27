package document

type Permission struct {
	RootId int `bson:rootId`
	AccessList struct {
		Id int `bson:id`
		Permission byte `bson:permission`
	}`bson:accessList`
}