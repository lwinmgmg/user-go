package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	pb "github.com/lwinmgmg/grpc_m/user_go"
	"github.com/lwinmgmg/user-go/env"
	"google.golang.org/grpc"
)

type userService struct {
	pb.UnimplementedUserServiceServer
}

func (*userService) CheckToken(context.Context, *pb.Token) (*pb.InternalUser, error) {
	slog.Info("Get a request")
	return &pb.InternalUser{
		Username: "",
		Email:    "",
		Code:     "A00001",
	}, nil
}

func main() {
	settings, err := env.LoadSettings()
	if err != nil {
		panic(err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", settings.GrpcServer.Host, settings.GrpcServer.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &userService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
