package user

import (
	"context"
	"log"
	"time"

	conf "github.com/A11Y2022/target-acquisition-api-fiber/app/configs"
	"github.com/gofiber/fiber/v2"
)

func PostUserLogin(c *fiber.Ctx) error {
	
	userInputCollection := conf.MI.DB.Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	UserInput := new(User)

	if err := c.BodyParser(UserInput); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	result, err := userInputCollection.InsertOne(ctx, UserInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "TestInput failed to insert",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "TestInput inserted successfully",
	})
}
