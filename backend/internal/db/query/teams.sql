-- name: ListTeams :many
-- Retrieve all teams available in the system.
SELECT * FROM teams;

-- name: GetTeamByID :one
-- Retrieve a specific team by its unique identifier.
SELECT * FROM teams WHERE id = $1;

-- name: CreateTeam :one
-- Insert a new team into the teams table.
-- If a team with the same id already exists, do nothing.
INSERT INTO teams (id, nickName, competition_id) 
VALUES ($1, $2, $3)
RETURNING *;

