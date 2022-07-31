package models

type Token struct {
	RGBA RGBA
	ID   string
}

type RGBA struct {
	R int
	G int
	B int
	A int
}
