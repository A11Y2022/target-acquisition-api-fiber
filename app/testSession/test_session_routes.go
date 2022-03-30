package testSession

import (
    "github.com/gofiber/fiber/v2"
)
func GetTestSessionRoutes(app *fiber.App){
    testApi := app.Group("/test")
    testApi.Post("/create", PostTestConfig )
}