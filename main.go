package main

import (
	"log"
	"net"

	"github.com/Faurby/Recibag/src/handlers"
	"github.com/Faurby/Recibag/src/repositories"
	"github.com/Faurby/Recibag/src/services"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	ru := services.NewRecipeUsecase(repositories.NewRecipeRepository())

	handlers.NewRecipeHandlerServer(s, ru)
	log.Printf("server listening at %v", lis.Addr())

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server")
	}
}
