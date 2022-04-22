package main

type iShoe interface {
	setShoeLogo(logo string)
	getShoeLogo() string
	setShoeSize(size int)
	getShoeSize() int
}

type shoe struct {
	logo string
	size int
}

func (c *shoe) setShoeLogo(logo string) {
	c.logo = logo
}

func (c *shoe) getShoeLogo() string {
	return c.logo
}

func (c *shoe) setShoeSize(size int) {
	c.size = size
}

func (c *shoe) getShoeSize() int {
	return c.size
}
