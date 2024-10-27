package repository

import (
    "vivo/model"
)

// FindByID retrieves a user by their ID.
func (repo *userRepository) FindByID(id string) (*model.User, error) {
    var user model.User
    err := repo.db.First(&user, "id = ?", id).Error
    return &user, err
}

// FindByEmail retrieves a user by their email.
func (repo *userRepository) FindByEmail(email string) (*model.User, error) {
    var user model.User
    err := repo.db.First(&user, "email = ?", email).Error
    return &user, err
}

// Create adds a new user to the database.
func (repo *userRepository) Create(user *model.User) (*model.User, error) {
    err := repo.db.Create(user).Error
    return user, err
}
