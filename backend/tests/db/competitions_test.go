package db

import (
	"context"
	"testing"
)

func TestListCompetitions(t *testing.T) {
	ctx := context.Background()

	competitions, err := testQueries.ListCompetitions(ctx)
	if err != nil {
		t.Fatalf("Failed to list competitions: %v", err)
	}

	expectedCount := 4 // Based on the assumption that there are 4 static competitions
	if len(competitions) != expectedCount {
		t.Fatalf("Expected %d competitions, got %d", expectedCount, len(competitions))
	}
}

func TestGetCompetitionByID(t *testing.T) {
	ctx := context.Background()

	competitionID := int32(111) // NRL
	competition, err := testQueries.GetCompetitionByID(ctx, competitionID)
	if err != nil {
		t.Fatalf("Failed to get competition by ID: %v", err)
	}

	if competition.ID != competitionID {
		t.Fatalf("Expected competition ID %d, got %d", competitionID, competition.ID)
	}
}
