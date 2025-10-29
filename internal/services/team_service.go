package services

import (
	"time"

	"github.com/jeremielumandong/nba-result/internal/models"
)

// TeamService handles team-related business logic
type TeamService struct {
	// In a real implementation, this would have a repository/database dependency
}

// NewTeamService creates a new TeamService instance
func NewTeamService() *TeamService {
	return &TeamService{}
}

// GetAllTeams returns all NBA teams
func (s *TeamService) GetAllTeams() ([]models.Team, error) {
	// Mock data - replace with actual database calls
	return []models.Team{
		{
			ID:           1,
			Name:         "Lakers",
			City:         "Los Angeles",
			Abbreviation: "LAL",
			Conference:   "Western",
			Division:     "Pacific",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           2,
			Name:         "Celtics",
			City:         "Boston",
			Abbreviation: "BOS",
			Conference:   "Eastern",
			Division:     "Atlantic",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}, nil
}

// GetTeamByID returns a team by its ID
func (s *TeamService) GetTeamByID(id int) (*models.Team, error) {
	teams, err := s.GetAllTeams()
	if err != nil {
		return nil, err
	}

	for _, team := range teams {
		if team.ID == id {
			return &team, nil
		}
	}

	return nil, models.ErrTeamNotFound
}

// GetTeamsByConference returns teams by conference
func (s *TeamService) GetTeamsByConference(conference string) ([]models.Team, error) {
	teams, err := s.GetAllTeams()
	if err != nil {
		return nil, err
	}

	var conferenceTeams []models.Team
	for _, team := range teams {
		if team.Conference == conference {
			conferenceTeams = append(conferenceTeams, team)
		}
	}

	return conferenceTeams, nil
}