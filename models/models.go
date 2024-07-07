package models

import "gorm.io/gorm"

type UserDetail struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	FullName    string `json:"full_name"`
	Bio         string `json:"bio"`
	PhoneNumber string `json:"phone_number"`
}

type User struct {
	gorm.Model
	Username       string       `json:"username"`
	Email          string       `json:"email"`
	Password       string       `json:"password"`
	UserDetail     UserDetail   `json:"user_detail"`
	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `json:"organization"`
	Roles          []Role       `json:"roles" gorm:"many2many:user_roles"`
}

type Organization struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `json:"users"`
}

type Role struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `json:"users" gorm:"many2many:user_roles"`
}
