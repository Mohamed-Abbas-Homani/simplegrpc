package models

import (
	"gorm.io/gorm"
	"log"
)

func Seed(db *gorm.DB) {
	// Create sample roles
	adminRole := Role{Name: "Admin"}
	userRole := Role{Name: "User"}

	if err := db.Create(&adminRole).Error; err != nil {
		log.Fatalf("could not seed roles: %v", err)
	}
	if err := db.Create(&userRole).Error; err != nil {
		log.Fatalf("could not seed roles: %v", err)
	}

	// Create sample organization
	organization := Organization{Name: "Tech Corp"}

	if err := db.Create(&organization).Error; err != nil {
		log.Fatalf("could not seed organization: %v", err)
	}

	// Create sample users with details
	user1 := User{
		Username:       "john_doe",
		Email:          "john.doe@example.com",
		Password:       "password123",
		OrganizationID: organization.ID,
		UserDetail: UserDetail{
			FullName:    "John Doe",
			Bio:         "Software Developer",
			PhoneNumber: "123-456-7890",
		},
		Roles: []Role{adminRole},
	}

	user2 := User{
		Username:       "jane_doe",
		Email:          "jane.doe@example.com",
		Password:       "password123",
		OrganizationID: organization.ID,
		UserDetail: UserDetail{
			FullName:    "Jane Doe",
			Bio:         "Project Manager",
			PhoneNumber: "098-765-4321",
		},
		Roles: []Role{userRole},
	}

	if err := db.Create(&user1).Error; err != nil {
		log.Fatalf("could not seed users: %v", err)
	}
	if err := db.Create(&user2).Error; err != nil {
		log.Fatalf("could not seed users: %v", err)
	}
}
