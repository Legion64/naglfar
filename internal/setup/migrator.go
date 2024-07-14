package setup

import (
	"auth-app/internal/db"
	"auth-app/internal/db/model"
)

func MigrateDatabase(db db.DatabaseConnection) error {
	if err := db.Conn.AutoMigrate(&model.User{}); err != nil {
		return err
	}

	return nil
}
