package db

import (
	"context"
	"testing"
	"time"

	"github.com/aussiebroadwan/tipping/backend/internal/db"
	"github.com/jackc/pgx/v5/pgtype"
)

func TestListFixtures(t *testing.T) {
	ctx := context.Background()

	_, err := testQueries.ListFixtures(ctx)
	if err != nil {
		t.Fatalf("Failed to list fixtures: %v", err)
	}
}

func TestCreateFixture(t *testing.T) {
	ctx := context.Background()

	arg := db.CreateFixtureParams{
		ID:             1,
		CompetitionID:  111,
		Roundtitle:     "Round 1",
		Matchstate:     "Upcoming",
		Venue:          "Stadium A",
		Venuecity:      "City A",
		Matchcentreurl: "http://example.com/match/1",
		Kickofftime:    pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	fixture, err := testQueries.CreateFixture(ctx, arg)
	if err != nil {
		t.Fatalf("Failed to create fixture: %v", err)
	}

	if fixture.CompetitionID != arg.CompetitionID || fixture.Roundtitle != arg.Roundtitle {
		t.Fatalf("Unexpected fixture data: %+v", fixture)
	}
}

func TestGetFixtureByID(t *testing.T) {
	ctx := context.Background()

	// Assuming a fixture with ID 1 exists
	fixtureID := int64(1)
	fixture, err := testQueries.GetFixtureByID(ctx, fixtureID)
	if err != nil {
		t.Fatalf("Failed to get fixture by ID: %v", err)
	}

	if fixture.ID != fixtureID {
		t.Fatalf("Expected fixture ID %d, got %d", fixtureID, fixture.ID)
	}
}

func TestGetFixturesByCompetitionID(t *testing.T) {
	ctx := context.Background()

	competitionID := int64(111) // NRL
	fixtures, err := testQueries.GetFixturesByCompetitionID(ctx, competitionID)
	if err != nil {
		t.Fatalf("Failed to get fixtures by competition ID: %v", err)
	}

	if len(fixtures) == 0 {
		t.Fatalf("Expected at least one fixture, got 0")
	}
}

func TestUpdateFixture(t *testing.T) {
	ctx := context.Background()

	// Assume a fixture with ID 1 exists
	fixtureID := int64(1)
	value := "FullTime"
	arg := db.UpdateFixtureParams{
		ID:         fixtureID,
		MatchState: &value,
	}

	fixture, err := testQueries.UpdateFixture(ctx, arg)
	if err != nil {
		t.Fatalf("Failed to update fixture: %v", err)
	}

	if fixture.Roundtitle != "Round 1" {
		t.Fatalf("Expected updated round title 'Round 2', got '%s'", fixture.Roundtitle)
	}
}
