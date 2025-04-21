package main

import (
	"database/sql"
	"log"

	"github.com/Aiyanu/simple-bank/api"
	db "github.com/Aiyanu/simple-bank/db/sqlc"
	"github.com/Aiyanu/simple-bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config: ", err)
	}
	dbDriver := config.DBDriver
	dbSource := config.DBSource
	serverAddress := config.ServerAddress

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("could not connect to database: ", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("could not start server: ", err)
	}
}
