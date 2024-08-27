package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/aussiebroadwan/tipping/backend/config"
	"github.com/aussiebroadwan/tipping/backend/internal/db"
	"github.com/aussiebroadwan/tipping/backend/internal/handlers"
	"github.com/aussiebroadwan/tipping/backend/internal/services"

	_ "github.com/aussiebroadwan/tipping/backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/jackc/pgx/v5"
)

var (
	lg *slog.Logger

	apiBase    string
	nrlApiBase string
)

func init() {
	// Create logger
	lg = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Check for required environment variables
	requiredDbEnvVars := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
	}

	numMissing := 0
	for _, envVar := range requiredDbEnvVars {
		if os.Getenv(envVar) == "" {
			lg.Error("Environment variable " + envVar + " is required")
			numMissing++
		}
	}

	if numMissing > 0 {
		os.Exit(1)
	}

	if apiBase = os.Getenv("API_BASE_URL"); apiBase == "" {
		lg.Warn("API_BASE_URL environment variable is not given, defaulting to http://localhost:8080")
		apiBase = "http://localhost:8080"
	}

	if nrlApiBase = os.Getenv("NRL_API_BASE_URL"); nrlApiBase == "" {
		lg.Error("NRL_API_BASE_URL environment variable is required")
		os.Exit(1)
	}
}

func connect(ctx context.Context) (*pgx.Conn, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	conn, err := pgx.Connect(ctx, psqlInfo)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// @title Tipping API
// @version 1.0
// @description This is the API for the Tipping Application to interact with NRL data.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := connect(ctx)
	if err != nil {
		lg.Error(fmt.Sprintf("Failed to connect to database: %s", err.Error()))
		os.Exit(1)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	// Initialize services
	nrlService := services.NewNRLService(os.Getenv("NRL_API_BASE_URL"))
	nrlDataService := services.NewNRLDataService(queries, ctx)
	apiDataService := services.NewAPIDataService(queries, ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", GetHealth)
	mux.HandleFunc("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL(apiBase+"/swagger/doc.json"),
	))
	handlers.RegisterRoutes(mux, apiDataService)
	go http.ListenAndServe(":8080", mux)

	// Define the competition IDs you want to fetch data for
	competitionIDs := []int64{
		config.CompetitionNRL,
		config.CompetitionNRLW,
		config.CompetitionStateOfOrigin,
		config.CompetitionStateOfOriginWomens,
	}

	// Initialize and start the scheduled service
	scheduledService := services.NewNRLScheduledService(nrlService, nrlDataService, competitionIDs)
	go scheduledService.Start(ctx)

	// Signal handler for graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		cancel()
	}()

	// Block until context is done
	<-ctx.Done()
	log.Println("Shutting down...")
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
