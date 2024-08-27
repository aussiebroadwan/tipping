package nrl

import (
	"context"
	"strconv"
	"testing"

	"github.com/aussiebroadwan/tipping/backend/internal/models"
	"github.com/aussiebroadwan/tipping/backend/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestStoreAndFetchNRLRound22Season2024(t *testing.T) {
	ctx := context.Background()

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
	dataService := services.NewNRLDataService(testQueries, ctx) // Assuming testQueries is already initialized with TestMain

	actual, err := c.FetchFixtures(111, 22, 2024)
	if err != nil {
		t.Fatalf("Failed to fetch fixtures: %v", err)
	}

	if len(actual) == 0 {
		t.Fatal("Expected 0 fixtures, got", len(actual))
	}

	actualResult := actual[0]

	// Insert fetched data into the database
	err = dataService.StoreFixtureAndDetails(actualResult)
	if err != nil {
		t.Fatalf("Failed to store fetched data into the database: %v", err)
	}

	// Fetch the stored fixture from the database
	storedFixture, err := testQueries.GetFixtureByID(ctx, parseFixtureID(expected.ID))
	if err != nil {
		t.Fatalf("Failed to fetch stored fixture from the database: %v", err)
	}

	storedMatchDetails, err := testQueries.GetMatchDetailsByFixtureID(ctx, parseFixtureID(expected.ID))
	if err != nil {
		t.Fatalf("Failed to fetch stored match details from the database: %v", err)
	}

	// Verify stored data
	assert.Equal(t, expected.RoundTitle, storedFixture.Roundtitle)
	assert.Equal(t, expected.MatchState, storedFixture.Matchstate)
	assert.Equal(t, expected.Venue, storedFixture.Venue)
	assert.Equal(t, expected.VenueCity, storedFixture.Venuecity)
	assert.Equal(t, expected.MatchCentreURL, storedFixture.Matchcentreurl)

	assert.Equal(t, *parseScore(expected.HomeTeam.Score), *storedMatchDetails.MatchDetail.HometeamScore)
	assert.Equal(t, *parseScore(expected.AwayTeam.Score), *storedMatchDetails.MatchDetail.AwayteamScore)

	assert.Equal(t, expected.KickOffTime, storedFixture.Kickofftime.Time.Format("2006-01-02T15:04:05Z"))
}

func TestStoreAndFetchNRLWRound5Season2024(t *testing.T) {
	ctx := context.Background()

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
	dataService := services.NewNRLDataService(testQueries, ctx) // Assuming testQueries is already initialized with TestMain

	actual, err := c.FetchFixtures(161, 5, 2024)
	if err != nil {
		t.Fatalf("Failed to fetch fixtures: %v", err)
	}

	if len(actual) == 0 {
		t.Fatal("Expected 0 fixtures, got", len(actual))
	}

	actualResult := actual[0]

	// Insert fetched data into the database
	err = dataService.StoreFixtureAndDetails(actualResult)
	if err != nil {
		t.Fatalf("Failed to store fetched data into the database: %v", err)
	}

	// Fetch the stored fixture from the database
	storedFixture, err := testQueries.GetFixtureByID(ctx, parseFixtureID(expected.ID))
	if err != nil {
		t.Fatalf("Failed to fetch stored fixture from the database: %v", err)
	}

	storedMatchDetails, err := testQueries.GetMatchDetailsByFixtureID(ctx, parseFixtureID(expected.ID))
	if err != nil {
		t.Fatalf("Failed to fetch stored match details from the database: %v", err)
	}

	// Verify stored data
	assert.Equal(t, expected.RoundTitle, storedFixture.Roundtitle)
	assert.Equal(t, expected.MatchState, storedFixture.Matchstate)
	assert.Equal(t, expected.Venue, storedFixture.Venue)
	assert.Equal(t, expected.VenueCity, storedFixture.Venuecity)
	assert.Equal(t, expected.MatchCentreURL, storedFixture.Matchcentreurl)

	assert.Equal(t, *parseScore(expected.HomeTeam.Score), *storedMatchDetails.MatchDetail.HometeamScore)
	assert.Equal(t, *parseScore(expected.AwayTeam.Score), *storedMatchDetails.MatchDetail.AwayteamScore)

	assert.Equal(t, expected.KickOffTime, storedFixture.Kickofftime.Time.Format("2006-01-02T15:04:05Z"))
}

// Helper function to convert fixture ID to int64
func parseFixtureID(id string) int64 {
	fixtureID, _ := strconv.ParseInt(id, 10, 64)
	return fixtureID
}

// Helper function to parse score
func parseScore(score *int) *int32 {
	if score == nil {
		return nil
	}
	val := int32(*score)
	return &val
}
