package models

type Ship interface {
	SetCode(code string)
	GetCode() string
	SetLocation(x, y int)
	GetLocation() (x, y int)
}
