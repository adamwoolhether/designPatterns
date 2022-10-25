package main

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 10,
		},
	}
}

// Concreate Product
type AdidasShoe struct {
	Shoe
}

// Concreate Product
type AdidasShirt struct {
	Shirt
}
