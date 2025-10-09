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

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func runMigrations(dbURL string) error {
	// Convert postgresql:// to pgx5:// for golang-migrate
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
	// Get database URL from environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Create database connection pool
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close()

	// Test database connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("Database connected successfully")

	// Run database migrations
	if err := runMigrations(dbURL); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	log.Println("Migrations completed successfully")

	// Initialize sqlc queries
	queries := db.New(pool)

	// Initialize layers
	courtRepo := repository.NewCourtRepository(queries)
	courtService := service.NewCourtService(courtRepo)
	courtController := controller.NewCourtController(courtService)

	// Setup routes
	mux := http.NewServeMux()
	courtController.RegisterRoutes(mux)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
