package services

import (
	"github.com/Faurby/Recibag/src/entities"
	"github.com/Faurby/Recibag/src/repositories"
)

type RecipeUsecases interface {
	GetRecipe(name string) (*entities.Recipe, error)
}

type recipeUsecase struct {
	recipeRepo repositories.IRecipeRepository
}

func NewRecipeUsecase(r repositories.IRecipeRepository) RecipeUsecases {
	return &recipeUsecase{r}
}

// CRUD operations
func (r *recipeUsecase) GetRecipe(name string) (*entities.Recipe, error) {
	return r.recipeRepo.GetRecipe(name)
}
