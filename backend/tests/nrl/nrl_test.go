package nrl

import (
	"encoding/json"
	"testing"

	"github.com/aussiebroadwan/tipping/backend/internal/models"
	"github.com/aussiebroadwan/tipping/backend/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestNRLRound22Season2024(t *testing.T) {

	score1, score2 := 30, 48

	// Expected results
	expected := models.NRLFixture{
		ID:             "20241112210",
		RoundTitle:     "Round 22",
		MatchState:     "FullTime",
		KickOffTime:    "2024-08-01T09:50:00Z",
		Venue:          "Leichhardt Oval",
		VenueCity:      "Sydney",
		MatchCentreURL: "/draw/nrl-premiership/2024/round-22/wests-tigers-v-cowboys/",
		HomeTeam: models.NRLTeam{
			ID:    500023,
			Name:  "Wests Tigers",
			Odds:  nil,
			Score: &score1,
			Form:  nil,
		},
		AwayTeam: models.NRLTeam{
			ID:    500012,
			Name:  "Cowboys",
			Odds:  nil,
			Score: &score2,
			Form:  nil,
		},
	}

	// Actual results
	c := services.NewNRLService("https://nrl.com")

	actual, err := c.FetchFixtures(111, 22, 2024)
	if err != nil {
		t.Fatalf("Failed to fetch fixtures: %v", err)
	}

	if len(actual) == 0 {
		t.Fatal("Expected 0 fixtures, got", len(actual))
	}

	actualResult := actual[0]

	expectedStr, _ := json.MarshalIndent(expected, "", "  ")
	actualStr, _ := json.MarshalIndent(actualResult, "", "  ")

	assert.Equal(t, string(expectedStr), string(actualStr))
}

func TestNRLWRound5Season2024(t *testing.T) {

	score1, score2 := 16, 36

	// Expected results
	expected := models.NRLFixture{
		ID:             "20241610510",
		RoundTitle:     "Round 5",
		MatchState:     "FullTime",
		KickOffTime:    "2024-08-24T01:00:00Z",
		Venue:          "Eric Tweedale Stadium",
		VenueCity:      "Sydney",
		MatchCentreURL: "/draw/womens-premiership/2024/round-5/eels-v-knights/",
		HomeTeam: models.NRLTeam{
			ID:    500692,
			Name:  "Eels",
			Odds:  nil,
			Score: &score1,
			Form:  nil,
		},
		AwayTeam: models.NRLTeam{
			ID:    500691,
			Name:  "Knights",
			Odds:  nil,
			Score: &score2,
			Form:  nil,
		},
	}

	// Actual results
	c := services.NewNRLService("https://nrl.com")

	actual, err := c.FetchFixtures(161, 5, 2024)
	if err != nil {
		t.Fatalf("Failed to fetch fixtures: %v", err)
	}

	if len(actual) == 0 {
		t.Fatal("Expected 0 fixtures, got", len(actual))
	}

	actualResult := actual[0]

	expectedStr, _ := json.MarshalIndent(expected, "", "  ")
	actualStr, _ := json.MarshalIndent(actualResult, "", "  ")

	assert.Equal(t, string(expectedStr), string(actualStr))

}
