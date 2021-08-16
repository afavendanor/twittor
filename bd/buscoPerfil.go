package bd

import (
	"context"
	"log"
	"github.com/afavendanor/twittor.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*	BuscarPerfil se encarga de buscar un perfil en BD */
func BuscarPerfil(ID string) (models.Usuario, error)  {

	context, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("twittor")
	colect := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := colect.FindOne(context, condicion).Decode(&perfil)
	perfil.Password=""
	if err != nil {
		log.Fatal("Registro no encontrado " + err.Error())
		return perfil, err
	}
	return perfil, nil
}