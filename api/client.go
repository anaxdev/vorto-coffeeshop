package main

import (
	"log"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vorto-coffeeshop/api/pkg/config"
	delivery "github.com/vorto-coffeeshop/api/pkg/grpc"
	"github.com/vorto-coffeeshop/api/pkg/server"
	"github.com/vorto-coffeeshop/api/pkg/service"
)

func main() {
	config.InitVars()
	rpc, err := delivery.NewGRPCService(config.VarGrpcAddr)
	if err != nil {
		log.Fatalf("don't create service: %s", err)
	}

	r := mux.NewRouter()
	err = service.ConfigureHandlers(r, rpc)
	if err != nil {
		log.Fatalf("don't setting handlers: %s", err)
	}
	port, _ := strconv.Atoi(config.VarHttpPort)
	server.Run(r, server.Server{Hostname: "0.0.0.0", HTTPPort: port})
}
