package main

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (c *director) setBuilder(b iBuilder) {
	c.builder = b
}

func (c *director) buildHouse() house {
	c.builder.setWindowType()
	c.builder.setDoorType()
	c.builder.setNumFloor()
	return c.builder.getHouse()
}
