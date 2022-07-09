package factory

type nike struct {
}

type nikeShoe struct {
	shoe
}

func (a *nike) makeShoe() {

}

func (a *nike) makeShirt() {

}
