package bd

import (
	"context"
	"github.com/afavendanor/twittor.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*
	InsertoTweet graba el tweet en la BD
 */
func InsertoTweet(t models.GraboTweet) (string, bool, error)  {
	context, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("twittor")
	colect := db.Collection("tweet")

	registro := bson.M{
		"userid": t.UserID,
		"mensaje": t.Mensaje,
		"fecha": t.Fecha,
	}

	result, err := colect.InsertOne(context, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil

}

