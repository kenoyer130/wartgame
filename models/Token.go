package models

type Token struct {
	RGBA RGBA
	ID   string
	Base int
}

type RGBA struct {
	R int
	G int
	B int
	A int
}
