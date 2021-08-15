package bd

import (
	"context"
	"github.com/afavendanor/twittor.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

/*
	CheckExistUser se encarga de validar si ya existe un usuario
*/
func CheckExistUser(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("twittor")
	colect := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario
	err := colect.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID

}
