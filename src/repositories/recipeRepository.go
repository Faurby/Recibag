package repositories

import (
	"errors"

	"github.com/Faurby/Recibag/src/entities"
)

type IRecipeRepository interface {
	GetRecipe(name string) (*entities.Recipe, error)
}

type RecipeRepository struct {
}

func NewRecipeRepository() IRecipeRepository {
	return &RecipeRepository{}
}

var chickenRecipe = &entities.Recipe{
	Name:        "Kylling",
	Description: "Protein baby",
	Persons:     2,
	Ingredients: []*entities.Ingredient{
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

func (r *RecipeRepository) GetRecipe(name string) (*entities.Recipe, error) {
	if name == chickenRecipe.Name {
		return chickenRecipe, nil
	}
	return nil, errors.New("Recipe not found")
}
