package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aussiebroadwan/tipping/backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetCompetitionsAPI(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/competitions", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerRouter.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var competitions []models.APICompetition
	err = json.Unmarshal(rr.Body.Bytes(), &competitions)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(competitions))
}

func TestGetFixturesAPI(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/fixtures", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerRouter.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fixtures []models.APIFixture
	err = json.Unmarshal(rr.Body.Bytes(), &fixtures)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(fixtures))
}

func TestGetCompetitionFixturesAPI(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/fixtures/111", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerRouter.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fixtures []models.APIFixture
	err = json.Unmarshal(rr.Body.Bytes(), &fixtures)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(fixtures))
}

func TestGetCompetitionRoundFixturesAPI(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/fixtures/111?round=26", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerRouter.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fixtures []models.APIFixture
	err = json.Unmarshal(rr.Body.Bytes(), &fixtures)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(fixtures))
}

func TestGetMatchDetailsAPI(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/fixtures/111/20241112610", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerRouter.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fixture models.APIFixture
	err = json.Unmarshal(rr.Body.Bytes(), &fixture)
	assert.NoError(t, err)

	assert.Equal(t, int64(20241112610), fixture.ID)
	assert.Equal(t, "Cowboys", fixture.HomeTeam.Nickname)
	assert.Equal(t, "Storm", fixture.AwayTeam.Nickname)
}

func TestGetCompetitionFixturesInvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/fixtures/999", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerRouter.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetMatchDetailsInvalidCompetitionID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/fixtures/999/20241112610", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerRouter.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetMatchDetailsInvalidMatchID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/fixtures/111/9999999999", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerRouter.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
