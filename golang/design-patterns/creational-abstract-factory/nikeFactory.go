package main

type nike struct {
}

func (c *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (c *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 18,
		},
	}
}
