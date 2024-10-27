package service

import (
    "vivo/model"
    "vivo/repository"
    "github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
    UserRepository repository.UserRepository
    validate       *validator.Validate
}

func (u *UserServiceImpl) SignUp(user model.User) (*model.User, error) {
    err := u.validate.Struct(user)
    if err != nil {
        return nil, err
    }
    return u.UserRepository.Create(&user)
}

func (u *UserServiceImpl) SignIn(email, password string) (*model.User, error) {
    user, err := u.UserRepository.FindByEmail(email)
    if err != nil || user.Password != password {
        return nil, err
    }
    return user, nil
}

func (u *UserServiceImpl) GetUserByID(id string) (*model.User, error) {
    return u.UserRepository.FindByID(id)
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
    return &UserServiceImpl{UserRepository: userRepository, validate: validate}
}
