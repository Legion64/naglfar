package db

import (
	"github.com/Legion64/naglfar/internal/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConnection struct {
	Conn     *gorm.DB
	database string
}

func New() (*DatabaseConnection, error) {
	db, err := gorm.Open(mysql.Open("root:123456@/users"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &DatabaseConnection{Conn: db, database: "users"}, nil
}

func (dc *DatabaseConnection) Migrate() {
	_ = dc.Conn.AutoMigrate(
		&model.Email{},
		&model.Profile{},
	)
}
