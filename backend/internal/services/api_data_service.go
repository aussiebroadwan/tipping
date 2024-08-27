package services

import (
	"context"

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
