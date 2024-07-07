package interfaces

import "tcpserver/models"

type UserRepository interface {
	GetUsers() ([]*models.User, error)
	GetUserById(userId uint) (*models.User, error)
	GetFullUserById(userId uint) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(userId uint) error
}
