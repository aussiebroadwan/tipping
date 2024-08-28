package handlers

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"

	"github.com/aussiebroadwan/tipping/backend/config"
	"github.com/aussiebroadwan/tipping/backend/internal/models"
	"github.com/aussiebroadwan/tipping/backend/internal/services"
	"github.com/aussiebroadwan/tipping/backend/internal/utils"
)

// Handlers struct to hold the data service.
type Handlers struct {
	dataService *services.APIDataService
}

// NewHandlers creates a new Handlers instance.
func NewHandlers(dataService *services.APIDataService) *Handlers {
	return &Handlers{
		dataService: dataService,
	}
}

// RegisterRoutes registers all the routes for the API.
func RegisterRoutes(mux *http.ServeMux, dataService *services.APIDataService) *Handlers {
	handlers := NewHandlers(dataService)

	mux.HandleFunc("/api/v1/competitions", handlers.GetCompetitions)
	mux.HandleFunc("/api/v1/fixtures", handlers.GetFixtures)
	mux.HandleFunc("/api/v1/fixtures/{competition_id}", handlers.GetCompetitionFixtures)
	mux.HandleFunc("/api/v1/fixtures/{competition_id}/{match_id}", handlers.GetMatchDetails)

	return handlers
}

// GetCompetitions retrieves a list of all available competitions.
// @Summary Retrieve a list of all available competitions
// @Description Get all competitions
// @Tags competitions
// @Produce json
// @Success 200 {array} models.APICompetition
// @Router /api/v1/competitions [get]
func (h *Handlers) GetCompetitions(w http.ResponseWriter, r *http.Request) {
	competitions, err := h.dataService.GetCompetitions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(competitions)
}

// GetFixtures retrieves all fixtures.
// @Summary Retrieve a list of all fixtures
// @Description Get all fixtures
// @Tags fixtures
// @Produce json
// @Success 200 {array} models.APIFixture
// @Router /api/v1/fixtures [get]
func (h *Handlers) GetFixtures(w http.ResponseWriter, r *http.Request) {
	fixtures, err := h.dataService.GetFixtures()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fixtures)
}

// GetCompetitionFixtures retrieves fixtures for a specific competition.
// @Summary Retrieve fixtures for a specific competition
// @Description Get fixtures by competition ID
// @Tags fixtures
// @Produce json
// @Param competition_id path int true "Competition ID" example(111)
// @Param round query int false "Round number" example(1)
// @Success 200 {array} models.APIFixture
// @Failure 400 "Invalid competition_id"
// @Router /api/v1/fixtures/{competition_id} [get]
func (h *Handlers) GetCompetitionFixtures(w http.ResponseWriter, r *http.Request) {
	competitionId := r.PathValue("competition_id")
	if competitionId == "" {
		http.Error(w, "Missing competition_id query parameter", http.StatusBadRequest)
		return
	}

	// Convert the competition ID to an integer
	competitionID, err := strconv.Atoi(competitionId)
	if err != nil {
		http.Error(w, "Invalid competition_id query parameter", http.StatusBadRequest)
		return
	}

	// Check if the competition exists
	competitions := []int{config.CompetitionNRL, config.CompetitionNRLW, config.CompetitionStateOfOrigin, config.CompetitionStateOfOriginWomens}
	if !slices.Contains(competitions, competitionID) {
		http.Error(w, "Invalid competition_id", http.StatusBadRequest)
		return
	}

	var fixtures []models.APIFixture

	round := r.URL.Query().Get("round")
	if round != "" {
		// Convert the round to an integer
		roundNum, err := strconv.Atoi(round)
		if err != nil {
			http.Error(w, "Invalid round query parameter", http.StatusBadRequest)
			return
		}

		fixtures, err = h.dataService.GetRoundCompetitionFixtures(int64(competitionID), roundNum)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		fixtures, err = h.dataService.GetCompetitionFixtures(int64(competitionID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fixtures)
}

// GetMatchDetails retrieves details for a specific match in a competition.
//
// @Summary Retrieve match details
// @Description Get detailed information for a specific match within a competition.
// @Tags fixtures
// @Produce json
// @Param competition_id path int true "Competition ID" example(111)
// @Param match_id path int true "Match ID" example(20241610510)
// @Success 200 {object} models.APIFixture
// @Failure 400 "Invalid competition_id or match_id, or Fixture does not belong to the specified competition"
// @Failure 500 "Internal server error"
// @Router /api/v1/fixtures/{competition_id}/{match_id} [get]
func (h *Handlers) GetMatchDetails(w http.ResponseWriter, r *http.Request) {
	// Retrieve query parameters
	competitionId := r.PathValue("competition_id")
	matchId := r.PathValue("match_id")

	// Check for missing parameters
	if competitionId == "" {
		http.Error(w, "Missing competition_id query parameter", http.StatusBadRequest)
		return
	}
	if matchId == "" {
		http.Error(w, "Missing match_id query parameter", http.StatusBadRequest)
		return
	}

	// Convert the competition ID to an integer
	competitionID, err := strconv.Atoi(competitionId)
	if err != nil {
		http.Error(w, "Invalid competition_id query parameter", http.StatusBadRequest)
		return
	}

	// Convert the match ID to an integer
	matchID, err := strconv.Atoi(matchId)
	if err != nil {
		http.Error(w, "Invalid match_id query parameter", http.StatusBadRequest)
		return
	}

	// Validate the competition ID
	validCompetitions := []int{config.CompetitionNRL, config.CompetitionNRLW, config.CompetitionStateOfOrigin, config.CompetitionStateOfOriginWomens}
	if !slices.Contains(validCompetitions, competitionID) {
		http.Error(w, "Invalid competition_id", http.StatusBadRequest)
		return
	}

	// Parse match ID
	_, comp, _, _ := utils.ParseMatchID(matchId)
	if !slices.Contains(validCompetitions, comp) {
		http.Error(w, "Invalid match_id", http.StatusBadRequest)
		return
	}

	// Fetch the fixture details from the database
	fixture, err := h.dataService.GetFixtureDetails(int64(matchID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ensure the fixture belongs to the specified competition
	if fixture.CompetitionID != int64(competitionID) {
		http.Error(w, "Fixture does not belong to the specified competition", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fixture)
}
