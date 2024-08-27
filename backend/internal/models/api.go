package models

import "time"

// APICompetition represents a competition in the API response.
type APICompetition struct {
	ID   int64  `json:"id" example:"111"`   // Unique identifier for the competition
	Name string `json:"name" example:"NRL"` // Name of the competition
}

// APIFixture represents a fixture in the API response.
type APIFixture struct {
	ID            int64     `json:"id" example:"20241610510"`                     // Unique identifier for the fixture
	CompetitionID int64     `json:"competition_id" example:"111"`                 // The competition ID this fixture belongs to
	RoundTitle    string    `json:"round_title" example:"Round 22"`               // The title of the round
	MatchState    string    `json:"match_state" example:"FullTime"`               // Current state of the match
	Venue         string    `json:"venue" example:"Leichhardt Oval"`              // Venue of the match
	VenueCity     string    `json:"venue_city" example:"Sydney"`                  // City where the venue is located
	KickOffTime   time.Time `json:"kick_off_time" example:"2024-08-24T01:00:00Z"` // Kickoff time of the match in RFC3339 format
	HomeTeam      APITeam   `json:"home_team"`                                    // Home team details
	AwayTeam      APITeam   `json:"away_team"`                                    // Away team details
}

// APITeam represents a team in the API response.
type APITeam struct {
	Nickname string   `json:"nickname" example:"Cowboys"`    // Nickname of the team
	Odds     *float64 `json:"odds,omitempty" example:"3.42"` // Odds for the team to win
	Score    *int32   `json:"score,omitempty" example:"40"`  // Final score of the team
	Form     string   `json:"form" example:"WLWWL"`          // Recent form of the team
}
