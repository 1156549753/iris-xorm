package models

import "time"

type Admin_users struct {
	Id        int
	Username  string `gorm:"size:60"`
	Password  string `gorm:"size:60"`
	Name      string `gorm:"size:255"`
	Avatar    string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
