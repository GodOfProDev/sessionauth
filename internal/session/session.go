package session

type Session interface {
	GenerateSession(userId string) (string, error)
	GetUserIDBySession(session string) (string, error)
	DeleteSession(session string) error
}
