package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"paddletraffic/internal/api/handler"
	"paddletraffic/internal/api/middleware"
	"paddletraffic/internal/config"
	"paddletraffic/internal/database/generated/db"
	"paddletraffic/internal/repository"
	"paddletraffic/internal/service"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func runMigrations(dbURL, migrationPath string) error {
	migrationURL := strings.Replace(dbURL, "postgresql://", "pgx5://", 1)
	migrationURL = strings.Replace(migrationURL, "postgres://", "pgx5://", 1)

	m, err := migrate.New(migrationPath, migrationURL)
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
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	ctx := context.Background()
	poolConfig, err := pgxpool.ParseConfig(cfg.Database.URL)
	if err != nil {
		log.Fatal("Failed to parse database config:", err)
	}

	poolConfig.MaxConns = int32(cfg.Database.MaxConnections)

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("Database connected successfully")

	if err := runMigrations(cfg.Database.URL, cfg.Database.MigrationPath); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	log.Println("Migrations completed successfully")

	queries := db.New(pool)

	// Create repositories
	courtRepo := repository.NewCourtRepository(queries)
	statusRepo := repository.NewStatusRepository(queries)

	// Create services
	statusService := service.NewStatusService(statusRepo)
	courtService := service.NewCourtService(courtRepo, statusService)

	// Create handlers
	courtHandler := handler.NewCourtHandler(courtService)
	statusHandler := handler.NewStatusHandler(statusService)

	healthHandler := handler.NewHealthHandler(pool)

	r := chi.NewRouter()

	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(middleware.Timeout(30 * time.Second))

	healthHandler.RegisterRoutes(r)
	courtHandler.RegisterRoutes(r)
	statusHandler.RegisterRoutes(r)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	go func() {
		log.Printf("Server starting on :%s", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error:", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down server gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
