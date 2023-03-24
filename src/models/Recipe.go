package models

type Recipe struct {
	Name         string
	Description  string
	Persons      int
	Ingredients  []*Ingredient
	Instructions []string
}

// Ingredient is a struct that represents an ingredient in a recipe.
type Ingredient struct {
	Name   string
	Amount float64
	Unit   string
}
