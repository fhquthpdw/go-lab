package main

type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

const (
	AdidasFactory = "adidas"
	NikeFactory   = "nike"
)

func getSportsFactory(brand string) iSportsFactory {
	switch brand {
	case AdidasFactory:
		return &adidas{}
	case NikeFactory:
		return &nike{}
	default:
		return nil
	}
}
