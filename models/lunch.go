package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lunch struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Date primitive.DateTime `json:"date" bson:"date"`
	Heater []primitive.ObjectID `json:"heater" bson:"heater"`
	Participants []primitive.ObjectID `json:"participants" bson:"participants"`
}
