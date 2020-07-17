package main

import (
	"fmt"
	"sirclo-test/ship-yard/models"
)

type ShipYard struct {
	ships []models.Ship
}

func (s *ShipYard) AddShip(ship models.Ship) {
	s.ships = append(s.ships, ship)
}

func (s *ShipYard) GetAllShip() []models.Ship {
	return s.ships
}

type ShipYardPrinter struct{}

func (p ShipYardPrinter) PrintAllLocation(ships []models.Ship,
	fn func(ship models.Ship) string) {

	for _, ship := range ships {
		fmt.Println(fn(ship))
	}
}

func main() {
	motorBoat := models.NewMotorBoat("MB001", 5, 2)
	sailBoat := models.NewSailBoat("SB001", 6, 8)
	cruiseShip := models.NewCruiseShip("CS001", 1, 5)

	shipYard := ShipYard{}
	shipYard.AddShip(motorBoat)
	shipYard.AddShip(sailBoat)
	shipYard.AddShip(cruiseShip)

	shipYardPrinter := ShipYardPrinter{}
	shipYardPrinter.PrintAllLocation(shipYard.GetAllShip(), func(ship models.Ship) string {
		x, y := ship.GetLocation()
		return fmt.Sprintf("Code: %s, X: %d, Y: %d", ship.GetCode(), x, y)
	})
}
