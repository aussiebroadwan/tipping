CREATE TABLE match_details (
  fixture_id BIGINT PRIMARY KEY REFERENCES fixtures(id) ON DELETE CASCADE,
  homeTeam_id BIGINT NOT NULL REFERENCES teams(id),
  awayTeam_id BIGINT NOT NULL REFERENCES teams(id),
  homeTeam_odds FLOAT,
  awayTeam_odds FLOAT,
  homeTeam_score INTEGER,
  awayTeam_score INTEGER,
  homeTeam_form VARCHAR(5) NOT NULL,
  awayTeam_form VARCHAR(5) NOT NULL,
  winner_teamId BIGINT REFERENCES teams(id)
);

COMMENT ON COLUMN match_details.fixture_id IS 'Foreign key referencing fixtures table';
COMMENT ON COLUMN match_details.homeTeam_id IS 'Foreign key for home team referencing teams table';
COMMENT ON COLUMN match_details.awayTeam_id IS 'Foreign key for away team referencing teams table';
COMMENT ON COLUMN match_details.homeTeam_odds IS 'Odds for the home team winning';
COMMENT ON COLUMN match_details.awayTeam_odds IS 'Odds for the away team winning';
COMMENT ON COLUMN match_details.homeTeam_score IS 'Score of the home team';
COMMENT ON COLUMN match_details.awayTeam_score IS 'Score of the away team';
COMMENT ON COLUMN match_details.homeTeam_form IS 'Recent form of the home team (e.g., WLWWL)';
COMMENT ON COLUMN match_details.awayTeam_form IS 'Recent form of the away team (e.g., LWWLL)';
COMMENT ON COLUMN match_details.winner_teamId IS 'Foreign key referencing the winning team';
