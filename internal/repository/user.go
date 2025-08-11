package repository

import (
	"github.com/Sahil2k07/gRPC-GO/internal/database"
	"github.com/Sahil2k07/gRPC-GO/internal/enum"
	interfaces "github.com/Sahil2k07/gRPC-GO/internal/interface"
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/util"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
)

type userRepository struct{}

func NewUserRepository() interfaces.UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUser(id uint) (model.User, error) {
	var user model.User

	if err := database.DB.Preload("Profile").First(&user, id).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) ListUsers(req view.ListUsers) ([]model.User, int64, error) {
	q := database.DB.Model(&model.User{}).
		Preload("Profile").
		Joins("JOIN profiles ON profiles.user_id = users.id")

	if req.Email != "" {
		q = q.Where("email ILIKE ?", "%"+req.Email+"%")
	}

	if req.UserName != "" {
		q = q.Where("user_name ILIKE ?", "%"+req.UserName+"%")
	}

	if req.FirstName != "" {
		q = q.Where("profiles.first_name ILIKE ?", "%"+req.FirstName+"%")
	}

	if req.LastName != "" {
		q = q.Where("profiles.first_name ILIKE ?", "%"+req.LastName+"%")
	}

	if req.City != "" {
		q = q.Where("profiles.city ILIKE ?", "%"+req.City+"%")
	}

	if req.State != "" {
		q = q.Where("profiles.state ILIKE ?", "%"+req.State+"%")
	}

	if len(req.Role) > 0 {
		q = q.Where("roles ILIKE ?", "%"+req.Role+"%")
	}

	var count int64
	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	q = util.AddPagination(q, req.Page, req.Sort)

	var models []model.User
	if err := q.Find(&models).Error; err != nil {
		return nil, 0, err
	}

	return models, count, nil
}

func (r *userRepository) UpdateUser(req view.UserView) error {
	user, err := r.GetUser(req.ID)
	if err != nil {
		return err
	}

	user.Roles = enum.RolesToString(req.Roles)
	user.UserName = req.UserName

	user.Profile = model.Profile{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Address:   req.Address,
		City:      req.City,
		State:     req.State,
		ZipCode:   req.ZipCode,
		Country:   req.Country,
	}

	if err := r.SaveUser(user); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(id uint) error {
	if err := database.DB.Delete(&model.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) SaveUser(u model.User) error {
	if err := database.DB.Save(&u).Error; err != nil {
		return err
	}

	return nil
}
