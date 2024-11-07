package repositories

import (
	"time"

	"keizer-auth-api/internal/models"

	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

// NewSessionRepository creates a new instance of SessionRepository
func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

// CreateSession saves a new session in the database
func (r *SessionRepository) CreateSession(session *models.Session) error {
	return r.db.Create(session).Error
}

// GetSession retrieves a session by its session ID
func (r *SessionRepository) GetSession(sessionId string) (*models.Session, error) {
	var session models.Session
	if err := r.db.First(&session, "session_id = ?", sessionId).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

// DeleteSession removes a session from the database
func (r *SessionRepository) DeleteSession(sessionId string) error {
	return r.db.Delete(&models.Session{}, sessionId).Error
}

// UpdateSession updates an existing session
func (r *SessionRepository) UpdateSession(session *models.Session) error {
	return r.db.Save(session).Error
}

// FindValidSession checks if the session is valid (not expired)
func (r *SessionRepository) FindValidSession(token string) (*models.Session, error) {
	var session models.Session
	if err := r.db.First(&session, "token = ? AND expires_at > ?", token, time.Now()).Error; err != nil {
		return nil, err
	}
	return &session, nil
}
