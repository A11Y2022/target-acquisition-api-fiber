package user

type User struct {
	UserId    int    `json:"id" bson:"id"`
	FirstName string `json:"givenName" bson:"firstname"`
	LastName  string `json:"familyName" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
}
