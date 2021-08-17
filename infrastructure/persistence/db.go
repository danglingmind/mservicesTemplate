package persistence

import (
	"errors"
	"mservicesTemplate/domain/repository"
)

type Repositories struct {
	User *repository.UserRepository
	// Db *gorm.DB      ----> db instance for your use
}

func NewRepositories(dbUser, dbPass string) (*Repositories, error) {
	// connect to the db and initialize the repos

	return &Repositories{
		User: NewUserRepository( /*db*/ ),
	}, errors.New("-")
}

// other db related operations
