package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (nb *IglooBuilder) setWindowType() {
	nb.windowType = "Snow Window"
}

func (nb *IglooBuilder) setDoorType() {
	nb.doorType = "Snow Door"
}

func (nb *IglooBuilder) setNumFloor() {
	nb.floor = 1
}

func (nb *IglooBuilder) getHouse() House {
	return House{
		windowType: nb.windowType,
		doorType:   nb.doorType,
		floor:      nb.floor,
	}
}
