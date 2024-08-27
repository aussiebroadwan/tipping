package db

import (
	"context"
	"testing"

	"github.com/aussiebroadwan/tipping/backend/internal/db"
)

func TestCreateMatchDetail(t *testing.T) {
	ctx := context.Background()

	odd1 := float64(1.54)
	odd2 := float64(2.49)

	homeTeam := db.CreateTeamParams{
		TeamID:   500001,
		Nickname: "homeTeam",
	}

	_, err := testQueries.CreateTeam(ctx, homeTeam)
	if err != nil {
		t.Fatalf("Failed to create home team: %v", err)
	}

	awayTeam := db.CreateTeamParams{
		TeamID:   500002,
		Nickname: "awayTeam",
	}

	_, err = testQueries.CreateTeam(ctx, awayTeam)
	if err != nil {
		t.Fatalf("Failed to create away team: %v", err)
	}

	// Assume a match detail with fixture ID 1 exists and is an upcoming match
	arg := db.CreateMatchDetailParams{
		FixtureID:     1,
		HometeamID:    500001,
		AwayteamID:    500002,
		HometeamOdds:  &odd1,
		AwayteamOdds:  &odd2,
		HometeamScore: nil,
		AwayteamScore: nil,
		HometeamForm:  "WLWLW",
		AwayteamForm:  "LWLWL",
		WinnerTeamid:  nil,
	}

	matchDetail, err := testQueries.CreateMatchDetail(ctx, arg)
	if err != nil {
		t.Fatalf("Failed to create match detail: %v", err)
	}

	if matchDetail.FixtureID != arg.FixtureID {
		t.Fatalf("Expected fixture ID %d, got %d", arg.FixtureID, matchDetail.FixtureID)
	}
}

func TestUpdateMatchDetail(t *testing.T) {
	ctx := context.Background()

	score1 := int32(35)
	score2 := int32(28)

	// Assume a match detail with fixture ID 1 exists
	fixtureID := int64(1)
	arg := db.UpdateMatchDetailParams{
		FixtureID:     fixtureID,
		HomeTeamScore: &score1,
		AwayTeamScore: &score2,
	}

	matchDetail, err := testQueries.UpdateMatchDetail(ctx, arg)
	if err != nil {
		t.Fatalf("Failed to update match detail: %v", err)
	}

	if *matchDetail.HometeamScore != 35 || *matchDetail.AwayteamScore != 28 {
		t.Fatalf("Unexpected match detail scores: %+v", matchDetail)
	}
}

func TestGetMatchDetailsByFixtureID(t *testing.T) {
	ctx := context.Background()

	// Assume a match detail with fixture ID 1 exists
	fixtureID := int64(1)
	matchDetail, err := testQueries.GetMatchDetailsByFixtureID(ctx, fixtureID)
	if err != nil {
		t.Fatalf("Failed to get match details by fixture ID: %v", err)
	}

	if matchDetail.Fixture.ID != fixtureID {
		t.Fatalf("Expected fixture ID %d, got %d", fixtureID, matchDetail.Fixture.ID)
	}
}

func TestListMatchDetails(t *testing.T) {
	ctx := context.Background()

	_, err := testQueries.ListMatchDetails(ctx)
	if err != nil {
		t.Fatalf("Failed to list match details: %v", err)
	}
}

func TestListMatchDetailsByCompetitionID(t *testing.T) {
	ctx := context.Background()

	competitionID := int64(111) // NRL
	matchDetails, err := testQueries.ListMatchDetailsByCompetitionID(ctx, competitionID)
	if err != nil {
		t.Fatalf("Failed to list match details by competition ID: %v", err)
	}

	if len(matchDetails) == 0 {
		t.Fatalf("Expected at least one match detail, got 0")
	}
}
