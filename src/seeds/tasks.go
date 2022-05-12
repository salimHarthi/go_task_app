package seeds

import (
	"go_task_app/src/models"
	"time"

	"github.com/jinzhu/gorm"
)

func CreateTask(db *gorm.DB, Title string, Description string, Due time.Time, UserId uint, Open bool) error {
	return db.Create(&models.Task{
		Title:       Title,
		Description: Description,
		Due:         Due,
		Open:        Open,
		UserId:      UserId}).Error

}
