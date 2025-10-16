package service

import "GoTwitter/m/db/respositories"

type UserService interface {
	Create() error
}

type UserServiceImpl struct {
	userRepository respositories.UserRepository
}

func NewUserService(_userRepository respositories.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (us *UserServiceImpl) Create() error {
	return nil
}
