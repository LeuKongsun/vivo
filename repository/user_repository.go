package repository

import (
	"vivo/model"

	"gorm.io/gorm"
)

// UserRepository defines the methods that any user repository must implement.
type UserRepository interface {
    FindByID(id string) (*model.User, error)
    FindByEmail(email string) (*model.User, error)
    Create(user *model.User) (*model.User, error)
}

// userRepository is the struct that implements the UserRepository interface.
type userRepository struct {
    db *gorm.DB
}

// NewUserRepository returns a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}
