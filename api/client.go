package main

import (
	"log"
	"strconv"

	"github.com/anaxdev/go-microservice/pkg/config"
	delivery "github.com/anaxdev/go-microservice/pkg/grpc"
	"github.com/anaxdev/go-microservice/pkg/server"
	"github.com/anaxdev/go-microservice/pkg/service"
	"github.com/gorilla/mux"
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
