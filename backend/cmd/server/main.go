package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/aussiebroadwan/tipping/backend/internal/db"
	"github.com/aussiebroadwan/tipping/backend/internal/services"

	"github.com/jackc/pgx/v5"
)

const (
	NRLCompetitionID                 = 111
	NRLWCompetitionID                = 161
	StateOfOriginCompetitionID       = 116
	StateOfOriginWomensCompetitionID = 156
)

var lg *slog.Logger

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

	if os.Getenv("NRL_API_BASE_URL") == "" {
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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := connect(ctx)
	if err != nil {
		lg.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	// Initialize services
	nrlService := services.NewNRLService(os.Getenv("NRL_API_BASE_URL"))
	dataService := services.NewNRLDataService(queries, ctx)

	// Define the competition IDs you want to fetch data for
	competitionIDs := []int64{
		NRLCompetitionID,
		NRLWCompetitionID,
		StateOfOriginCompetitionID,
		StateOfOriginWomensCompetitionID,
	}

	// Initialize and start the scheduled service
	scheduledService := services.NewNRLScheduledService(nrlService, dataService, competitionIDs)
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
