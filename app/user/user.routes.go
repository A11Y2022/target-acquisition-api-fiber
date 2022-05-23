package user

import (
    "github.com/gofiber/fiber/v2"
)

func GetUserRoutes(app *fiber.App){
    userApi := app.Group("/user")
    userApi.Post("/create", PostUserLogin )
}
