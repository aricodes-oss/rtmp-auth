package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connection *gorm.DB
var Path string

func init() {
	Path = os.Getenv("DB_PATH")
	if Path == "" {
		Path = "data.db"
	}

	Connection, _ = gorm.Open(sqlite.Open(Path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}
