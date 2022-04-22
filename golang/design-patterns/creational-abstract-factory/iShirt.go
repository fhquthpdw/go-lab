package main

type iShirt interface {
	setShirtLogo(logo string)
	getShirtLogo() string
	setShirtSize(size int)
	getShirtSize() int
}

type shirt struct {
	logo string
	size int
}

func (c *shirt) setShirtLogo(logo string) {
	c.logo = logo
}

func (c *shirt) getShirtLogo() string {
	return c.logo
}

func (c *shirt) setShirtSize(size int) {
	c.size = size
}

func (c *shirt) getShirtSize() int {
	return c.size
}
