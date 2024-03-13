package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// TODO: work on field specs
type FieldSpecs struct {
	Required bool `json:"is_required"`
	Unique   bool `json:"is_unique"`
}

type Field struct {
	Name  string     `json:"name"`
	Type  string     `json:"type"`
	Specs FieldSpecs `json:"specs"`
}

type SingleTypeModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	ApiID  string             `json:"apiId"`
	Name   string             `json:"name"`
	Fields []Field            `json:"fields"`
}
