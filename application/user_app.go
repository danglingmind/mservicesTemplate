package application

import (
	"mservicesTemplate/domain/entity"
	"mservicesTemplate/domain/repository"
)

type UserAppInterface interface {
	GetById(id int) (entity.User, error)
}

// this will implement above interface
type UserApp struct {
	User repository.UserRepository // -----> this is where user API is being linked to the actual db entity
}

func (u *UserApp) GetById(id int) (entity.User, error) {
	return u.User.GetById(id)
}
