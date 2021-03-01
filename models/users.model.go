package models

import (
	"context"
	"log"
	"time"

	"github.com/TelegramServer/structs"
	"github.com/TelegramServer/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionName = "users"
)

// UserModel type
type UserModel utils.Model

// GetUserModel returns user model
func GetUserModel() *UserModel {
	return &UserModel{utils.GetCollection("users")}
}

// GetUserByPhone return a single user from collection
func (model *UserModel) GetUserByPhone(phone string) *structs.User {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "phone", Value: phone}}
	var result = structs.User{}
	err := model.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}
