ALTER TABLE competitions
ADD COLUMN round VARCHAR(50);

COMMENT ON COLUMN competitions.round IS 'Indicates the current round or game for the competition (e.g., Round 1 for NRL, Game 1 for State of Origin)';
