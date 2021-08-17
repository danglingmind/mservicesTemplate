package persistence

import (
	"fmt"
	"mservicesTemplate/domain/repository"
)

type UserRepo struct {
	// db *gorm.DB ----> db instance
}

var _ repository.UserRepository = &UserRepo{}

func NewUserRepository( /* db *gorm.DB*/ ) *UserRepo {
	return &UserRepo{}
}

func (u *UserRepo) Say() {
	fmt.Println("hello")
}
