package testData

type ClickData struct {
	Coordinate      Coordinate `json:"coordinate" bson:"coordinate"`
	Time            Time       `json:"time" bson:"time"`
	ValidCoordinate string     `json:"on_target" bson:"on_target"`
	NodeId          string     `json:"node_id" bson:"node_id"`
	TrialNum        string     `json:"trial_num" bson:"trial_num"`
	TrialID         string     `json:"trial_id" bson:"trial_id"`
}

type Coordinate struct {
	X_Pos string `json:"x" bson:"x"`
	Y_Pos string `json:"y" bson:"y"`
}

type Time struct {
	Minutes      string `json:"minutes" bson:"minutes"`
	Seconds      string `json:"seconds" bson:"seconds"`
	Centiseconds string `json:"centiseconds" bson:"centiseconds"`
}
