package main

import (
	"gorm.io/gen"

	"rauth/db"
	"rauth/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db.Connection) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	models := []interface{}{
		model.User{},
	}
	g.ApplyBasic(models...)
	db.Connection.AutoMigrate(models...)

	// Generate the code
	g.Execute()
}
