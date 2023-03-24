package api

import (
	"context"
	"log"

	pb "github.com/Faurby/Recibag/src/api/recipe_grpc"
	"github.com/Faurby/Recibag/src/models"
	"github.com/Faurby/Recibag/src/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedRecipeHandlerServer
	recipeUsecase usecase.RecipeUsecases
}

// Depends on the interface 'RecipeUsecases' from the usecase package.
func NewRecipeHandlerServer(gserver *grpc.Server, r usecase.RecipeUsecases) {
	recipeServer := &server{recipeUsecase: r}
	pb.RegisterRecipeHandlerServer(gserver, recipeServer)

	// Register reflection service on gRPC server. This is needed for the grpcurl tool.
	reflection.Register(gserver)
}

func (s *server) GetRecipe(ctx context.Context, in *pb.RecipeRequest) (*pb.Recipe, error) {
	recipe, err := s.recipeUsecase.GetRecipe(in.Name)
	if err != nil {
		log.Printf("Error getting recipe:" + err.Error())
		return nil, err
	}
	return s.TransformRecipeRPC(recipe), nil
}

// Helper functions for transforming between the gRPC and the internal models.
func (s *server) TransformRecipeRPC(r *models.Recipe) *pb.Recipe {
	return &pb.Recipe{
		Name:         r.Name,
		Description:  r.Description,
		Persons:      int32(r.Persons),
		Ingredients:  (s.TransformIngredientsRPC(r.Ingredients)),
		Instructions: r.Instructions,
	}
}

func (s *server) TransformIngredientsRPC(i []*models.Ingredient) (ingredients []*pb.Ingredient) {
	for _, ingredient := range i {
		ingredients = append(ingredients, &pb.Ingredient{
			Name:   ingredient.Name,
			Amount: ingredient.Amount,
			Unit:   ingredient.Unit,
		})
	}
	return ingredients
}
