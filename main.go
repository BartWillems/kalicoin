package main

import (
	"os"

	"gitlab.com/bartwillems/kalicoin/pkg/api"
	"gitlab.com/bartwillems/kalicoin/pkg/db"
	"gitlab.com/bartwillems/kalicoin/pkg/jaeger"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	defer db.Conn.Close()

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	closer, err := jaeger.Init()

	if err != nil {
		log.Fatal(err)
	}

	defer closer.Close()

	router := api.New(db.Conn)

	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
