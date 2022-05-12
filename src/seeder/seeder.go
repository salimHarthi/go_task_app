package seeder

import (
	"log"

	"go_task_app/src/seeds"
	"go_task_app/src/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Seeder() {
	dbConn := openConnection()
	defer dbConn.Close()

	for _, seed := range seeds.All() {
		if err := seed.Run(dbConn); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}

func openConnection() *gorm.DB {
	db, err := gorm.Open("mysql", util.GoDotEnvVariable("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}
