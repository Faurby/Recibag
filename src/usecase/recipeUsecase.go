package usecase

import (
	"errors"

	"github.com/Faurby/Recibag/src/models"
	"github.com/Faurby/Recibag/src/repository"
)

type RecipeUsecases interface {
	GetRecipe(name string) (*models.Recipe, error)
}

type recipeUsecase struct {
	recipeRepo repository.IRecipeRepository
}

var chickenRecipe = &models.Recipe{
	Name:        "Kylling",
	Description: "En god kyllingret med MASSERE af salt",
	Persons:     2,
	Ingredients: []*models.Ingredient{
		{
			Name:   "Kylling",
			Amount: 100,
			Unit:   "g",
		},
		{
			Name:   "Salt",
			Amount: 1,
			Unit:   "tsk",
		},
	},
	Instructions: []string{
		"Krydre kyllingen",
		"Steg kyllingen",
	},
}

func NewRecipeUsecase(r repository.IRecipeRepository) RecipeUsecases {
	return &recipeUsecase{r}
}

// CRUD operations
func (r *recipeUsecase) GetRecipe(name string) (*models.Recipe, error) {
	if name == "Kylling" {
		return chickenRecipe, nil
	} else {
		return nil, errors.New("Recipe not found")
	}
	// return r.recipeRepo.GetRecipe(name)
}
