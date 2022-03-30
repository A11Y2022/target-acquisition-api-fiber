package testSession

import (
	"context"
	"log"
	"time"

	conf "github.com/A11Y2022/target-acquisition-api-fiber/config"
	"github.com/gofiber/fiber/v2"
)


func PostTestConfig(c *fiber.Ctx) error {
	testInputCollection := conf.MI.DB.Collection("TestInput")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	TestInput := new(Config)

	if err := c.BodyParser(TestInput); err != nil {
        log.Println(err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "success": false,
            "message": "Failed to parse body",
            "error":   err,
        })
    }

	result, err := testInputCollection.InsertOne(ctx, TestInput)
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
