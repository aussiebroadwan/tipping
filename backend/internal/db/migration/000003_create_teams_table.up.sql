CREATE TABLE teams (
    team_id BIGINT PRIMARY KEY,
    nickName VARCHAR(255) NOT NULL,
    competition_id BIGINT NOT NULL REFERENCES competitions(id) ON DELETE CASCADE
);

COMMENT ON COLUMN teams.team_id IS 'Unique identifier for each team';
COMMENT ON COLUMN teams.nickName IS 'Nickname or short name for the team (e.g., Cowboys)';
