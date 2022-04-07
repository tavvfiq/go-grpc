package main

import (
	"context"
	"fmt"
	"grpc_service/internal/common/config"
	"grpc_service/internal/common/model"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.BookingListByUser

func init() {
	localStorage = new(model.BookingListByUser)
	localStorage.List = make(map[string]*model.BookingList)
}

type BookingService struct{}

var counter int

func (b *BookingService) Create(_ context.Context, param *model.BookingDetailAndUserId) (*emptypb.Empty, error) {
	counter++
	userId := param.Id
	bookingDetail := param.Detail
	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = &model.BookingList{}
		localStorage.List[userId].List = make([]*model.Booking, 0)
	}
	booking := &model.Booking{
		Id:     fmt.Sprintf("b%d", counter),
		UserId: userId,
		Detail: bookingDetail,
	}
	log.Println("create booking with detail:", bookingDetail.String(), "for user:", userId)
	localStorage.List[userId].List = append(localStorage.List[userId].List, booking)
	return &emptypb.Empty{}, nil
}

func (b *BookingService) List(_ context.Context, param *model.UserId) (*model.BookingList, error) {
	return localStorage.List[param.Id], nil
}

func main() {
	srv := grpc.NewServer()
	var bookingSrv BookingService
	model.RegisterBookingsServer(srv, &bookingSrv)
	log.Println("Starting RPC server at", config.SERVICE_BOOKING_PORT)
	listener, err := net.Listen("tcp", config.SERVICE_BOOKING_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_BOOKING_PORT, err)
	}
	log.Fatal(srv.Serve(listener))
}
