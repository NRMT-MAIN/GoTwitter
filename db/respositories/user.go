package respositories

import "database/sql"

type UserRepository interface {
	//Create() error
}

type UserRepositoryImpl struct {
	db *sql.DB
}
