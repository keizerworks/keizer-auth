package models

import "github.com/google/uuid"

type AuthProviderType string

const (
	AuthProviderEmail AuthProviderType = "email"
	// AuthProviderGoogle    AuthProviderType = "google"
	// AuthProviderGithub    AuthProviderType = "github"
	// AuthProviderMicrosoft AuthProviderType = "microsoft"
)

type ApplicationAuthProvider struct {
	Provider AuthProviderType `gorm:"type:varchar(50);not null" json:"provider"`

	// Specific configuration for each provider
	ClientID     string `gorm:"type:varchar(255)" json:"client_id,omitempty"`
	ClientSecret string `gorm:"type:varchar(255)" json:"client_secret,omitempty"`

	Base

	// Additional provider-specific configurations can be added as needed
	Scopes []string `gorm:"type:text[];serializer:json" json:"scopes,omitempty"`

	Application   Application `gorm:"foreignKey:ApplicationID"`
	ApplicationID uuid.UUID   `gorm:"type:uuid;not null;index" json:"application_id"`
	IsEnabled     bool        `gorm:"default:false" json:"is_enabled"`
}
