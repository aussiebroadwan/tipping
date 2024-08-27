package services

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aussiebroadwan/tipping/backend/config"
	"github.com/aussiebroadwan/tipping/backend/internal/db"
	"github.com/aussiebroadwan/tipping/backend/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
)

// NRLDataService defines a service for handling data conversion and integration with the database.
type NRLDataService struct {
	queries *db.Queries
	ctx     context.Context
}

// NewNRLDataService creates a new instance of NRLDataService.
func NewNRLDataService(queries *db.Queries, ctx context.Context) *NRLDataService {
	return &NRLDataService{
		queries: queries,
		ctx:     ctx,
	}
}

// StoreFixtureAndDetails converts NRLFixture to database models and stores them.
func (s *NRLDataService) StoreFixtureAndDetails(fixture models.NRLFixture) error {
	// Parse fixture ID
	fixtureID, err := strconv.ParseInt(fixture.ID, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse fixture ID: %w", err)
	}

	// Parse match ID components
	_, compID, _, _ := parseMatchID(fixture.ID)

	// Parse kickoff time
	kickOffTime, err := time.Parse(time.RFC3339, fixture.KickOffTime)
	if err != nil {
		return fmt.Errorf("failed to parse fixture kickoff time: %w", err)
	}

	// Create fixture in the database
	err = s.createOrUpdateFixture(fixtureID, compID, fixture, kickOffTime)
	if err != nil {
		return err
	}

	// Store each team
	if err := s.storeTeam(fixture.HomeTeam, compID); err != nil {
		return fmt.Errorf("failed to store home team: %w", err)
	}
	if err := s.storeTeam(fixture.AwayTeam, compID); err != nil {
		return fmt.Errorf("failed to store away team: %w", err)
	}

	// Store match details
	if err := s.storeMatchDetails(fixtureID, fixture); err != nil {
		return fmt.Errorf("failed to store match details: %w", err)
	}

	return nil
}

func (s *NRLDataService) UpdateMatchState(fixtureID string, matchState string) error {
	// Parse fixture ID
	id, err := strconv.ParseInt(fixtureID, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse fixture ID: %w", err)
	}

	_, err = s.queries.UpdateFixture(s.ctx, db.UpdateFixtureParams{
		ID:         id,
		MatchState: &matchState,
	})
	if err != nil {
		return fmt.Errorf("failed to update fixture: %w", err)
	}
	return nil
}

// createOrUpdateFixture creates or updates a fixture in the database.
func (s *NRLDataService) createOrUpdateFixture(fixtureID int64, compID int, fixture models.NRLFixture, kickOffTime time.Time) error {
	pgxKickOffTime := pgtype.Timestamp{Time: kickOffTime, Valid: true}

	// Check if fixture exists
	if _, err := s.queries.GetFixtureByID(s.ctx, fixtureID); err == nil {
		// Update fixture
		_, err := s.queries.UpdateFixture(s.ctx, db.UpdateFixtureParams{
			ID:         fixtureID,
			MatchState: &fixture.MatchState,
		})
		if err != nil {
			return fmt.Errorf("failed to update fixture: %w", err)
		}
		return nil
	}

	_, err := s.queries.CreateFixture(s.ctx, db.CreateFixtureParams{
		ID:             fixtureID,
		CompetitionID:  int64(compID),
		Roundtitle:     fixture.RoundTitle,
		Matchstate:     fixture.MatchState,
		Venue:          fixture.Venue,
		Venuecity:      fixture.VenueCity,
		Matchcentreurl: fixture.MatchCentreURL,
		Kickofftime:    pgxKickOffTime,
	})
	if err != nil {
		return fmt.Errorf("failed to store fixture: %w", err)
	}

	return nil
}

// storeTeam stores a team in the database, creating it if it does not exist.
func (s *NRLDataService) storeTeam(team models.NRLTeam, competitionId int) error {
	// Check if team exists
	_, err := s.queries.GetTeamByID(s.ctx, int64(team.ID))
	if err == nil {
		return nil
	}

	// Create team
	_, err = s.queries.CreateTeam(s.ctx, db.CreateTeamParams{
		TeamID:        int64(team.ID),
		Nickname:      team.Name,
		CompetitionID: int64(competitionId),
	})
	if err != nil {
		return fmt.Errorf("failed to store team: %w", err)
	}
	return nil
}

// storeMatchDetails converts and stores match details in the database.
func (s *NRLDataService) storeMatchDetails(fixtureID int64, fixture models.NRLFixture) error {
	// Check if match details exist
	_, err := s.queries.GetMatchDetailsByFixtureID(s.ctx, fixtureID)
	if err == nil {
		// Match details already exist Update them
		_, err := s.queries.UpdateMatchDetail(s.ctx, db.UpdateMatchDetailParams{
			FixtureID:     fixtureID,
			HomeTeamScore: parseScore(fixture.HomeTeam.Score),
			AwayTeamScore: parseScore(fixture.AwayTeam.Score),
			HomeTeamOdds:  parseOdds(fixture.HomeTeam.Odds),
			AwayTeamOdds:  parseOdds(fixture.AwayTeam.Odds),
			WinnerTeamId:  parseWinnerTeamID(fixture),
		})
		if err != nil {
			return fmt.Errorf("failed to update match details: %w", err)
		}
		return nil
	}

	_, err = s.queries.CreateMatchDetail(s.ctx, db.CreateMatchDetailParams{
		FixtureID:     fixtureID,
		HometeamID:    int64(fixture.HomeTeam.ID),
		AwayteamID:    int64(fixture.AwayTeam.ID),
		HometeamOdds:  parseOdds(fixture.HomeTeam.Odds),
		AwayteamOdds:  parseOdds(fixture.AwayTeam.Odds),
		HometeamScore: parseScore(fixture.HomeTeam.Score),
		AwayteamScore: parseScore(fixture.AwayTeam.Score),
		HometeamForm:  parseForm(fixture.HomeTeam.Form),
		AwayteamForm:  parseForm(fixture.AwayTeam.Form),
		WinnerTeamid:  parseWinnerTeamID(fixture),
	})
	if err != nil {
		return fmt.Errorf("failed to store match details: %w", err)
	}
	return nil
}

// Helper functions to parse different data types.

func parseOdds(odds *string) *float64 {
	if odds == nil {
		return nil
	}
	oddsValue, _ := strconv.ParseFloat(*odds, 64)
	return &oddsValue
}

func parseScore(score *int) *int32 {
	if score == nil {
		return nil
	}
	scoreValue := int32(*score)
	return &scoreValue
}

func parseForm(form []models.NRLForm) string {
	if len(form) == 0 {
		return ""
	}
	var result string
	for _, f := range form {
		if f.Result == "Won" {
			result += "W"
		} else {
			result += "L"
		}
	}
	return result
}

func parseWinnerTeamID(fixture models.NRLFixture) *int64 {

	if fixture.MatchState != config.MatchStateFullTime {
		return nil
	}

	if fixture.HomeTeam.Score != nil && fixture.AwayTeam.Score != nil {
		if *fixture.HomeTeam.Score > *fixture.AwayTeam.Score {
			homeTeamID := int64(fixture.HomeTeam.ID)
			return &homeTeamID
		} else if *fixture.HomeTeam.Score < *fixture.AwayTeam.Score {
			awayTeamID := int64(fixture.AwayTeam.ID)
			return &awayTeamID
		}
	}
	return nil
}

func parseMatchID(id string) (season, competition, round, game int) {
	season, _ = strconv.Atoi(id[0:4])
	competition, _ = strconv.Atoi(id[4:7])
	round, _ = strconv.Atoi(id[7:9])
	game, _ = strconv.Atoi(string(id[9]))

	return
}
