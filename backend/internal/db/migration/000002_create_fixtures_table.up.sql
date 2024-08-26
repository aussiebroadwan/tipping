CREATE TABLE fixtures (
  id BIGINT PRIMARY KEY,
  competition_id BIGINT NOT NULL REFERENCES competitions(id) ON DELETE CASCADE,
  roundTitle VARCHAR(255) NOT NULL,
  matchState VARCHAR(50) NOT NULL,
  venue VARCHAR(255) NOT NULL,
  venueCity VARCHAR(255) NOT NULL,
  matchCentreUrl VARCHAR(255) NOT NULL,
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
