package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/vorto-coffeeshop/api/pkg/config"
	delivery "github.com/vorto-coffeeshop/api/pkg/grpc"
	"github.com/vorto-coffeeshop/api/pkg/service"
)

func main() {
	config.InitVars()
	conn, err := config.Connect()
	if err != nil {
		panic(err)
	}

	deliveryService := service.NewService(conn)
	deliveryController := delivery.NewDeliveryController(deliveryService)

	// start a gRPC server
	server := grpc.NewServer()
	delivery.RegisterDeliveryServiceServer(server, deliveryController)
	reflection.Register(server)
	con, err := net.Listen("tcp", config.VarGrpcAddr)
	if err != nil {
		panic(err)
	}
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
