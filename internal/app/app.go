package app

import (
	"context"
	"log"
	"net/http"
	"tms/internal/config"
	"tms/internal/logger"
	"tms/internal/postgres"
	"tms/internal/project"

	"github.com/go-chi/chi/v5"
)

func Run() {
	ctx := context.Background()

	cfg := config.New()

	l := logger.New("") // исправить позже

	db, err := postgres.New(ctx, cfg.Postgres.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	projectRepo := project.NewRepository(db)
	projectService := project.NewService(projectRepo)
	projectHandler := project.NewHandler(projectService)

	r := chi.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
