-- name: GetMatchDetailsByFixtureID :one
-- Retrieve match details for a specific fixture by its unique fixture ID.
SELECT * FROM match_details WHERE fixture_id = $1;

-- name: ListMatchDetails :many
-- Retrieve all match details available in the system.
SELECT * FROM match_details;

-- name: ListMatchDetailsByCompetitionID :many
-- Retrieve all match details for a specific competition ID.
-- This query performs a JOIN between match_details and fixtures to get all 
-- match details that are part of a specific competition.
SELECT md.*
FROM match_details md
JOIN fixtures f ON md.fixture_id = f.id
WHERE f.competition_id = $1
ORDER BY f.kickOffTime;

-- name: CreateMatchDetail :one
-- Insert a new match detail record into the match_details table.
-- If a match detail with the same fixture_id already exists, do nothing.
INSERT INTO match_details (
  fixture_id, homeTeam_id, awayTeam_id, homeTeam_odds, awayTeam_odds, 
  homeTeam_score, awayTeam_score, homeTeam_form, awayTeam_form, winner_teamId
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateMatchDetail :one
-- Conditionally update match detail fields based on provided arguments.
-- Only updates fields where the argument is not NULL.
UPDATE match_details 
SET 
    homeTeam_odds = COALESCE(sqlc.narg('homeTeam_odds'), homeTeam_odds), 
    awayTeam_odds = COALESCE(sqlc.narg('awayTeam_odds'), awayTeam_odds), 
    homeTeam_score = COALESCE(sqlc.narg('homeTeam_score'), homeTeam_score), 
    awayTeam_score = COALESCE(sqlc.narg('awayTeam_score'), awayTeam_score), 
    winner_teamId = COALESCE(sqlc.narg('winner_teamId'), winner_teamId)
WHERE fixture_id = $1
RETURNING *;