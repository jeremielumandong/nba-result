package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeam_FullName(t *testing.T) {
	tests := []struct {
		name     string
		city     string
		teamName string
		expected string
	}{
		{"Lakers", "Los Angeles", "Lakers", "Los Angeles Lakers"},
		{"Celtics", "Boston", "Celtics", "Boston Celtics"},
		{"Heat", "Miami", "Heat", "Miami Heat"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			team := &Team{
				City: tt.city,
				Name: tt.teamName,
			}
			assert.Equal(t, tt.expected, team.FullName())
		})
	}
}