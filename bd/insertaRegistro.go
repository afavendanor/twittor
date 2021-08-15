package bd

import (
	"context"
	"github.com/afavendanor/twittor.git/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*
	InsertUSer es ka encargada de insertar los registros en la base de datos
*/
func InsertUSer(t models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("twittor")
	colect := db.Collection("usuarios")

	t.Password, _ = EncriptarPassword(t.Password)

	result, err := colect.InsertOne(ctx, t, nil)
	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil

}
