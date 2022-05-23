package testSession

import (
	t "github.com/A11Y2022/target-acquisition-api-fiber/app/testData"
)

type Config struct {
	SubjectId  string        `json:"id" bson:"id"`
	NumTrials  string        `json:"trials" bson:"trials"`
	Difficulty string        `json:"difficulty" bson:"difficulty"`
	TrialData  []t.ClickData `json:"trial_data" bson:"trial_data"`
	Accuracy   float32       `json:"accuracy" bson:"accuracy"`
}
