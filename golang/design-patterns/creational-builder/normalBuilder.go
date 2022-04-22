package main

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (c *normalBuilder) setWindowType() {
	c.windowType = "Normal Wooden Window"
}
func (c *normalBuilder) setDoorType() {
	c.doorType = "Normal Wooden Door"
}
func (c *normalBuilder) setNumFloor() {
	c.floor = 2
}
func (c *normalBuilder) getHouse() house {
	return house{
		doorType:   c.doorType,
		windowType: c.windowType,
		floor:      c.floor,
	}
}
