package services

import (
	"context"
	"fmt"

	"github.com/aussiebroadwan/tipping/backend/internal/db"
	"github.com/aussiebroadwan/tipping/backend/internal/models"
)

// APIDataService defines a service for handling data conversion and integration with the database.
type APIDataService struct {
	queries *db.Queries
	ctx     context.Context
}

// NewAPIDataService creates a new instance of APIDataService.
func NewAPIDataService(queries *db.Queries, ctx context.Context) *APIDataService {
	return &APIDataService{
		queries: queries,
		ctx:     ctx,
	}
}

// GetCompetitions fetches all competitions from the database and returns them
// as a list of APICompetition models.
func (s *APIDataService) GetCompetitions() ([]models.APICompetition, error) {
	competitions, err := s.queries.ListCompetitions(s.ctx)
	if err != nil {
		return nil, err
	}

	// Convert to API Model
	comps := make([]models.APICompetition, 0)
	for _, c := range competitions {
		comps = append(comps, models.APICompetition{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return comps, nil
}

// GetFixtures fetches all fixtures with match details from the database and converts them to API models.
func (s *APIDataService) GetFixtures() ([]models.APIFixture, error) {
	fixtures, err := s.queries.ListMatchDetails(s.ctx)
	if err != nil {
		return nil, err
	}

	// Convert database models to API models.
	apiFixtures := make([]models.APIFixture, 0)
	for _, f := range fixtures {
		apiFixtures = append(apiFixtures, models.APIFixture{
			ID:            f.Fixture.ID,
			CompetitionID: f.Fixture.CompetitionID,
			RoundTitle:    f.Fixture.Roundtitle,
			MatchState:    f.Fixture.Matchstate,
			Venue:         f.Fixture.Venue,
			VenueCity:     f.Fixture.Venuecity,
			HomeTeam: models.APITeam{
				Nickname: f.Team.Nickname,
				Score:    f.MatchDetail.HometeamScore,
				Odds:     f.MatchDetail.HometeamOdds,
				Form:     f.MatchDetail.HometeamForm,
			},
			AwayTeam: models.APITeam{
				Nickname: f.Team_2.Nickname,
				Score:    f.MatchDetail.AwayteamScore,
				Odds:     f.MatchDetail.AwayteamOdds,
				Form:     f.MatchDetail.AwayteamForm,
			},
			KickOffTime: f.Fixture.Kickofftime.Time,
		})
	}

	return apiFixtures, nil
}

// GetFixtures fetches all fixtures with match details from the database with a given round and converts them to API models.
func (s *APIDataService) GetRoundFixtures(round int) ([]models.APIFixture, error) {
	fixtures, err := s.queries.ListRoundMatchDetails(s.ctx, fmt.Sprintf("Round %d", round))
	if err != nil {
		return nil, err
	}

	// Convert database models to API models.
	apiFixtures := make([]models.APIFixture, 0)
	for _, f := range fixtures {
		apiFixtures = append(apiFixtures, models.APIFixture{
			ID:            f.Fixture.ID,
			CompetitionID: f.Fixture.CompetitionID,
			RoundTitle:    f.Fixture.Roundtitle,
			MatchState:    f.Fixture.Matchstate,
			Venue:         f.Fixture.Venue,
			VenueCity:     f.Fixture.Venuecity,
			HomeTeam: models.APITeam{
				Nickname: f.Team.Nickname,
				Score:    f.MatchDetail.HometeamScore,
				Odds:     f.MatchDetail.HometeamOdds,
				Form:     f.MatchDetail.HometeamForm,
			},
			AwayTeam: models.APITeam{
				Nickname: f.Team_2.Nickname,
				Score:    f.MatchDetail.AwayteamScore,
				Odds:     f.MatchDetail.AwayteamOdds,
				Form:     f.MatchDetail.AwayteamForm,
			},
			KickOffTime: f.Fixture.Kickofftime.Time,
		})
	}

	return apiFixtures, nil
}

// GetCompetitionFixtures fetches fixtures for a specific competition and converts them to API models.
func (s *APIDataService) GetCompetitionFixtures(competitionId int64) ([]models.APIFixture, error) {
	fixtures, err := s.queries.ListMatchDetailsByCompetitionID(s.ctx, competitionId)
	if err != nil {
		return nil, err
	}

	// Convert database models to API models.
	apiFixtures := make([]models.APIFixture, 0)
	for _, f := range fixtures {
		apiFixtures = append(apiFixtures, models.APIFixture{
			ID:            f.Fixture.ID,
			CompetitionID: f.Fixture.CompetitionID,
			RoundTitle:    f.Fixture.Roundtitle,
			MatchState:    f.Fixture.Matchstate,
			Venue:         f.Fixture.Venue,
			VenueCity:     f.Fixture.Venuecity,
			HomeTeam: models.APITeam{
				Nickname: f.Team.Nickname,
				Score:    f.MatchDetail.HometeamScore,
				Odds:     f.MatchDetail.HometeamOdds,
				Form:     f.MatchDetail.HometeamForm,
			},
			AwayTeam: models.APITeam{
				Nickname: f.Team_2.Nickname,
				Score:    f.MatchDetail.AwayteamScore,
				Odds:     f.MatchDetail.AwayteamOdds,
				Form:     f.MatchDetail.AwayteamForm,
			},
			KickOffTime: f.Fixture.Kickofftime.Time,
		})
	}

	return apiFixtures, nil
}

// GetCompetitionFixtures fetches fixtures for a specific competition and converts them to API models.
func (s *APIDataService) GetRoundCompetitionFixtures(competitionId int64, round int) ([]models.APIFixture, error) {
	fixtures, err := s.queries.ListRoundMatchDetailsByCompetitionID(s.ctx, db.ListRoundMatchDetailsByCompetitionIDParams{
		CompetitionID: competitionId,
		Roundtitle:    fmt.Sprintf("Round %d", round),
	})
	if err != nil {
		return nil, err
	}

	// Convert database models to API models.
	apiFixtures := make([]models.APIFixture, 0)
	for _, f := range fixtures {
		apiFixtures = append(apiFixtures, models.APIFixture{
			ID:            f.Fixture.ID,
			CompetitionID: f.Fixture.CompetitionID,
			RoundTitle:    f.Fixture.Roundtitle,
			MatchState:    f.Fixture.Matchstate,
			Venue:         f.Fixture.Venue,
			VenueCity:     f.Fixture.Venuecity,
			HomeTeam: models.APITeam{
				Nickname: f.Team.Nickname,
				Score:    f.MatchDetail.HometeamScore,
				Odds:     f.MatchDetail.HometeamOdds,
				Form:     f.MatchDetail.HometeamForm,
			},
			AwayTeam: models.APITeam{
				Nickname: f.Team_2.Nickname,
				Score:    f.MatchDetail.AwayteamScore,
				Odds:     f.MatchDetail.AwayteamOdds,
				Form:     f.MatchDetail.AwayteamForm,
			},
			KickOffTime: f.Fixture.Kickofftime.Time,
		})
	}

	return apiFixtures, nil
}

// GetFixtureDetails fetches details for a specific fixture and converts them to API models.
func (s *APIDataService) GetFixtureDetails(fixtureId int64) (*models.APIFixture, error) {
	fixture, err := s.queries.GetMatchDetailsByFixtureID(s.ctx, fixtureId)
	if err != nil {
		return nil, err
	}

	apiFixture := models.APIFixture{
		ID:            fixture.Fixture.ID,
		CompetitionID: fixture.Fixture.CompetitionID,
		RoundTitle:    fixture.Fixture.Roundtitle,
		MatchState:    fixture.Fixture.Matchstate,
		Venue:         fixture.Fixture.Venue,
		VenueCity:     fixture.Fixture.Venuecity,
		HomeTeam: models.APITeam{
			Nickname: fixture.Team.Nickname,
			Score:    fixture.MatchDetail.HometeamScore,
			Odds:     fixture.MatchDetail.HometeamOdds,
			Form:     fixture.MatchDetail.HometeamForm,
		},
		AwayTeam: models.APITeam{
			Nickname: fixture.Team_2.Nickname,
			Score:    fixture.MatchDetail.AwayteamScore,
			Odds:     fixture.MatchDetail.AwayteamOdds,
			Form:     fixture.MatchDetail.AwayteamForm,
		},
		KickOffTime: fixture.Fixture.Kickofftime.Time,
	}

	return &apiFixture, nil
}
