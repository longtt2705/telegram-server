package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct
type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Phone    string             `json:"phone" bson:"phone"`
	Location string             `json:"location" bson:"location"`
	Status   string             `json:"status" bson:"status"`
	Avatar   string             `json:"avatar" bson:"avatar"`
	WorkAt   []string           `json:"workAt" bson:"workAt"`
	Password string             `json:"password" bson:"password"`
}
