package controllers

import (
	"go_task_app/src/database"
	"go_task_app/src/models"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {

	var user models.User
	database.DB.Where("id=?", c.Params("id")).First(&user)
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	return c.JSON(user)
}
