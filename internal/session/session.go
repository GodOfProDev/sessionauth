package session

type Session interface {
	GenerateSession(userId string) (string, error)
	GetUserBySession(session string) (string, error)
	DeleteSession(session string) error
}
