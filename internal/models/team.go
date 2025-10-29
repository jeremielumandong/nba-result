package models

import "time"

// Team represents an NBA team
type Team struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	City         string    `json:"city"`
	Abbreviation string    `json:"abbreviation"`
	Conference   string    `json:"conference"` // "Eastern" or "Western"
	Division     string    `json:"division"`
	LogoURL      string    `json:"logo_url,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// FullName returns the full team name (City + Name)
func (t *Team) FullName() string {
	return t.City + " " + t.Name
}