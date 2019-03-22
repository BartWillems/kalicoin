package main

import (
	"kalicoin/pkg/api"
	"kalicoin/pkg/db"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	defer db.Conn.Close()

	if err := db.Migrate(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	router := api.New(db.Conn)

	if err := router.Run(":8000"); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	os.Exit(0)
}
