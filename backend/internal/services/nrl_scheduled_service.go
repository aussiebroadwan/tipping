package services

import (
	"context"
	"log"
	"time"
)

// NRLScheduledService handles the scheduled fetching of data from the NRL API.
type NRLScheduledService struct {
	nrlService     *NRLService
	dataService    *NRLDataService
	competitionIDs []int64
}

// NewNRLScheduledService creates a new instance of NRLScheduledService.
func NewNRLScheduledService(nrlService *NRLService, dataService *NRLDataService, competitionIDs []int64) *NRLScheduledService {
	return &NRLScheduledService{
		nrlService:     nrlService,
		dataService:    dataService,
		competitionIDs: competitionIDs,
	}
}

// Start starts the daily scheduled fetch of NRL data.
func (s *NRLScheduledService) Start(ctx context.Context) {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	// Perform an initial fetch on startup.
	s.FetchAndStoreData(ctx)

	// Perform fetches on every tick (daily).
	for {
		select {
		case <-ticker.C:
			s.FetchAndStoreData(ctx)
		case <-ctx.Done():
			log.Println("NRL scheduled service stopped")
			return
		}
	}
}

// FetchAndStoreData fetches data from the NRL API and stores it in the database.
func (s *NRLScheduledService) FetchAndStoreData(ctx context.Context) {
	log.Println("Starting scheduled fetch of NRL data")

	for _, competitionID := range s.competitionIDs {
		// Fetch fixtures for the current season.
		fixtures, err := s.nrlService.FetchFixtures(competitionID, 0, time.Now().Year())
		if err != nil {
			log.Printf("Error fetching fixtures for competition %d: %v", competitionID, err)
			continue
		}

		// Store each fetched fixture and its details.
		for _, fixture := range fixtures {
			err := s.dataService.StoreFixtureAndDetails(fixture)
			if err != nil {
				log.Printf("Error storing fixture ID %s: %v", fixture.ID, err)
				continue
			}
		}
	}

	log.Println("Completed scheduled fetch of NRL data")
}
