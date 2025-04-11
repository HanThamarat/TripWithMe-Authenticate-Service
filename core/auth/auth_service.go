package core

import (
	"time"

	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/conf"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Authenticate(auth Auth) (AuthResponse, error)
}

type AuthServiceImpl struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) Authenticate(auth Auth) (AuthResponse, error) {
	authenticatedUser, err := s.repo.Authenticate(auth)

	if err != nil {
		return AuthResponse{}, err
	}

	config := conf.GetConfig();

	claims := jwt.MapClaims{
		"id": authenticatedUser.User.ID,
		"email": authenticatedUser.User.Email,
		"first_name": authenticatedUser.User.FirstName,
		"last_name": authenticatedUser.User.LastName,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return AuthResponse{}, err;
	}

	 authData := AuthResponse{
		AuthToken: t,
		User: authenticatedUser.User,
	}

	return authData, nil
}

