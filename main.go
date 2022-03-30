package main

import (
	"log"
	"os"

	conf "github.com/A11Y2022/target-acquisition-api-fiber/config"
    test "github.com/A11Y2022/target-acquisition-api-fiber/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)


func initRoutes(app *fiber.App){
    test.GetTestSessionRoutes(app)
}

func main() {
    if os.Getenv("APP_ENV") != "production" {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
    }
    
    app := fiber.New()
    app.Use(logger.New())
    app.Use(cors.New())
    
    conf.ConnectDB()
    initRoutes(app)

    app.Listen(":3001")
}
