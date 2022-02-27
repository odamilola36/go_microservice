package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"microservices/api/resthandlers"
	"microservices/api/routes"
	"microservices/pb"
	"net/http"
)

var (
	// Version is the current version of the API.
	Port int
	Host string
)

func init() {
	flag.IntVar(&Port, "port", 9000, "api service port")
	flag.StringVar(&Host, "auth-host", "localhost:9001", "auth service host")
	flag.Parse()
}

func main() {
	// Create a new gRPC client.
	conn, err := grpc.Dial(Host, grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewAuthServiceClient(conn)
	authHandler := resthandlers.NewAuthHandlers(client)
	authRoutes := routes.NewAuthRoutes(authHandler)

	router := mux.NewRouter().StrictSlash(true)
	routes.Install(router, authRoutes)

	log.Printf("Starting API server on port %d", Port)

	err = http.ListenAndServe(":"+fmt.Sprintf("%d", Port), routes.WithCors(router))
	if err != nil {
		log.Panicln(err)
		return
	}

}
