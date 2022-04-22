package main

type adidas struct {
}

func (c adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 15,
		},
	}
}

func (c adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 19,
		},
	}
}
