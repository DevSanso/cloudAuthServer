package document

type ContainerMapping struct {
	userId int `bson:userId`
	ContainerId int `bson:cotainerId`
}

const MappingName = "container_mapping"