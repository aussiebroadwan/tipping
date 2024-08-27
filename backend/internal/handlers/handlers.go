package handlers

import (
	"encoding/json"
	"net/http"

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
	// mux.HandleFunc("/api/v1/fixtures/", handlers.GetFixtures)
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
