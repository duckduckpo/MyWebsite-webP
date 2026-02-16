package controllers

import (
	"trip_splitter/database"
	"trip_splitter/models"

	"github.com/gofiber/fiber/v2"
)

// Create a new Trip
func CreateTrip(c *fiber.Ctx) error {
	trip := new(models.Trip)

	// Parse JSON body
	if err := c.BodyParser(trip); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// 🛠️ FIX: Handle Existing Users
	// We check if a member already exists by email.
	// If yes, we use the existing user (so we get their valid ID).
	// If no, we let GORM create a new user.
	var processedMembers []models.User
	
	for _, member := range trip.Members {
		var existingUser models.User
		
		// Check DB for user with this email
		if member.Email != "" {
			err := database.DB.Where("email = ?", member.Email).First(&existingUser).Error
			if err == nil {
				// User Found! Use the existing record (which has the correct ID)
				processedMembers = append(processedMembers, existingUser)
			} else {
				// User Not Found! Append as new (GORM will create them)
				processedMembers = append(processedMembers, member)
			}
		} else {
			// No email provided? Just try to create new
			processedMembers = append(processedMembers, member)
		}
	}
	
	// Replace the members list with our processed list
	trip.Members = processedMembers

	// Save to DB
	if result := database.DB.Create(&trip); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Could not create trip",
			"details": result.Error.Error(),
		})
	}

	return c.Status(201).JSON(trip)
}

// Get Trip Dashboard Data
func GetTripDashboard(c *fiber.Ctx) error {
	id := c.Params("id")
	var trip models.Trip

	// Find trip and preload members
	if result := database.DB.Preload("Members").First(&trip, "id = ?", id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Trip not found"})
	}

	// Calculate Total Spent
	var totalSpent float64
	database.DB.Model(&models.Expense{}).Where("trip_id = ?", id).Select("sum(amount)").Scan(&totalSpent)

	return c.JSON(fiber.Map{
		"trip":        trip,
		"total_spent": totalSpent,
	})
}