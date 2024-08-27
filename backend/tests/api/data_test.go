package api

import (
	"testing"
)

func TestGetCompetitions(t *testing.T) {
	competitions, err := dataService.GetCompetitions()
	if err != nil {
		t.Fatalf("Failed to get competitions: %v", err)
	}

	expectedCount := 4 // Based on the assumption that there are 4 static competitions
	if len(competitions) != expectedCount {
		t.Fatalf("Expected %d competitions, got %d", expectedCount, len(competitions))
	}

	// Check if the competitions are the expected ones
	expectedCompetitionIDs := []int64{111, 161, 116, 156}
	for i, competition := range competitions {
		if competition.ID != expectedCompetitionIDs[i] {
			t.Fatalf("Expected competition ID %d, got %d", expectedCompetitionIDs[i], competition.ID)
		}
	}
}

func TestGetFixtures(t *testing.T) {
	fixtures, err := dataService.GetFixtures()
	if err != nil {
		t.Fatalf("Failed to get fixtures: %v", err)
	}

	if len(fixtures) != 3 {
		t.Fatal("Expected 3 fixtures, got", len(fixtures))
	}
}

func TestGetFixtureByID(t *testing.T) {
	fixtureID := int64(20241112610)
	fixture, err := dataService.GetFixtureDetails(fixtureID)
	if err != nil {
		t.Fatalf("Failed to get fixture by ID: %v", err)
	}

	if fixture.ID != fixtureID {
		t.Fatalf("Expected fixture ID %d, got %d", fixtureID, fixture.ID)
	}

	if fixture.HomeTeam.Form != "WLWWW" {
		t.Fatalf("Unexpected home team form: %s", fixture.HomeTeam.Form)
	}

	if fixture.AwayTeam.Form != "WWWLW" {
		t.Fatalf("Unexpected away team form: %s", fixture.AwayTeam.Form)
	}
}

func TestGetCompetitionFixtures(t *testing.T) {
	competitionID := int64(111) // NRL
	fixtures, err := dataService.GetCompetitionFixtures(competitionID)
	if err != nil {
		t.Fatalf("Failed to get fixtures by competition ID: %v", err)
	}

	if len(fixtures) != 2 {
		t.Fatalf("Expected 2 fixtures, got %d", len(fixtures))
	}
}

func TestCompetitionsFixturesNRLW(t *testing.T) {
	competitionID := int64(161) // NRLW
	fixtures, err := dataService.GetCompetitionFixtures(competitionID)
	if err != nil {
		t.Fatalf("Failed to get fixtures by competition ID: %v", err)
	}

	if len(fixtures) != 1 {
		t.Fatalf("Expected 1 fixture, got %d", len(fixtures))
	}
}
