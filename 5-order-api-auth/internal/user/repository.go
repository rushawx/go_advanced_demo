package user

import "5-order-api-auth/pkg/db"

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{Database: database}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) GetByPhone(phone string) (*User, error) {
	user := User{}
	result := repo.Database.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

type SessionRepository struct {
	Database *db.Db
}

func NewSessionRepository(database *db.Db) *SessionRepository {
	return &SessionRepository{Database: database}
}

func (repo *SessionRepository) Create(session *Session) (*Session, error) {
	result := repo.Database.Create(session)
	if result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}

func (repo *SessionRepository) GetBySessionId(sessionId string) (*Session, error) {
	session := Session{}
	result := repo.Database.Where("session_id = ?", sessionId).First(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}
