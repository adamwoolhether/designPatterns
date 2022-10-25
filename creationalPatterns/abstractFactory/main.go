package main

import (
	"fmt"
	"log"
)

func main() {
	adidasFactory, err := GetSportsFactory("adidas")
	if err != nil {
		log.Println(err)
	}
	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()
	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	nikeFactory, err := GetSportsFactory("adidas")
	if err != nil {
		log.Println(err)
	}
	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()
	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s\nSize: %d\n", s.getLogo(), s.getSize())
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s\nSize: %d\n", s.getLogo(), s.getSize())
}
