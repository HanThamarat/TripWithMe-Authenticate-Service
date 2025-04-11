package core

type UserRepository interface {
	Save(user User) (*User, error)
}