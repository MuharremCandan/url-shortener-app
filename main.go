package main

import (
	"log"

	"github.com/MuharremCandan/url-shortenerapp/config"
	"github.com/MuharremCandan/url-shortenerapp/database"
	"github.com/MuharremCandan/url-shortenerapp/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	pgDb := database.NewPostgreDB(cfg)

	db, err := pgDb.ConnectDB()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	s := server.NewFiberServer(cfg, db)

	s.Start()

}
