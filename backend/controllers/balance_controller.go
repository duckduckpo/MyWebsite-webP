package controllers

import (
	"trip_splitter/database"
	"trip_splitter/models"

	"github.com/gofiber/fiber/v2"
)

func GetSimplifiedBalances(c *fiber.Ctx) error {
	tripID := c.Params("id")
	var expenses []models.Expense

	// 1. Fetch all expenses with their splits
	database.DB.Preload("Splits").Where("trip_id = ?", tripID).Find(&expenses)

	// 2. Calculate Net Balances
	// CHANGE: Map key must be 'uint' to match User.ID
	balances := make(map[uint]float64)

	for _, exp := range expenses {
		// Payer gets positive balance (They paid X)
		balances[exp.PaidBy] += exp.Amount

		// Splitters get negative balance (They consumed Y)
		for _, split := range exp.Splits {
			balances[split.UserID] -= split.Amount
		}
	}

	// 3. Algorithm: Simplify Debts
	var debts []models.Debt

	// NOTE: If you implement the simplification algorithm here, 
	// ensure it handles 'uint' IDs for the users.

	return c.JSON(fiber.Map{
		"raw_balances":     balances,
		"simplified_debts": debts,
	})
}