package models

type Unit struct {
	Name           string
	Army           string
	Models         []Model
	DefaultWeapons []string
	Power          int
	Location       Location
}