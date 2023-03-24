package repository

import "github.com/Faurby/Recibag/src/models"

type IRecipeRepository interface {
	GetRecipe(name string) (*models.Recipe, error)
}
