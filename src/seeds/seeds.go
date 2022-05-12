package seeds

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func All() []Seed {
	return []Seed{
		Seed{
			Name: "Createfd",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "fd", "aaa", "jonfxxoo@test.com")
			},
		},
		Seed{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "jon", "foo", "jonfoo@test.com")
			},
		},
		Seed{
			Name: "CreateTask1",
			Run: func(db *gorm.DB) error {
				return CreateTask(db, "task1", "task 1 desc", time.Now().Add(48*time.Hour), 1, true)
			},
		},
		Seed{
			Name: "CreateTask2Closed",
			Run: func(db *gorm.DB) error {
				return CreateTask(db, "task1", "task 1 desc", time.Now().Add(48*time.Hour), 1, false)
			},
		},
	}
}
