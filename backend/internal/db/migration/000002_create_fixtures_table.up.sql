CREATE TABLE fixtures (
  id SERIAL PRIMARY KEY,
  competition_id INTEGER NOT NULL REFERENCES competitions(id) ON DELETE CASCADE,
  roundTitle VARCHAR(255),
  matchState VARCHAR(50),
  venue VARCHAR(255),
  venueCity VARCHAR(255),
  matchCentreUrl VARCHAR(255),
  kickOffTime TIMESTAMP WITHOUT TIME ZONE
);

COMMENT ON COLUMN fixtures.id IS 'Unique identifier for each fixture';
COMMENT ON COLUMN fixtures.competition_id IS 'Foreign key referencing competitions table';
COMMENT ON COLUMN fixtures.roundTitle IS 'Title of the round (e.g., Round 1)';
COMMENT ON COLUMN fixtures.matchState IS 'Current state of the match (e.g., Upcoming, Completed)';
COMMENT ON COLUMN fixtures.venue IS 'Venue name where the match will take place';
COMMENT ON COLUMN fixtures.venueCity IS 'City where the venue is located';
COMMENT ON COLUMN fixtures.matchCentreUrl IS 'URL to the match center page';
COMMENT ON COLUMN fixtures.kickOffTime IS 'Scheduled kickoff time of the match';
