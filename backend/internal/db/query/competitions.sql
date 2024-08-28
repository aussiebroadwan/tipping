-- The competitions table is a static table that stores information about the 
-- competitions that are available in the system. Other tables in the system
-- reference this table to establish a relationship. 

-- name: ListCompetitions :many
-- Retrieve all competitions available in the system.
SELECT * FROM competitions;

-- name: GetCompetitionByID :one
-- Retrieve a specific competition by its unique identifier.
SELECT * FROM competitions WHERE id = $1;

-- The following commands for creating, updating, and deleting competitions
-- are not required since this is a static table with fixed records:
-- - NRL (111)
-- - NRLW (161)
-- - State of Origin (116)
-- - State of Origin Womens (156)
--
-- However, if future updates to this table are needed (e.g., new competitions),
-- you may add additional commands to handle such changes.


-- name: UpdateCompetitionRound :one
-- Update the current round for a competition.
-- This query updates the round field for a specific competition based on the
-- provided competition ID.
UPDATE competitions
SET round = $2
WHERE id = $1
RETURNING *;