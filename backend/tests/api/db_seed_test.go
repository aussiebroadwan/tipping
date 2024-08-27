package api

import (
	"context"
	"log"

	"github.com/aussiebroadwan/tipping/backend/internal/models"
	"github.com/aussiebroadwan/tipping/backend/internal/services"
)

func seedDatabase() {

	// Due to migration files the competitions table should already be populated
	// NOP

	// Add Some Fixtures
	if err := addCowboysVsStormUpcomingFixture(); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	if err := addBulldogsVsSeaEaglesUpcomingFixture(); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	if err := addTitansVsSharksFixture(); err != nil { // This is NRLW
		log.Fatalf("Failed to seed database: %v", err)
	}
}

func addCowboysVsStormUpcomingFixture() error {
	oddsHome := "1.23"
	oddsAway := "4.25"

	fixture := models.NRLFixture{
		ID:             "20241112610",
		RoundTitle:     "Round 26",
		MatchState:     "Upcoming",
		KickOffTime:    "2024-08-27T01:16:09Z",
		Venue:          "Queensland Country Bank Stadium",
		VenueCity:      "Townsville",
		MatchCentreURL: "/draw/nrl-premiership/2024/round-26/cowboys-v-storm/",
		HomeTeam: models.NRLTeam{
			ID:    500012,
			Name:  "Cowboys",
			Odds:  &oddsHome,
			Score: nil,
			Form: []models.NRLForm{
				{
					Result: "Won",
					Score:  "42-4",
				},
				{
					Result: "Lost",
					Score:  "18-42",
				},
				{
					Result: "Won",
					Score:  "38-30",
				},
				{
					Result: "Won",
					Score:  "30-22",
				},
				{
					Result: "Won",
					Score:  "20-18",
				},
			},
		},
		AwayTeam: models.NRLTeam{
			ID:    500021,
			Name:  "Storm",
			Odds:  &oddsAway,
			Score: nil,
			Form: []models.NRLForm{
				{
					Result: "Won",
					Score:  "48-6",
				},
				{
					Result: "Won",
					Score:  "24-22",
				},
				{
					Result: "Won",
					Score:  "28-16",
				},
				{
					Result: "Lost",
					Score:  "16-18",
				},
				{
					Result: "Won",
					Score:  "32-14",
				},
			},
		},
	}

	dataService := services.NewNRLDataService(testQueries, context.Background())
	return dataService.StoreFixtureAndDetails(fixture)
}

func addBulldogsVsSeaEaglesUpcomingFixture() error {
	oddsHome := "1.63"
	oddsAway := "2.30"

	fixture := models.NRLFixture{
		ID:             "20241112620",
		RoundTitle:     "Round 26",
		MatchState:     "Upcoming",
		KickOffTime:    "2024-08-30T08:00:00Z",
		Venue:          "Accor Stadium",
		VenueCity:      "Sydney",
		MatchCentreURL: "/draw/nrl-premiership/2024/round-26/bulldogs-v-sea-eagles/",
		HomeTeam: models.NRLTeam{
			ID:    500010,
			Name:  "Bulldogs",
			Odds:  &oddsHome,
			Score: nil,
			Form: []models.NRLForm{
				{
					Result: "Won",
					Score:  "34-18",
				},
				{
					Result: "Won",
					Score:  "30-10",
				},
				{
					Result: "Won",
					Score:  "28-10",
				},
				{
					Result: "Won",
					Score:  "22-18",
				},
				{
					Result: "Won",
					Score:  "41-16",
				},
			},
		},
		AwayTeam: models.NRLTeam{
			ID:    500002,
			Name:  "Sea Eagles",
			Odds:  &oddsAway,
			Score: nil,
			Form: []models.NRLForm{
				{
					Result: "Lost",
					Score:  "26-34",
				},
				{
					Result: "Won",
					Score:  "24-10",
				},
				{
					Result: "Won",
					Score:  "46-24",
				},
				{
					Result: "Lost",
					Score:  "30-34",
				},
				{
					Result: "Won",
					Score:  "38-8",
				},
			},
		},
	}

	dataService := services.NewNRLDataService(testQueries, context.Background())
	return dataService.StoreFixtureAndDetails(fixture)
}

func addTitansVsSharksFixture() error {
	oddsHome := "1.71"
	oddsAway := "2.15"

	fixture := models.NRLFixture{
		ID:             "20241610610",
		RoundTitle:     "Round 6",
		MatchState:     "Upcoming",
		KickOffTime:    "2024-08-27T07:16:05Z",
		Venue:          "Cbus Super Stadium",
		VenueCity:      "Gold Coast",
		MatchCentreURL: "/draw/womens-premiership/2024/round-6/titans-v-sharks/",
		HomeTeam: models.NRLTeam{
			ID:    500690,
			Name:  "Titans",
			Odds:  &oddsHome,
			Score: nil,
			Form: []models.NRLForm{
				{
					Result: "Won",
					Score:  "26-6",
				},
				{
					Result: "Lost",
					Score:  "10-11",
				},
				{
					Result: "Lost",
					Score:  "4-44",
				},
				{
					Result: "Won",
					Score:  "28-12",
				},
				{
					Result: "Won",
					Score:  "18-10",
				},
			},
		},
		AwayTeam: models.NRLTeam{
			ID:    500786,
			Name:  "Sharks",
			Odds:  &oddsAway,
			Score: nil,
			Form: []models.NRLForm{
				{
					Result: "Won",
					Score:  "28-4",
				},
				{
					Result: "Won",
					Score:  "14-12",
				},
				{
					Result: "Won",
					Score:  "24-12",
				},
				{
					Result: "Won",
					Score:  "18-16",
				},
				{
					Result: "Won",
					Score:  "14-0",
				},
			},
		},
	}

	dataService := services.NewNRLDataService(testQueries, context.Background())
	return dataService.StoreFixtureAndDetails(fixture)
}
