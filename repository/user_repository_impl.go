package repository

import (
	"gorm.io/gorm"
	"tcpserver/models"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (r *UserRepositoryImpl) GetUsers() ([]*models.User, error) {
	var users []*models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) GetUserById(userId uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetFullUserById(userId uint) (*models.User, error) {
	var user models.User
	if err := r.DB.
		Model(&models.User{}).
		Preload("Roles").
		Preload("UserDetail").
		Preload("Organization").
		First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Create(user *models.User) (*models.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Update(user *models.User) (*models.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Delete(userId uint) error {
	if err := r.DB.Delete(&models.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}
