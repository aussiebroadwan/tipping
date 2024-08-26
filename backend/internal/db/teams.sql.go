// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: teams.sql

package db

import (
	"context"
)

const createTeam = `-- name: CreateTeam :one
INSERT INTO teams (team_id, nickName, competition_id) 
VALUES ($1, $2, $3)
RETURNING team_id, nickname, competition_id
`

type CreateTeamParams struct {
	TeamID        int64
	Nickname      string
	CompetitionID int64
}

// Insert a new team into the teams table.
// If a team with the same team_id already exists, do nothing.
func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (*Team, error) {
	row := q.db.QueryRow(ctx, createTeam, arg.TeamID, arg.Nickname, arg.CompetitionID)
	var i Team
	err := row.Scan(&i.TeamID, &i.Nickname, &i.CompetitionID)
	return &i, err
}

const getTeamByID = `-- name: GetTeamByID :one
SELECT team_id, nickname, competition_id FROM teams WHERE team_id = $1
`

// Retrieve a specific team by its unique identifier.
func (q *Queries) GetTeamByID(ctx context.Context, teamID int64) (*Team, error) {
	row := q.db.QueryRow(ctx, getTeamByID, teamID)
	var i Team
	err := row.Scan(&i.TeamID, &i.Nickname, &i.CompetitionID)
	return &i, err
}

const listTeams = `-- name: ListTeams :many
SELECT team_id, nickname, competition_id FROM teams
`

// Retrieve all teams available in the system.
func (q *Queries) ListTeams(ctx context.Context) ([]*Team, error) {
	rows, err := q.db.Query(ctx, listTeams)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(&i.TeamID, &i.Nickname, &i.CompetitionID); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
