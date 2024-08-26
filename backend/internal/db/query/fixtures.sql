-- name: ListFixtures :many
-- Retrieve all fixtures available in the system.
-- This query is used to list all fixtures without filtering by any criteria.
SELECT * FROM fixtures;


-- name: GetFixtureByID :one
-- Retrieve a specific fixture by its unique identifier.
-- Useful for fetching details about a single fixture based on its ID.
SELECT * FROM fixtures WHERE id = $1;

-- name: GetFixturesByCompetitionID :many
-- Retrieve fixtures for a specific competition, ordered by kickoff time.
-- This query fetches all fixtures for a given competition ID, ordered by their
-- kickoff time to display them in chronological order.
SELECT * FROM fixtures 
WHERE competition_id = $1
ORDER BY kickOffTime;

-- name: CreateFixture :one
-- Insert a new fixture into the fixtures table.
-- This query adds a new fixture record with the specified details, such as
-- competition ID, round title, match state, venue, venue city, match center URL,
-- and kickoff time.
INSERT INTO fixtures (
  competition_id, roundTitle, matchState, venue, venueCity, matchCentreUrl, kickOffTime
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateFixture :one
-- Conditionally update fixture details based on provided arguments.
-- This query updates the fields of a fixture record where the provided arguments
-- are not NULL. It uses the COALESCE function to retain the existing value if 
-- the argument is NULL.
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
