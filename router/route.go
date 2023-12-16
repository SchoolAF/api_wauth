package router

import (
	"api/handler"
	"github.com/gofiber/fiber/v2"
)

// Setup our router
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User Group
	user := api.Group("/user")
	check := api.Group("/check")

	// User Routes
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUserByID)
	check.Get("/", handler.CheckPhoneNumber)
    check.Get("/:phoneNumber", handler.CheckPhoneNumber)
}