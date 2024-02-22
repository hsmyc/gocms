package handlers

import (
	"context"
	"hsmyc/gocms/db"
	"hsmyc/gocms/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var singleTypeCollection *mongo.Collection = db.GetCollection(db.DB, "singleType")

func InsertSingleType(singleType models.SingleTypeModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := singleTypeCollection.InsertOne(ctx, singleType)
	if err != nil {
		log.Printf("Error while inserting new singleType into db, Reason: %v\n", err)
		return err
	}

	log.Printf("SingleType inserted successfully with id: %v\n", result.InsertedID)
	return nil
}
