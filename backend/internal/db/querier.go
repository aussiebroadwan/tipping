// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	// Insert a new fixture into the fixtures table.
	// This query adds a new fixture record with the specified details, such as
	// competition ID, round title, match state, venue, venue city, match center URL,
	// and kickoff time.
	CreateFixture(ctx context.Context, arg CreateFixtureParams) (*Fixture, error)
	// Insert a new match detail record into the match_details table.
	// If a match detail with the same fixture_id already exists, do nothing.
	CreateMatchDetail(ctx context.Context, arg CreateMatchDetailParams) (*MatchDetail, error)
	// Insert a new team into the teams table.
	// If a team with the same team_id already exists, do nothing.
	CreateTeam(ctx context.Context, arg CreateTeamParams) (*Team, error)
	// Retrieve a specific competition by its unique identifier.
	GetCompetitionByID(ctx context.Context, id int64) (*Competition, error)
	// Retrieve a specific fixture by its unique identifier.
	// Useful for fetching details about a single fixture based on its ID.
	GetFixtureByID(ctx context.Context, id int64) (*Fixture, error)
	// Retrieve fixtures for a specific competition, ordered by kickoff time.
	// This query fetches all fixtures for a given competition ID, ordered by their
	// kickoff time to display them in chronological order.
	GetFixturesByCompetitionID(ctx context.Context, competitionID int64) ([]*Fixture, error)
	// Retrieve match details for a specific fixture by its unique fixture ID.
	GetMatchDetailsByFixtureID(ctx context.Context, fixtureID int64) (*GetMatchDetailsByFixtureIDRow, error)
	// Retrieve a specific team by its unique identifier.
	GetTeamByID(ctx context.Context, teamID int64) (*Team, error)
	// The competitions table is a static table that stores information about the
	// competitions that are available in the system. Other tables in the system
	// reference this table to establish a relationship.
	// Retrieve all competitions available in the system.
	ListCompetitions(ctx context.Context) ([]*Competition, error)
	// Retrieve all fixtures available in the system.
	// This query is used to list all fixtures without filtering by any criteria.
	ListFixtures(ctx context.Context) ([]*Fixture, error)
	// Retrieve all match details available in the system.
	ListMatchDetails(ctx context.Context) ([]*ListMatchDetailsRow, error)
	// Retrieve all match details for a specific competition ID.
	// This query performs a JOIN between match_details and fixtures to get all
	// match details that are part of a specific competition.
	ListMatchDetailsByCompetitionID(ctx context.Context, competitionID int64) ([]*ListMatchDetailsByCompetitionIDRow, error)
	// Retrieve all teams available in the system.
	ListTeams(ctx context.Context) ([]*Team, error)
	// Conditionally update fixture details based on provided arguments.
	// This query updates the fields of a fixture record where the provided arguments
	// are not NULL. It uses the COALESCE function to retain the existing value if
	// the argument is NULL.
	UpdateFixture(ctx context.Context, arg UpdateFixtureParams) (*Fixture, error)
	// Conditionally update match detail fields based on provided arguments.
	// Only updates fields where the argument is not NULL.
	UpdateMatchDetail(ctx context.Context, arg UpdateMatchDetailParams) (*MatchDetail, error)
}

var _ Querier = (*Queries)(nil)
