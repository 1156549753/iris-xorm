package do

import (
	"free/models"
	"log"

	"github.com/go-xorm/xorm"
)

type AdminUser struct {
	engine *xorm.Engine
}

func NewAdminUser(engine *xorm.Engine) *AdminUser {
	return &AdminUser{
		engine: engine,
	}
}

func (u *AdminUser) Find() (*models.Admin_users, bool) {
	data := &models.Admin_users{}
	ok, err := u.engine.Where("id = 1").Desc("id").Get(data)
	if err != nil {
		log.Fatal(err)
	}
	defer u.engine.NewSession().Close()
	return data, ok
}
