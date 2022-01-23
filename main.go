package main

import (
	"database/sql"
	"log"

	"github.com/dhruv42/bank/api"
	db "github.com/dhruv42/bank/db/sqlc"
	"github.com/dhruv42/bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start the server:", err)
	}
}
