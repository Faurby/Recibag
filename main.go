package main

import (
	"log"
	"net"

	"github.com/Faurby/Recibag/src/api"
	"github.com/Faurby/Recibag/src/usecase"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	ru := usecase.NewRecipeUsecase(nil)

	api.NewRecipeHandlerServer(s, ru)
	log.Printf("server listening at %v", lis.Addr())

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server")
	}
}
