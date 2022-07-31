package models

type Squad struct {
	Name           string
	Army           string
	Units          []Unit
	DefaultWeapons []string
	Power          int
	Location       Location
}