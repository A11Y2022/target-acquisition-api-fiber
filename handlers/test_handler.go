package handlers

import (
	"fmt"

	"github.com/A11Y2022/target-acquisition-api-fiber/models"
	"github.com/gofiber/fiber/v2"
)

//Handler function to get the test config settings and post them
// Accept: text/*, application/json
func CreateTestHandler(ctx *fiber.Ctx) error {
	body := struct {
		SubjectId  string `json:"id"`
		NumTrials  string `json:"trials"`
		Difficulty string `json:"difficulty"`
	}{}

	err := ctx.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		ctx.Status(fiber.StatusBadRequest)
		return err
	}

	config := models.Config{
		SubjectId:  body.SubjectId,
		NumTrials:  body.NumTrials,
		Difficulty: body.Difficulty,
	}
	fmt.Println(config)
	return ctx.Status(fiber.StatusOK).JSON(config)
}
