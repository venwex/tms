package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"tms/internal/config"
	"tms/internal/postgres"
	"tms/internal/project"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()

	cfg := config.New()

	db, err := postgres.New(ctx, cfg.Postgres.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	projectRepo := project.NewRepository(db)
	projectService := project.NewService(projectRepo)

	r := chi.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
