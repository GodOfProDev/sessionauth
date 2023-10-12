package session

type Session interface {
	GenerateSession()
	GetSession()
	LogIn()
	LogOut()
}
