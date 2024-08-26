package db

import (
	"context"
	"testing"

	"github.com/aussiebroadwan/tipping/backend/internal/db"
)

func TestListTeams(t *testing.T) {
	ctx := context.Background()

	_, err := testQueries.ListTeams(ctx)
	if err != nil {
		t.Fatalf("Failed to list teams: %v", err)
	}
}

func TestCreateTeam(t *testing.T) {
	ctx := context.Background()

	arg := db.CreateTeamParams{
		TeamID:   500012,
		Nickname: "Cowboys",
	}

	team, err := testQueries.CreateTeam(ctx, arg)
	if err != nil {
		t.Fatalf("Failed to create team: %v", err)
	}

	if team.TeamID != arg.TeamID || team.Nickname != arg.Nickname {
		t.Fatalf("Unexpected team data: %+v", team)
	}
}

func TestGetTeamByID(t *testing.T) {
	ctx := context.Background()

	teamID := int32(500012) // Cowboys, from the previous test
	team, err := testQueries.GetTeamByID(ctx, teamID)
	if err != nil {
		t.Fatalf("Failed to get team by ID: %v", err)
	}

	if team.TeamID != teamID {
		t.Fatalf("Expected team ID %d, got %d", teamID, team.TeamID)
	}
}
