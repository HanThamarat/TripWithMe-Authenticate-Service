package core

import "go.mongodb.org/mongo-driver/bson/primitive"

// @Description Auth information
// @Description Contains username and password
type Auth struct {
	// User's unique username
	// @example administrator
	Email string `json:"email"`
	
	// User's password
	// @example 123456
	Password string `json:"password"`
}

type AuthResponse struct {
	AuthToken string 		`json:"authToken"`
	User      UserDTO   	`json:"user"`
}

type UserDTO struct {
	ID       string    `json:"id"`
	FirstName string `json:"firstname"`
	LastName string  `json:"lastname"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type MongoUser struct {
	ID        primitive.ObjectID `bson:"_id"`
	Password  string             `bson:"password"`
	FirstName string             `bson:"firstname"`
	LastName  string             `bson:"lastname"`
	Email     string             `bson:"email"`
}
