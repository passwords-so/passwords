package services

import (
	"context"
	"os"

	"github.com/dickeyy/passwords/api/log"
	"github.com/dickeyy/passwords/api/structs"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var DB *pg.DB

func ConnectDB(ctx context.Context) {
	DB = pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDRESS") + ":" + os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	err := DB.Ping(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
		os.Exit(1)
	}

	// create schema
	err = createSchema(DB)
	if err != nil {
		log.Error().Err(err).Msg("failed to create schema")
		os.Exit(1)
	}

	log.Info().Msg("connected to database")
}

func CloseDB() {
	DB.Close()
	log.Warn().Msg("closed database connection")
}

// createSchema creates database schema for User models.
func createSchema(db *pg.DB) error {
	// Check if the table already exists
	exists, err := db.Model((*structs.User)(nil)).Table().Exists()
	if err != nil && err.Error() != "ERROR #42P01 relation \"users\" does not exist" {
		return err
	}

	if !exists {
		err = db.Model((*structs.User)(nil)).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true, // This is a safeguard, but we're already checking existence
		})
		if err != nil {
			return err
		}
		log.Info().Msg("Users table created successfully")
	} else {
		log.Info().Msg("Users table already exists")
	}

	return nil
}
