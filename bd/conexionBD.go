package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
	MongoConnection es el objeto de conexion a la base de datos
*/
var MongoConnection = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@twittor.xqgeg.mongodb.net/test")

/*
	ConectarBD es la funciòn que me permite conectarme a la base de datos de mongoDB
*/
func ConectarBD() *mongo.Client {
	client, error := mongo.Connect(context.TODO(), clientOptions)
	if error != nil {
		log.Fatal(error.Error())
		return client
	}
	error = client.Ping(context.TODO(), nil)
	if error != nil {
		log.Fatal(error.Error())
		return client
	}
	log.Println("Conexiòn exitosa con la BD")
	return client
}

/*
	CheckConnection hace ping a la base de datos
*/
func CheckConnection() int {
	error := MongoConnection.Ping(context.TODO(), nil)
	if error != nil {
		return 0
	}
	return 1
}
