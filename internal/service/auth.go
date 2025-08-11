package service

import (
	"github.com/Sahil2k07/gRPC-GO/internal/authentication"
	errz "github.com/Sahil2k07/gRPC-GO/internal/error"
	interfaces "github.com/Sahil2k07/gRPC-GO/internal/interface"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
)

type authService struct {
	repo interfaces.AuthRepository
}

func NewAuthService(r interfaces.AuthRepository) interfaces.AuthService {
	return &authService{repo: r}
}

func (s *authService) SignUp(req view.SignUpRequest) error {
	exists, err := s.repo.CheckUserExist(req.Email)
	if err != nil {
		return err
	}

	if exists {
		return errz.NewAlreadyExists("user with same email already exists")
	}

	passwordHash, err := authentication.HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = passwordHash

	err = s.repo.AddUser(req)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) SignIn(req view.SignInRequest) (string, error) {
	user, err := s.repo.GetUser(req.Email)
	if err != nil {
		return "", err
	}

	err = authentication.CheckPassword(user.Password, req.Password)
	if err != nil {
		return "", errz.NewUnauthorized("Wrong email or password")
	}

	token, err := authentication.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
