package seeds

import (
	"go_task_app/src/models"

	"github.com/jinzhu/gorm"
)

func CreateUser(db *gorm.DB, FirstName string, LastName string, Email string) error {
	return db.Create(&models.User{
		FirstName: FirstName,
		LastName:  LastName,
		Email:     Email}).Error
}
