package main

import (
	"flag"
	"go_task_app/src/database"
	"go_task_app/src/routes"
	"go_task_app/src/seeder"
	"go_task_app/src/server"
	"log"
	"os"
)

func main() {
	handleArgs()
	database.Connect()
	database.AutoMigrate()
	app := server.Setup()

	routes.Setup(app)
	log.Fatal(app.Listen(":8000"))
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			println("seeding...")
			seeder.Seeder()
			println("Done seeding")
			os.Exit(0)
		}
	}
}
