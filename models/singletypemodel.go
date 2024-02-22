package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Field struct {
	Name string `json:"name"`
}

type SingleTypeModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Fields []Field            `json:"fields"`
}
