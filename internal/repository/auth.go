package repository

import (
	"github.com/Sahil2k07/gRPC-GO/internal/database"
	"github.com/Sahil2k07/gRPC-GO/internal/enum"
	errz "github.com/Sahil2k07/gRPC-GO/internal/error"
	interfaces "github.com/Sahil2k07/gRPC-GO/internal/interface"
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
)

type authRepository struct{}

func NewAuthRepository() interfaces.AuthRepository {
	return &authRepository{}
}

func (r *authRepository) CheckUserExist(email string) (bool, error) {
	var count int64

	err := database.DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return true, err
	}

	if count > 0 {
		return true, errz.NewAlreadyExists("user with email already exists")
	}

	return false, nil
}

func (r *authRepository) GetUser(email string) (model.User, error) {
	var user model.User

	err := database.DB.Model(&model.User{}).Preload("Profile").Where("email = ?", email).Find(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *authRepository) AddUser(req view.SignUpRequest) error {
	roles := enum.RolesToString(req.Roles)

	user := model.User{
		Email:    req.Email,
		Password: req.Password,
		Roles:    roles,
		UserName: req.UserName,
		Profile: model.Profile{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Phone:     req.Phone,
			Address:   req.Address,
			City:      req.City,
			State:     req.State,
			Country:   req.Country,
			ZipCode:   req.ZipCode,
		},
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
