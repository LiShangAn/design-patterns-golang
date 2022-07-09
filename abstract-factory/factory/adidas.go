package factory

type adidasShoe struct {
	shoe
}

type adidasShirt struct {
	shirt
}

type Adidas struct{}

func (a *Adidas) MakeShoe() IShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidase",
			size: 10,
		},
	}
}

func (a *Adidas) MakeShirt() IShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidase",
			size: 100,
		},
	}
}
