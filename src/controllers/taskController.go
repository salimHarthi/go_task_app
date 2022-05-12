package controllers

import (
	"go_task_app/src/database"
	"go_task_app/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetallTasks(c *fiber.Ctx) error {

	var task models.Task
	database.DB.Find(&task)
	if task.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "no tasks for this user",
		})
	}
	return c.JSON(task)
}

func CreateTasks(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return err
	}
	database.DB.Create(&task)
	return c.JSON(task)
}
func UpdateTasks(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var task models.Task
	task.Id = uint(id)
	if err := c.BodyParser(&task); err != nil {
		return err
	}
	database.DB.Model(&task).Updates(&task)
	return c.JSON(task)
}

func DeleteTasks(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))
	var task models.Task
	task.Id = uint(id)

	database.DB.Delete(&task)
	return c.JSON(fiber.Map{
		"message": "deleted",
	})
}
