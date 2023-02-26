package main

import (
	"context"
	"log"

	"github.com/Nuriddin-Olimjon/url-shortener/config"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/controller"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/repository"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	store := repository.NewStore(ctx, cfg.DBSource)

	runDBMigration(cfg.MigrationURL, cfg.DBSource)

	controller := controller.NewController(&cfg, store)

	err = controller.Start(cfg.HttpServerAddress)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatalf("cannot create new migrate instance: %s", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrate up: %s", err)
	}
}
