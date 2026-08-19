package app

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"tms/internal/auth"
	"tms/internal/config"
	"tms/internal/logger"
	"tms/internal/postgres"
	"tms/internal/project"
	users "tms/internal/user"

	"github.com/go-chi/chi/v5"
)

func Run() {
	ctx := context.Background()

	cfg := config.New()

	l := logger.New("production") // hard coded, gotta fix it later

	db, err := postgres.New(ctx, cfg.Postgres.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := users.NewRepository(db)

	authService := auth.NewService(userRepo)
	authHandler := auth.NewHandler(authService, l)

	projectRepo := project.NewRepository(db)
	projectService := project.NewService(projectRepo)
	projectHandler := project.NewHandler(projectService, l)

	r := chi.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
