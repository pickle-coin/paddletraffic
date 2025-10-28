package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"paddletraffic/internal/controller"
	"paddletraffic/internal/database/generated/db"
	"paddletraffic/internal/repository"
	"paddletraffic/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func runMigrations(dbURL string) error {
	migrationURL := strings.Replace(dbURL, "postgresql://", "pgx5://", 1)
	migrationURL = strings.Replace(migrationURL, "postgres://", "pgx5://", 1)

	m, err := migrate.New(
		"file://internal/database/migrations",
		migrationURL,
	)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("Database connected successfully")

	if err := runMigrations(dbURL); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	log.Println("Migrations completed successfully")

	queries := db.New(pool)

	courtRepo := repository.NewCourtRepository(queries)
	courtService := service.NewCourtService(courtRepo)
	courtController := controller.NewCourtController(courtService)
	healthController := controller.NewHealthController(pool)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	healthController.RegisterRoutes(r)
	courtController.RegisterRoutes(r)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
