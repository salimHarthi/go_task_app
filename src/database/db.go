package database

import (
	"go_task_app/src/models"
	"go_task_app/src/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open(util.GoDotEnvVariable("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{}, models.Task{})
}
