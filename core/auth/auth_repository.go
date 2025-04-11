package core

type AuthRepository interface {
	Authenticate(auth Auth) (AuthResponse, error)
}