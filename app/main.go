package main

import (
	"log"

    "os"
	conf "github.com/A11Y2022/target-acquisition-api-fiber/app/configs"
	user "github.com/A11Y2022/target-acquisition-api-fiber/app/user"
	test "github.com/A11Y2022/target-acquisition-api-fiber/app/testSession"
	trial "github.com/A11Y2022/target-acquisition-api-fiber/app/testData"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)


func initRoutes(app *fiber.App) {
	test.GetTestSessionRoutes(app)
	trial.GetTestDataRoutes(app)
	user.GetUserRoutes(app)
	
}

func main() {
	currDir, _ := os.Getwd()
	envDir := "/configs/.env"
	fullPath := currDir + envDir

	err := godotenv.Load(fullPath)
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
