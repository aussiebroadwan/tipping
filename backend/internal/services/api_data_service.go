package services

import (
	"context"

	"github.com/aussiebroadwan/tipping/backend/internal/db"
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
