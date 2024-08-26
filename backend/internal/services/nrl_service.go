package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aussiebroadwan/tipping/backend/internal/models"
)

// NRLService defines a service that interacts with the NRL API to fetch data.
type NRLService struct {
	baseURL string
	client  *http.Client
}

// NewNRLService creates a new instance of NRLService with default settings.
func NewNRLService(baseURL string) *NRLService {
	return &NRLService{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// FetchFixtures fetches all fixtures for a given competition ID and enriches each fixture
// with additional details such as odds and recent form from their respective matchCentreURLs.
func (s *NRLService) FetchFixtures(competitionID, roundNum, season int) ([]models.NRLFixture, error) {
	url := fmt.Sprintf("%s/draw/data?competition=%d", s.baseURL, competitionID)

	if competitionID == 0 {
		return nil, fmt.Errorf("competition ID is required")
	}

	if roundNum > 0 {
		url = fmt.Sprintf("%s&round=%d", url, roundNum)
	}

	if season > 0 {
		url = fmt.Sprintf("%s&season=%d", url, season)
	}

	// Step 1: Fetch basic fixtures data from the main draw endpoint.
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch fixtures: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code from fixtures endpoint: %d", resp.StatusCode)
	}

	var response struct {
		Fixtures []models.NRLFixture `json:"fixtures"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode fixtures response: %w", err)
	}

	// Step 2: Iterate over each fixture to fetch additional details.
	for i, fixture := range response.Fixtures {
		matchDetail, err := s.fetchMatchDetail(fixture.MatchCentreURL)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch match details for fixture %s: %w", fixture.ID, err)
		}

		// Update the fixture with the additional data fetched.
		response.Fixtures[i].ID = matchDetail.ID
		response.Fixtures[i].KickOffTime = matchDetail.KickOffTime
		response.Fixtures[i].HomeTeam.Odds = matchDetail.HomeTeam.Odds
		response.Fixtures[i].AwayTeam.Odds = matchDetail.AwayTeam.Odds
		response.Fixtures[i].HomeTeam.Score = matchDetail.HomeTeam.Score
		response.Fixtures[i].AwayTeam.Score = matchDetail.AwayTeam.Score
		response.Fixtures[i].HomeTeam.Form = matchDetail.HomeTeam.Form
		response.Fixtures[i].AwayTeam.Form = matchDetail.AwayTeam.Form
	}

	return response.Fixtures, nil
}

// fetchMatchDetail fetches additional match details for a specific fixture using its matchCentreURL.
func (s *NRLService) fetchMatchDetail(matchCentreURL string) (*models.NRLFixture, error) {
	// Full URL for the match details data
	url := fmt.Sprintf("%s%s/data", s.baseURL, matchCentreURL)
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch match details from %s: %w", matchCentreURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code from match details endpoint: %d", resp.StatusCode)
	}

	var matchDetail models.NRLFixture
	err = json.NewDecoder(resp.Body).Decode(&matchDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to decode match details response: %w", err)
	}

	return &matchDetail, nil
}
