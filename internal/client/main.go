package main

import (
	"context"
	"fmt"
	"grpc_service/internal/common/config"
	"grpc_service/internal/common/model"
	"log"

	"google.golang.org/grpc"
)

func NewUserService() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s with error %v\n", port, err)
	}
	return model.NewUsersClient(conn)
}

func NewBookingService() model.BookingsClient {
	port := config.SERVICE_BOOKING_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s with error %v\n", port, err)
	}
	return model.NewBookingsClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "u001",
		Name:     "Noval Agung",
		Password: "kw8d hl12/3m,a",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}
	detail := model.BookingDetail{
		Name:       "Justin Bieber Concert",
		TotalPrice: 6000000,
		Type:       model.BookingType(model.BookingType_value["CONCERT"]),
	}
	param := model.BookingDetailAndUserId{
		Id:     user1.Id,
		Detail: &detail,
	}
	c := NewUserService()
	cc := NewBookingService()
	fmt.Println("testing user service gRPC")
	c.Register(context.Background(), &user1)
	fmt.Println("testing booking service gRPC")
	cc.Create(context.Background(), &param)
}
