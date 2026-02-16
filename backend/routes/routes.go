package routes

import (
	"trip_splitter/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/trips", controllers.CreateTrip)          // POST only
	api.Get("/trips/:id", controllers.GetTripDashboard) // GET only

	// Expense Routes
	api.Post("/trips/:id/expenses", controllers.AddExpense) // POST only
	api.Get("/trips/:id/expenses", controllers.GetExpenses) //

	// Balance Routes
	api.Get("/trips/:id/balances", controllers.GetSimplifiedBalances)
}
