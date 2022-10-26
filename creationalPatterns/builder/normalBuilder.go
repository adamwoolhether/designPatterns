package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (nb *NormalBuilder) setWindowType() {
	nb.windowType = "Wooden Window"
}

func (nb *NormalBuilder) setDoorType() {
	nb.doorType = "Wooden Door"
}

func (nb *NormalBuilder) setNumFloor() {
	nb.floor = 2
}

func (nb *NormalBuilder) getHouse() House {
	return House{
		windowType: nb.windowType,
		doorType:   nb.doorType,
		floor:      nb.floor,
	}
}
