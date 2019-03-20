package db

import (
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
	log "github.com/sirupsen/logrus"
)

// Conn is the shared database connection pool
var Conn *pop.Connection

// Connect initialises the `Conn` connection
func Connect() error {
	var env = envy.Get("env", "development")
	var err error
	Conn, err = pop.Connect(env)

	if err != nil {
		return err
	}

	log.Info("Established DB connection.")
	pop.Debug = env == "development"
	return nil
}

// Migrate tries to run the migrations
func Migrate() error {
	fileMigrator, err := pop.NewFileMigrator("./migrations", Conn)

	if err != nil {
		return err
	}

	fileMigrator.Status()

	return fileMigrator.Up()
}