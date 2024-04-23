package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	pb "github.com/lwinmgmg/grpc_m/user_go"
	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/controller"
	"google.golang.org/grpc"
)

type userService struct {
	pb.UnimplementedUserServiceServer
	ctrl *controller.Controller
}

func (us *userService) CheckThirdpartyToken(ctx context.Context, tkn *pb.Token) (*pb.ThirdpartySubject, error) {
	slog.Info("Get a request")
	sub, err := us.ctrl.CheckThirdPartyTkn(tkn.Token)
	if err != nil {
		return nil, err
	}
	return &pb.ThirdpartySubject{
		Uid: sub.UserID,
		Cid: sub.ClientID,
		Scp: sub.Scopes,
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
	pb.RegisterUserServiceServer(s, &userService{
		ctrl: controller.NewContoller(settings),
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
