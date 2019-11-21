package services

import (
	"fmt"
	"free/datasource"
	"free/do"
)

type userService struct {
	do *do.AdminUser
}

type UserService interface {
	Sigunp() string
}

func NewUserService() UserService {
	return &userService{
		do: do.NewAdminUser(datasource.InstanceMaster()),
	}
}

func (u userService) Sigunp() string {
	user, ok := u.do.Find()
	fmt.Println(user)
	fmt.Println(ok)
	return "mindsd"
}
