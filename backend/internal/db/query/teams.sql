-- name: ListTeams :many
-- Retrieve all teams available in the system.
SELECT * FROM teams;

-- name: GetTeamByID :one
-- Retrieve a specific team by its unique identifier.
SELECT * FROM teams WHERE team_id = $1;

-- name: CreateTeam :one
-- Insert a new team into the teams table.
-- If a team with the same team_id already exists, do nothing.
INSERT INTO teams (team_id, nickName) 
VALUES ($1, $2)
RETURNING *;

