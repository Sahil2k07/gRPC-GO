package interfaces

import (
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
)

type (
	AuthRepository interface {
		CheckUserExist(email string) (bool, error)

		GetUser(email string) (model.User, error)

		AddUser(req view.SignUpRequest) error
	}

	AuthService interface {
		SignUp(req view.SignUpRequest) error

		SignIn(req view.SignInRequest) (string, error)
	}
)
