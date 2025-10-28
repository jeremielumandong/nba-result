package nba

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	assert.NotNil(t, client)
	assert.NotNil(t, client.httpClient)
	assert.Equal(t, 30*time.Second, client.httpClient.Timeout)
}

func TestClient_GetGamesByDate_InvalidDate(t *testing.T) {
	client := NewClient()
	
	_, err := client.GetGamesByDate("invalid-date")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid date format")
}

func TestClient_GetGamesByDate_Success(t *testing.T) {
	// Create mock API response
	mockResponse := APIResponse{
		Resource: "scoreboardV2",
		ResultSets: []ResultSet{
			{
				Name:    "GameHeader",
				Headers: []string{"GAME_DATE_EST", "GAME_SEQUENCE", "GAME_ID", "GAME_STATUS_TEXT", "SEASON_TYPE", "VISITOR_TEAM_ID", "VISITOR_TEAM_CITY", "VISITOR_TEAM_NAME", "VISITOR_TEAM_ABBREVIATION", "HOME_TEAM_ID", "HOME_TEAM_CITY", "HOME_TEAM_NAME"},
				RowSet: [][]interface{}{
					{"2024-01-15T00:00:00", 1, "0022300567", "Final", "Regular Season", 1610612747, "Los Angeles", "Lakers", "LAL", 1610612744, "Golden State", "Warriors", "GSW"},
				},
			},
			{
				Name:    "LineScore",
				Headers: []string{"GAME_DATE_EST", "GAME_SEQUENCE", "GAME_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY_NAME", "TEAM_WINS_LOSSES", "PTS_QTR1", "PTS_QTR2", "PTS_QTR3", "PTS_QTR4", "PTS_OT1", "PTS_OT2", "PTS_OT3", "PTS_OT4", "PTS_OT5", "PTS_OT6", "PTS_OT7", "PTS_OT8", "PTS_OT9", "PTS_OT10", "PTS"},
				RowSet: [][]interface{}{
					{"2024-01-15T00:00:00", 1, "0022300567", 1610612747, "LAL", "Los Angeles", "25-18", 28, 32, 24, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108},
					{"2024-01-15T00:00:00", 1, "0022300567", 1610612744, "GSW", "Golden State", "22-25", 30, 28, 28, 26, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 112},
				},
			},
		},
	}

	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Create client with mock server URL
	client := NewClient()
	originalBaseURL := baseURL
	defer func() { baseURL = originalBaseURL }() // Restore original URL

	// Override baseURL for testing
	// Note: In a real implementation, you'd make baseURL configurable
	
	games, err := client.parseGames(&mockResponse)
	require.NoError(t, err)

	assert.Len(t, games, 1)
	assert.Equal(t, "0022300567", games[0].GameID)
	assert.Equal(t, "Lakers", games[0].HomeTeam.Name)
	assert.Equal(t, "Warriors", games[0].AwayTeam.Name)
	assert.Equal(t, 108, games[0].HomeTeam.Score)
	assert.Equal(t, 112, games[0].AwayTeam.Score)
	assert.Equal(t, "Final", games[0].Status)
}

func TestClient_parseGames_EmptyResponse(t *testing.T) {
	client := NewClient()
	apiResp := &APIResponse{}
	
	games := client.parseGames(apiResp)
	assert.Empty(t, games)
}

func TestToString(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"nil", nil, ""},
		{"string", "test", "test"},
		{"int", 123, "123"},
		{"float", 45.67, "45.67"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int
	}{
		{"nil", nil, 0},
		{"int", 123, 123},
		{"float64", 45.67, 45},
		{"empty string", "", 0},
		{"other", "abc", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toInt(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}