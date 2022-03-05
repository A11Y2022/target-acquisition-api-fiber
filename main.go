package main

import (
    "github.com/A11Y2022/target-acquisition-api-fiber/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func initRoutes(app *fiber.App){
    testApi := app.Group("/test")
    testApi.Post("/create", handlers.CreateTestHandler)

}

func main() {
    app := fiber.New()
    app.Use(logger.New())
    app.Use(cors.New())

    initRoutes(app)

    app.Listen(":3001")
}
