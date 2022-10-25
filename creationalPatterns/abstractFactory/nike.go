package main

type Nike struct {
}

func (a *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (a *Nike) makeShirt() IShirt {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 10,
		},
	}
}

// Concreate Product
type NikeShoe struct {
	Shoe
}

// Concreate Product
type NikeShirt struct {
	Shirt
}
