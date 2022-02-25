package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"microservices/db"
)

var (
	local bool
)

func init() {
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
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()
	log.Println("Connected to database")
}
