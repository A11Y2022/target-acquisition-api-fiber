package testData

import (
    "github.com/gofiber/fiber/v2"
)
func GetTestDataRoutes(app *fiber.App) {
    trialApi := app.Group("/trial")
    trialApi.Patch("/data", PostClickData )
    trialApi.Get("/data", GetTrialData )
}