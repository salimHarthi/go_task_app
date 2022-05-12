package models

type User struct {
	Id        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Task      []Task `gorm:"foreignKey:UserId"`
}
