package models

//Model to describe test configuration
type Config struct {
	SubjectId	string 	`json:"id"`
	NumTrials	string 	`json:"trials"`
	Difficulty	string 	`json:"difficulty"`
}
