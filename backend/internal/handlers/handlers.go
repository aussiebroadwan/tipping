package handlers

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"

	"github.com/aussiebroadwan/tipping/backend/config"
	"github.com/aussiebroadwan/tipping/backend/internal/services"
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
	// mux.HandleFunc("/api/v1/fixtures/{competition_id}/{match_id}", handlers.GetMatchDetails)
	// mux.HandleFunc("/api/v1/teams", handlers.GetTeams)
	// mux.HandleFunc("/api/v1/teams/{team_id}", handlers.GetTeamByID)

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
// @Param competition_id path int true "Competition ID"
// @Success 200 {array} models.APIFixture
// @Failure 400 "Invalid competition_id"
// @Router /api/v1/fixtures/{competition_id} [get]
func (h *Handlers) GetCompetitionFixtures(w http.ResponseWriter, r *http.Request) {
	competitionId := r.URL.Query().Get("competition_id")
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
	if slices.Contains(competitions, competitionID) {
		http.Error(w, "Invalid competition_id", http.StatusBadRequest)
		return
	}

	fixtures, err := h.dataService.GetCompetitionFixtures(int64(competitionID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fixtures)
}
