package interfaces

import "demo-shop-backend/application"

type User struct {
	us application.UserAppInterface
}

func NewUser(u application.UserAppInterface) *User {

}

// Define the user handlers here

// func (u *User) GetUserById()   ------> this will become an endpoint for the user
