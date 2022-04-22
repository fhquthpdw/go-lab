package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

const (
	NormalBuilderType = "normal"
	IglooBuilderType  = "igloo"
)

func getBuilder(builderType string) iBuilder {
	switch builderType {
	case NormalBuilderType:
		return &normalBuilder{}
	case IglooBuilderType:
		return &iglooBuilder{}
	default:
		return nil
	}
}
