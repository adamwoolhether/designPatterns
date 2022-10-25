package main

import (
	"fmt"
	"log"
)

func main() {
	ak47, err := getGun("ak47")
	if err != nil {
		log.Print(err)
	}

	musket, err := getGun("musket")
	if err != nil {
		log.Print(err)
	}

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}
