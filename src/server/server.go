package server

import "github.com/gofiber/fiber/v2"

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {
	// Initialize a new app
	app := fiber.New()

	// Register the index route with a simple
	// "OK" response. It should return status
	// code 200
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK aaaa")
	})
	// Return the configured app
	return app
}
