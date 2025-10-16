package db

import "GoTwitter/m/db/respositories"

type Storage struct{
	UserRepository respositories.UserRepository
}

func NewStorage() *Storage {
	return &Storage{
		UserRepository: &respositories.UserRepositoryImpl{}
	}
}


