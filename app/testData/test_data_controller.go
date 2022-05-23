package testData

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	conf "github.com/A11Y2022/target-acquisition-api-fiber/app/configs"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func findClickData(c *fiber.Ctx, clickInput map[string]interface{}) {
	testInputCollection := conf.MI.DB.Collection("TestInput")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	trialID := fmt.Sprintf("%v", clickInput["trial_id"])
	objectID, err := primitive.ObjectIDFromHex(trialID)

	fmt.Println(clickInput)
	filter := bson.M{"_id": objectID}
	update := bson.D{
		{"$push", bson.D{{"trial_data", clickInput}}},
	}
	result, err := testInputCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	
}

func PostClickData(c *fiber.Ctx) error {
	ClickInput := new(ClickData)

	if err := c.BodyParser(ClickInput); err != nil {
		log.Println(err, ClickInput)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	var clickInputInterface map[string]interface{}
	clickInputJM, _ := json.Marshal(ClickInput)
	json.Unmarshal(clickInputJM, &clickInputInterface)

	findClickData(c, clickInputInterface)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    nil,
		"success": true,
		"message": "ClickInput inserted successfully",
	})

}

func calculateAccuracy(trialData interface{}) {
	myMap := trialData.(map[string]interface{})
	fmt.Println(myMap)
	// var trialDataInterface map[string]interface{}
	// trialDataJM, _ := json.Marshal(trialData)
	// json.Unmarshal(trialDataJM, &trialDataInterface)

	// fmt.Println(trialDataInterface)
	
	// if obj, ok := trialData.(map[string]interface{}); ok{
	// 	for _, data := range obj {
	// 		if dataObj, ok := data.(map[string]interface{}); ok {
	// 			fmt.Println(dataObj["on_target"])
	// 		}
	// 	}
	// }
	
}

func GetTrialData(c *fiber.Ctx) error {
	testInputCollection := conf.MI.DB.Collection("TestInput")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	
	cursor, err := testInputCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	
	var testConfigs []bson.M
	if err = cursor.All(ctx, &testConfigs); err != nil {
		log.Fatal(err)
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    testConfigs,
		"success": true,
		"message": "Test Configs Query Successfull",
	})
}
