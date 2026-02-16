package controllers

import (
	"trip_splitter/database"
	"trip_splitter/models"

	"github.com/gofiber/fiber/v2"
)

// AddExpense saves a new expense and its splits to the DB
func AddExpense(c *fiber.Ctx) error {
	expense := new(models.Expense)

	if err := c.BodyParser(expense); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Save Expense (GORM handles the auto-increment ID and splits automatically)
	if result := database.DB.Create(&expense); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Could not save expense",
			"details": result.Error.Error(),
		})
	}

	return c.Status(201).JSON(expense)
}

// GetExpenses fetches all expenses for a specific trip
func GetExpenses(c *fiber.Ctx) error {
	tripID := c.Params("id")
	var expenses []models.Expense

	database.DB.Preload("Splits").Where("trip_id = ?", tripID).Find(&expenses)

	return c.JSON(expenses)
}