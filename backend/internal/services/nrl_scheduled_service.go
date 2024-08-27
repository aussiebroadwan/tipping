package services

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/aussiebroadwan/tipping/backend/config"
	"github.com/aussiebroadwan/tipping/backend/internal/models"
)

// NRLScheduledService handles the scheduled fetching of data from the NRL API.
type NRLScheduledService struct {
	nrlService       *NRLService
	dataService      *NRLDataService
	competitionIDs   []int64
	scheduleChan     chan models.NRLFixture
	scheduledMatches map[string]struct{}
	mu               sync.Mutex
}

// NewNRLScheduledService creates a new instance of NRLScheduledService.
func NewNRLScheduledService(nrlService *NRLService, dataService *NRLDataService, competitionIDs []int64) *NRLScheduledService {
	return &NRLScheduledService{
		nrlService:       nrlService,
		dataService:      dataService,
		competitionIDs:   competitionIDs,
		scheduleChan:     make(chan models.NRLFixture, 100),
		scheduledMatches: make(map[string]struct{}),
	}
}

// Start starts the daily scheduled fetch of NRL data.
func (s *NRLScheduledService) Start(ctx context.Context) {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	// Start match monitoring goroutine
	go s.MonitorMatches(ctx)

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

			// Schedule match monitoring for "Upcoming" fixtures
			if fixture.MatchState == config.MatchStateUpcoming {
				s.scheduleMatchMonitoring(fixture)
			}
		}

		log.Printf("Fetched and stored %d fixtures for competition %d", len(fixtures), competitionID)
	}

	log.Println("Completed scheduled fetch of NRL data")
}

// scheduleMatchMonitoring schedules a match to be checked 80 minutes after its kickoff time.
func (s *NRLScheduledService) scheduleMatchMonitoring(fixture models.NRLFixture) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.scheduledMatches[fixture.ID]; exists {
		// Match is already scheduled, skip it to prevent duplication.
		log.Printf("Match ID %s is already scheduled for monitoring, skipping", fixture.ID)
		return
	}

	kickoffTime, err := time.Parse(time.RFC3339, fixture.KickOffTime)
	if err != nil {
		log.Printf("Error parsing kickoff time for fixture %s: %v", fixture.ID, err)
		return
	}

	checkTime := kickoffTime.Add(80 * time.Minute)
	delay := time.Until(checkTime)

	log.Printf("Scheduling match monitoring for fixture %s at %v", fixture.ID, checkTime)

	// Mark match as scheduled
	s.scheduledMatches[fixture.ID] = struct{}{}

	// Send the fixture to the scheduler channel after the calculated delay
	go func() {
		time.Sleep(delay)
		s.scheduleChan <- fixture
	}()
}

// MonitorMatches monitors matches based on scheduled times and checks their status.
func (s *NRLScheduledService) MonitorMatches(ctx context.Context) {
	for {
		select {
		case fixture := <-s.scheduleChan:
			s.checkMatchStatus(ctx, fixture)
		case <-ctx.Done():
			log.Println("Match monitoring stopped")
			return
		}
	}
}

// checkMatchStatus checks the status of a match and reschedules if it is not yet "FullTime".
func (s *NRLScheduledService) checkMatchStatus(ctx context.Context, fixture models.NRLFixture) {
	log.Printf("Checking status for match ID %s", fixture.ID)

	// Fetch the latest match details from the NRL API
	updatedFixture, err := s.nrlService.fetchMatchDetail(fixture.MatchCentreURL)
	if err != nil {
		log.Printf("Error fetching match details for fixture %s: %v", fixture.ID, err)
		return
	}

	// Update the match details in the database
	err = s.dataService.StoreFixtureAndDetails(*updatedFixture)
	if err != nil {
		log.Printf("Error updating match details for fixture %s: %v", fixture.ID, err)
		return
	}

	// If the match is still not "FullTime", reschedule another check in 5 minutes
	if updatedFixture.MatchState != config.MatchStateFullTime {
		log.Printf("Match ID %s is not yet FullTime, rescheduling check in 5 minutes", fixture.ID)

		go func() {
			select {
			case <-time.After(5 * time.Minute):
				s.scheduleChan <- *updatedFixture
			case <-ctx.Done():
				log.Printf("Match monitoring for fixture %s stopped", fixture.ID)
			}
		}()
	} else {
		log.Printf("Match ID %s is FullTime", fixture.ID)

		// Remove match from the scheduled set after it is "FullTime"
		s.mu.Lock()
		delete(s.scheduledMatches, fixture.ID)
		s.mu.Unlock()

		// Update the State
		if err = s.dataService.UpdateMatchState(fixture.ID, config.MatchStateFullTime); err != nil {
			log.Printf("Error updating match state for fixture %s: %v", fixture.ID, err)
			return
		}

		// Update the Score and Winner
		if err = s.dataService.UpdateMatchScores(fixture.ID, fixture.HomeTeam.ID, fixture.HomeTeam.Score, fixture.AwayTeam.ID, fixture.AwayTeam.Score); err != nil {
			log.Printf("Error updating match scores for fixture %s: %v", fixture.ID, err)
			return
		}
	}
}
