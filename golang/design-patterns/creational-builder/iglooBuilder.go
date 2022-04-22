package main

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (c *iglooBuilder) setWindowType() {
	c.windowType = "Igloo Snow Window"
}
func (c *iglooBuilder) setDoorType() {
	c.doorType = "Igloo Snow Door"
}
func (c *iglooBuilder) setNumFloor() {
	c.floor = 1
}
func (c *iglooBuilder) getHouse() house {
	return house{
		doorType:   c.doorType,
		windowType: c.windowType,
		floor:      c.floor,
	}
}
