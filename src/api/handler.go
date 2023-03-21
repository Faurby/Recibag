package api

import (
	"context"

	pb "github.com/Faurby/Recibag/src/api/recipe_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedRecipeHandlerServer
}

func NewRecipeHandlerServer(gserver *grpc.Server) {
	pb.RegisterRecipeHandlerServer(gserver, &server{})
	reflection.Register(gserver)
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hej!"}, nil
}
