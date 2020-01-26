package session

import "github.com/TenaHub/api/entity"

// SessionRepository specifies logged in user session related database operations
type SessionRepository interface {
	Session(sessionID string) (*entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionID string) (*entity.Session, []error)
}