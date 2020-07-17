package models

type motorBoat struct {
	code string
	x, y int
}

func NewMotorBoat(code string, x, y int) Ship {
	return &motorBoat{code, x, y}
}

func (mb *motorBoat) SetCode(code string) {
	mb.code = code
}

func (mb *motorBoat) GetCode() string {
	return mb.code
}

func (mb *motorBoat) SetLocation(x, y int) {
	mb.x = x
	mb.y = y
}

func (mb *motorBoat) GetLocation() (x, y int) {
	return mb.x, mb.y
}
