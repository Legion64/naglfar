package setup

import (
	"github.com/Legion64/naglfar/internal/db"
	"github.com/Legion64/naglfar/internal/db/model"
)

func MigrateDatabase(db db.DatabaseConnection) error {
	if err := db.Conn.AutoMigrate(&model.User{}); err != nil {
		return err
	}

	return nil
}
