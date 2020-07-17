package models

type sailBoat struct {
	code string
	x, y int
}

func NewSailBoat(code string, x, y int) Ship {
	return &sailBoat{code, x, y}
}

func (sb *sailBoat) SetCode(code string) {
	sb.code = code
}

func (sb *sailBoat) GetCode() string {
	return sb.code
}

func (sb *sailBoat) SetLocation(x, y int) {
	sb.x = x
	sb.y = y
}

func (sb *sailBoat) GetLocation() (x, y int) {
	return sb.x, sb.y
}
