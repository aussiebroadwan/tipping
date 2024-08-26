// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: teams.sql

package sqlc

import (
	"context"
)

const createTeam = `-- name: CreateTeam :one
INSERT INTO teams (team_id, nickName) 
VALUES ($1, $2)
ON CONFLICT (team_id) DO NOTHING
RETURNING team_id, nickname
`

type CreateTeamParams struct {
	TeamID   int32
	Nickname string
}

// Insert a new team into the teams table.
// If a team with the same team_id already exists, do nothing.
func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (*Team, error) {
	row := q.db.QueryRow(ctx, createTeam, arg.TeamID, arg.Nickname)
	var i Team
	err := row.Scan(&i.TeamID, &i.Nickname)
	return &i, err
}

const getTeamByID = `-- name: GetTeamByID :one
SELECT team_id, nickname FROM teams WHERE team_id = $1
`

// Retrieve a specific team by its unique identifier.
func (q *Queries) GetTeamByID(ctx context.Context, teamID int32) (*Team, error) {
	row := q.db.QueryRow(ctx, getTeamByID, teamID)
	var i Team
	err := row.Scan(&i.TeamID, &i.Nickname)
	return &i, err
}

const listTeams = `-- name: ListTeams :many
SELECT team_id, nickname FROM teams
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
		if err := rows.Scan(&i.TeamID, &i.Nickname); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
