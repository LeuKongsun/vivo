package service

import "vivo/model"

// UserService defines the methods for user-related operations.
type UserService interface {
    SignUp(user model.User) (*model.User, error)
    SignIn(email, password string) (*model.User, error)
    GetUserByID(id string) (*model.User, error)
}
