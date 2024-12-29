package auth

import (
	"5-order-api-auth/internal/user"
	"errors"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type AuthService struct {
	UserRepository    *user.UserRepository
	SessionRepository *user.SessionRepository
}

func NewAuthService(userRepository *user.UserRepository, sessionRepository *user.SessionRepository) *AuthService {
	return &AuthService{UserRepository: userRepository, SessionRepository: sessionRepository}
}

func (service *AuthService) Register(phone string) (string, error) {
	existedUser, _ := service.UserRepository.GetByPhone(phone)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	user := &user.User{
		Phone: phone,
	}
	_, err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Phone, nil
}

func (service *AuthService) GetSessionId(phone string) (string, error) {
	_, err := service.UserRepository.GetByPhone(phone)
	if err != nil {
		return "", err
	}
	session := &user.Session{
		SessionId: uuid.New().String(),
		Code:      fmt.Sprintf("%06d", rand.Intn(1000000)),
	}
	_, err = service.SessionRepository.Create(session)
	if err != nil {
		return "", err
	}
	fmt.Println(session.Code)
	return session.SessionId, nil
}

func (service *AuthService) Login(sessionId, code string) (string, error) {
	session, err := service.SessionRepository.GetBySessionId(sessionId)
	if err != nil {
		return "", err
	}
	if session.Code != code {
		return "", errors.New(ErrInvalidCode)
	}
	return session.SessionId, nil
}
