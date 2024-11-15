package services

import (
	"encoding/json"
	"fmt"
	"keizer-auth/internal/models"
	"keizer-auth/internal/repositories"
	"keizer-auth/internal/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionService struct {
	redisRepo *repositories.RedisRepository
	userRepo  *repositories.UserRepository
}

func NewSessionService(redisRepo *repositories.RedisRepository, userRepo *repositories.UserRepository) *SessionService {
	return &SessionService{redisRepo: redisRepo, userRepo: userRepo}
}

func (ss *SessionService) CreateSession(user *models.User) (string, error) {
	sessionID := utils.GenerateSessionID()

	userJson, err := json.Marshal(user)
	if err != nil {
		return "", fmt.Errorf("error occured %w", err)
	}

	if err = ss.redisRepo.Set(
		"dashboard-user-session-"+sessionID,
		string(userJson),
		utils.SessionExpiresIn,
	); err != nil {
		return "", fmt.Errorf("error in setting session %w", err)
	}

	return sessionID, nil
}

func (ss *SessionService) GetSession(
	sessionId string,
	user *models.User,
) error {
	userSession, err := ss.redisRepo.Get("dashboard-user-session-" + sessionId)
	if err != nil {
		return fmt.Errorf("no session found")
	}

	if err = json.Unmarshal([]byte(userSession), user); err != nil {
		return fmt.Errorf("error in unmarshalling")
	}

	return nil
}

func (ss *SessionService) UpdateSession(sessionId string) error {
	val, err := ss.redisRepo.Get("dashboard-user-session-" + sessionId)
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("session not found")
		}
		return err
	}
	if err = ss.redisRepo.Set(
		"dashboard-user-session-"+sessionId,
		val,
		utils.SessionExpiresIn,
	); err != nil {
		return fmt.Errorf("error in updating session %w", err)
	}
	return nil
}

func (ss *SessionService) TTL(sessionId string) (time.Duration, error) {
	ttl, err := ss.redisRepo.TTL("dashboard-user-session-" + sessionId)
	return ttl, err
}
