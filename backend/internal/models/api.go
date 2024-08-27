package models

import "time"

// Competition represents a competition in the API response.
type Competition struct {
	ID   int64  `json:"id"`   // Unique identifier for the competition
	Name string `json:"name"` // Name of the competition (e.g., "NRL", "NRLW")
}

// Fixture represents a fixture in the API response.
type Fixture struct {
	ID             int64     `json:"id"`               // Unique identifier for the fixture
	CompetitionID  int64     `json:"competition_id"`   // The competition ID this fixture belongs to
	RoundTitle     string    `json:"round_title"`      // The title of the round (e.g., "Round 22")
	MatchState     string    `json:"match_state"`      // Current state of the match (e.g., "Upcoming", "FullTime")
	Venue          string    `json:"venue"`            // Venue of the match (e.g., "Leichhardt Oval")
	VenueCity      string    `json:"venue_city"`       // City where the venue is located (e.g., "Sydney")
	MatchCentreURL string    `json:"match_centre_url"` // URL to the match centre details
	KickOffTime    time.Time `json:"kick_off_time"`    // Kickoff time of the match in RFC3339 format
}

// Team represents a team in the API response.
type Team struct {
	ID       int64  `json:"id"`       // Unique identifier for the team
	Nickname string `json:"nickname"` // Nickname of the team (e.g., "Cowboys")
}

// MatchDetail represents the detailed information of a match in the API response.
type MatchDetail struct {
	FixtureID     int64    `json:"fixture_id"`      // Unique identifier for the fixture (match)
	HomeTeamID    int64    `json:"home_team_id"`    // ID of the home team
	AwayTeamID    int64    `json:"away_team_id"`    // ID of the away team
	HomeTeamOdds  *float64 `json:"home_team_odds"`  // Odds for the home team to win
	AwayTeamOdds  *float64 `json:"away_team_odds"`  // Odds for the away team to win
	HomeTeamScore *int32   `json:"home_team_score"` // Final score of the home team
	AwayTeamScore *int32   `json:"away_team_score"` // Final score of the away team
	HomeTeamForm  string   `json:"home_team_form"`  // Recent form of the home team (e.g., "WLWWL")
	AwayTeamForm  string   `json:"away_team_form"`  // Recent form of the away team (e.g., "LLLWW")
	WinnerTeamID  *int64   `json:"winner_team_id"`  // ID of the winning team, if available
}

// ErrorResponse represents the structure of an error response in the API.
type ErrorResponse struct {
	Message string `json:"message"` // Error message describing what went wrong
}
