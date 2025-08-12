package interfaces

import (
	"github.com/Sahil2k07/gRPC-GO/internal/auth"
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
)

type (
	UserRepository interface {
		ListUsers(req view.ListUsers) ([]model.User, int64, error)

		GetUser(id uint) (model.User, error)

		UpdateUser(req view.UserView) error

		DeleteUser(id uint) error

		SaveUser(model.User) error
	}

	UserService interface {
		ListUsers(req view.ListUsers) (view.ListResponse, error)

		GetUser(id string) (view.UserView, error)

		UpdateUser(req view.UserView) error

		DeleteUser(id string) error

		UpdatePassword(u *auth.UserData, req view.ChangePasswordRequest) error
	}
)
