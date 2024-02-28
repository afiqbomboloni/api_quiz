package service

import (
	"github.com/afiqbomboloni/api_quiz/entity"
	"github.com/afiqbomboloni/api_quiz/modules/auth/v1/repository"
	"github.com/afiqbomboloni/api_quiz/request"
	"context"
	"errors"

	"github.com/afiqbomboloni/api_quiz/utils"

	"golang.org/x/crypto/bcrypt"
)


type AuthService interface {
	SaveUser(authRequest request.AuthRequest) (entity.User, error)
	GenerateAccessToken(ctx context.Context, user *entity.User) (string, error)
	AuthValidate(username, password string) (*entity.User, error)
	// BlacklistToken(token string) error
}


type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthRepository(authRepository repository.AuthRepository) *authService {
	return &authService{authRepository}
}


// func(s *authService) BlacklistToken(token string) error {
// 	expiryTime, err := s.Get
// }

func(s *authService) SaveUser(authRequest request.AuthRequest) (entity.User, error) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(authRequest.Password), bcrypt.DefaultCost)

	if(authRequest.Role == "") {
		authRequest.Role = "user"
	}

	auth := entity.User {
		Nama: authRequest.Nama,
		Password:string(hash),
		Email: authRequest.Email,
		Role: authRequest.Role,
	}
	
	newAuth, err := s.authRepository.SaveUser(auth)

	return newAuth, err
}

func(s *authService) GenerateAccessToken(ctx context.Context, user *entity.User) (string, error) {

	token, err := utils.GenerateToken(user.ID)


	return token, err
}

func(s *authService) AuthValidate(email, password string) (*entity.User, error) {
	user, err := s.authRepository.GetEmail(email)

	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.New("invalid Credentials")
	}

	return user, err
}