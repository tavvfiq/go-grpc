package main

import (
	"context"
	"grpc_service/internal/common/config"
	"grpc_service/internal/common/model"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UserService struct{}

func (u *UserService) Register(_ context.Context, param *model.User) (*emptypb.Empty, error) {
	localStorage.List = append(localStorage.List, param)
	log.Printf("Registering user: %s\n", param.String())
	return new(emptypb.Empty), nil
}

func (u *UserService) List(_ context.Context, _ *emptypb.Empty) (*model.UserList, error) {
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	var userSrv UserService
	model.RegisterUsersServer(srv, &userSrv)
	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)
	listener, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}
	log.Fatal(srv.Serve(listener))
}
