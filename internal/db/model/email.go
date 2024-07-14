package model

import "gorm.io/gorm"

type Email struct {
	gorm.Model
	IsPrimary bool
	Email     string
}
