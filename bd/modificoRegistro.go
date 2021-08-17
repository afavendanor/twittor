package bd

import (
	"context"
	"github.com/afavendanor/twittor.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*	ModificarRegistro permite modificarlos datos del perfil del usuario */
func ModificarRegistro(u models.Usuario, ID string) (bool, error) {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("twittor")
	colect := db.Collection("usuarios")

	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellido) > 0 {
		registro["apellido"] = u.Apellido
	}
	registro["fechaNaciiento"] = u.FechaNacimiento
	if len(u.Email) > 0 {
		registro["email"] = u.Email
	}
	if len(u.Password) > 0 {
		registro["password"] = u.Password
	}
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioweb"] = u.SitioWeb
	}

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := colect.UpdateOne(context, filtro, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
