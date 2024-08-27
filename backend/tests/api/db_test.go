package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/aussiebroadwan/tipping/backend/internal/db"
	"github.com/aussiebroadwan/tipping/backend/internal/handlers"
	"github.com/aussiebroadwan/tipping/backend/internal/services"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	testDB        *pgxpool.Pool
	testQueries   *db.Queries
	handlerRouter *http.ServeMux
	dataService   *services.APIDataService
)

// TestMain sets up the test environment by initializing the PostgreSQL container,
// running migrations, and tearing down the environment after tests are completed.
func TestMain(m *testing.M) {
	ctx := context.Background()

	// Start the PostgreSQL container
	container, err := startPostgresContainer(ctx)
	if err != nil {
		log.Fatalf("Could not start PostgreSQL container: %v", err)
	}
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			log.Fatalf("Could not terminate PostgreSQL container: %v", err)
		}
	}()

	// Connect to the database
	host, port, err := getContainerHostPort(container, ctx)
	if err != nil {
		log.Fatalf("Could not get container host/port: %v", err)
	}

	dsn := fmt.Sprintf("postgres://postgres:password@%s:%s/testdb?sslmode=disable", host, port)
	testDB, err = connectToDatabase(ctx, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer testDB.Close()

	// Run migrations
	if err := runMigrations(dsn); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialise test queries
	testQueries = db.New(testDB)
	seedDatabase()

	// Initialise API data service for testing
	dataService = services.NewAPIDataService(testQueries, context.Background())

	// Initialise handler router for testing the API requests
	handlerRouter = http.NewServeMux()
	handlers.RegisterRoutes(handlerRouter, dataService)

	// Run tests
	os.Exit(m.Run())
}

// startPostgresContainer starts a PostgreSQL container using testcontainers.
func startPostgresContainer(ctx context.Context) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15.3-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(60 * time.Second),
	}
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

// getContainerHostPort retrieves the host and mapped port of the PostgreSQL container.
func getContainerHostPort(container testcontainers.Container, ctx context.Context) (string, string, error) {
	host, err := container.Host(ctx)
	if err != nil {
		return "", "", fmt.Errorf("could not get container host: %w", err)
	}

	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return "", "", fmt.Errorf("could not get mapped port: %w", err)
	}

	return host, mappedPort.Port(), nil
}

// connectToDatabase establishes a connection to the PostgreSQL database.
func connectToDatabase(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DSN: %w", err)
	}

	dbPool, err := pgxpool.New(ctx, config.ConnString())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return dbPool, nil
}

// runMigrations applies the database migrations using golang-migrate.
func runMigrations(dsn string) error {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database for migrations: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../internal/db/migration", // Path to the migrations directory
		"postgres",                           // Database name
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
