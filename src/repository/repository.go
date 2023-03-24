package repository

import "github.com/Faurby/Recibag/src/models"

type RecipeRepository interface {
	GetRecipe(name string) (*models.Recipe, error)
}
