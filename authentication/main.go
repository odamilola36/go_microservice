package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"microservices/authentication/repositories"
	"microservices/authentication/service"
	"microservices/dbCon"
	"microservices/pb"
	"net"
)

var (
	local bool
	port  int
)

func init() {
	flag.IntVar(&port, "port", 9001, "port to run auth server on")
	flag.BoolVar(&local, "local", true, "Run the service locally")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}
	cfg := dbCon.NewConfig()
	conn, err := dbCon.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	defer func(conn dbCon.Connection, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, conn.DBContext())
	log.Println("Connected to database")

	userRepo := repositories.NewUsersRepository(conn)
	authService := service.NewAuthService(userRepo)

	log.Println("Starting auth server on port: ", port)
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panicln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)

	err = grpcServer.Serve(server)

	log.Println("Auth server started")
	if err != nil {
		return
	}

}
