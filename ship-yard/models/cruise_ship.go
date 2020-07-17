package models

type cruiseShip struct {
	code string
	x, y int
}

func NewCruiseShip(code string, x, y int) Ship {
	return &cruiseShip{code, x, y}
}

func (cs *cruiseShip) SetCode(code string) {
	cs.code = code
}

func (cs *cruiseShip) GetCode() string {
	return cs.code
}

func (cs *cruiseShip) SetLocation(x, y int) {
	cs.x = x
	cs.y = y
}

func (cs *cruiseShip) GetLocation() (x, y int) {
	return cs.x, cs.y
}
