package testSession

type Config struct {
	SubjectId  int    `json:"id" bson:"id"`
	NumTrials  int    `json:"trials" bson:"trials"`
	Difficulty string `json:"difficulty" bson:"difficulty"`
}
