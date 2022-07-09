package factory

import (
	"fmt"
)

type iSportFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetSportsFactory(name string) (iSportFactory, error) {
	if name == "adidas" {
		return &Adidas{}, nil
	}
	// if brand == "nike" {
	// 	return &nike{}, nil
	// }

	return nil, fmt.Errorf("wrong type")
}
