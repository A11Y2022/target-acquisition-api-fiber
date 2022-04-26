package main

import (
	"log"

	conf "github.com/A11Y2022/target-acquisition-api-fiber/app/configs"
	test "github.com/A11Y2022/target-acquisition-api-fiber/app/testSession"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)


func initRoutes(app *fiber.App){
    test.GetTestSessionRoutes(app)
}

func main() {
    err := godotenv.Load("./configs/.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    app := fiber.New()
    app.Use(logger.New())
    app.Use(cors.New())
    
    conf.ConnectDB()
    initRoutes(app)

    app.Listen(":3001")
}
