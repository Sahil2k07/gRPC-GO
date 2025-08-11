package view

import (
	"github.com/Sahil2k07/gRPC-GO/internal/enum"
	"github.com/Sahil2k07/gRPC-GO/internal/model"
)

type (
	SignInRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	SignUpRequest struct {
		Email     string      `json:"email" validate:"required,email"`
		UserName  string      `json:"userName" validate:"required"`
		Password  string      `json:"password" validate:"required"`
		FirstName string      `json:"firstName" validate:"required"`
		LastName  string      `json:"lastName" validate:"required"`
		Phone     string      `json:"phone"`
		Address   string      `json:"address"`
		City      string      `json:"city"`
		State     string      `json:"state"`
		ZipCode   string      `json:"zipCode"`
		Country   string      `json:"country"`
		Roles     []enum.Role `json:"roles" validate:"required,dive"`
	}

	ListUsers struct {
		Email     string     `json:"email"`
		UserName  string     `json:"userName"`
		FirstName string     `json:"firstName"`
		LastName  string     `json:"lastName"`
		City      string     `json:"city"`
		State     string     `json:"stage"`
		Role      enum.Role  `json:"role"`
		Page      PageFilter `json:"page"`
		Sort      SortFilter `json:"sort"`
	}

	UserView struct {
		ID        uint        `json:"id" validate:"required"`
		Email     string      `json:"email" validate:"required,email"`
		UserName  string      `json:"userName" validate:"required"`
		FirstName string      `json:"firstName" validate:"required"`
		LastName  string      `json:"lastName" validate:"required"`
		Phone     string      `json:"phone"`
		Address   string      `json:"address"`
		City      string      `json:"city"`
		State     string      `json:"state"`
		ZipCode   string      `json:"zipCode"`
		Country   string      `json:"country"`
		Roles     []enum.Role `json:"roles" validate:"required,dive"`
	}

	ChangePasswordRequest struct {
		OldPassword string `json:"oldPassword" validate:"required"`
		NewPassword string `json:"newPassword" validate:"required,min=8"`
	}
)

func NewUserResponse(m model.User) UserView {
	return UserView{
		ID:        m.ID,
		Email:     m.Email,
		UserName:  m.UserName,
		Roles:     enum.StringToRoles(m.Roles),
		FirstName: m.Profile.FirstName,
		LastName:  m.Profile.LastName,
		Phone:     m.Profile.Phone,
		Address:   m.Profile.Address,
		City:      m.Profile.City,
		State:     m.Profile.State,
		ZipCode:   m.Profile.ZipCode,
		Country:   m.Profile.ZipCode,
	}
}
