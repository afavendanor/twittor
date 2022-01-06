package bd

import (
	"context"
	"github.com/afavendanor/twittor.git/models"
	"time"
)

/*
	Función para grabar relación en BD
 */
func InsertoRelacion(t models.Relacion) (bool, error)  {
	context, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoConnection.Database("twittor")
	colect := db.Collection("relacion")

	_, err := colect.InsertOne(context, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
