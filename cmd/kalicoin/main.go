package main

import (
	"kalicoin/pkg/db"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	if err := db.Migrate(); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	db.Conn.Close()

	os.Exit(0)
}
