package document

import (
	"go.mongodb.org/mongo-driver/bson"
)

type ContainerMapping struct {
	UserId int `bson:userId`
	ContainerId int `bson:cotainerId`
}
const (
	CMCollectionUserIdFieldName = "userId"
	CmCollectionContainerIdFieldName = "cotainerId"
)
func(cm *ContainerMapping)ToBson() bson.D {
	return bson.D {
		{CMCollectionUserIdFieldName,cm.UserId},
		{CmCollectionContainerIdFieldName,cm.ContainerId},
	}
	
}

const MappingName = "container_mapping"