// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Competition struct {
	// Unique identifier for each competition
	ID int64
	// Name of the competition (e.g., NRL, NRLW)
	Name string
	// Indicates the current round or game for the competition (e.g., Round 1 for NRL, Game 1 for State of Origin)
	Round *string
}

type Fixture struct {
	// Unique identifier for each fixture
	ID int64
	// Foreign key referencing competitions table
	CompetitionID int64
	// Title of the round (e.g., Round 1)
	Roundtitle string
	// Current state of the match (e.g., Upcoming, Completed)
	Matchstate string
	// Venue name where the match will take place
	Venue string
	// City where the venue is located
	Venuecity string
	// URL to the match center page
	Matchcentreurl string
	// Scheduled kickoff time of the match
	Kickofftime pgtype.Timestamp
}

type MatchDetail struct {
	// Foreign key referencing fixtures table
	FixtureID int64
	// Foreign key for home team referencing teams table
	HometeamID int64
	// Foreign key for away team referencing teams table
	AwayteamID int64
	// Odds for the home team winning
	HometeamOdds *float64
	// Odds for the away team winning
	AwayteamOdds *float64
	// Score of the home team
	HometeamScore *int32
	// Score of the away team
	AwayteamScore *int32
	// Recent form of the home team (e.g., WLWWL)
	HometeamForm string
	// Recent form of the away team (e.g., LWWLL)
	AwayteamForm string
	// Foreign key referencing the winning team
	WinnerTeamid *int64
}

type Team struct {
	// Unique identifier for each team
	ID int64
	// Nickname or short name for the team (e.g., Cowboys)
	Nickname      string
	CompetitionID int64
}
