package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Last_name string `json:"last_name" bson:"last_name"`
    Nick_name string `json:"nick_name" bson:"nick_name"`
    Last_lunch primitive.DateTime `json:"last_lunch" bson:"last_lunch"`
    Benefits int `json:"benefits" bson:"benefits"`
}