package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Username string
	Password string
	Name     string
	Surname  string
	Emails   []Email
}
