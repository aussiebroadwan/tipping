-- name: ListFixtures :many
SELECT * FROM fixtures;

-- name: GetFixtureByID :one
SELECT * FROM fixtures WHERE id = $1;

-- name: GetFixturesByCompetitionID :many
SELECT * FROM fixtures 
WHERE competition_id = $1
ORDER BY kickOffTime;

-- name: CreateFixture :one
INSERT INTO fixtures (competition_id, roundTitle, matchState, venue, venueCity, matchCentreUrl, kickOffTime) 
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateFixture :one
-- Conditionally update fixture details based on provided arguments.
UPDATE fixtures 
SET 
    competition_id = COALESCE($2, competition_id), 
    roundTitle = COALESCE($3, roundTitle), 
    matchState = COALESCE($4, matchState), 
    venue = COALESCE($5, venue), 
    venueCity = COALESCE($6, venueCity), 
    matchCentreUrl = COALESCE($7, matchCentreUrl), 
    kickOffTime = COALESCE($8, kickOffTime)
WHERE id = $1
RETURNING *;
-- name: DeleteFixture :exec
DELETE FROM fixtures WHERE id = $1;