package testSessionHandler

import (
	session "github.com/A11Y2022/target-acquisition-api-fiber/testSession"
    "github.com/gofiber/fiber/v2"
)
func GetTestSessionRoutes(app *fiber.App){
    testApi := app.Group("/test")
    testApi.Post("/create", session.PostTestConfig )
}