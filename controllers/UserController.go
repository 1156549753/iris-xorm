package controllers

import (
	"free/services"

	"github.com/kataras/iris"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

func (u *UserController) GetSignup() string {
	// name := u.Ctx.Request().FormValue("name")
	// fmt.Println("1:", name)
	// password := r.Get("password")
	// if name

	// u.Ctx.Writef(n)
	return u.Service.Sigunp()
}
