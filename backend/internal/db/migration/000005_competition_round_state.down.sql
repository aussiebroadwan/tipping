-- Down migration to remove the 'round' column from the competitions table
ALTER TABLE competitions
DROP COLUMN round;