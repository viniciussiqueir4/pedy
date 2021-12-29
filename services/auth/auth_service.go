package auth

import (
	"errors"
	"pedy/interfaces/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository repositories.IUserRepository
}

func (s *AuthService) Auth(data AuthDto) (AuthResponse, error) {
	user, err := s.repository.GetByEmail(data.Email)

	if err != nil {
		return AuthResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return AuthResponse{}, errors.New("Invalid credentials")
	}

	token, err := NewJwtService().GenerateToken(user.ID)

	if err != nil {
		return AuthResponse{}, err
	}

	user.Password = ""
	return AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

func NewAuthService(r repositories.IUserRepository) *AuthService {
	return &AuthService{
		repository: r,
	}
}
