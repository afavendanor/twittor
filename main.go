package main

import (
	"log"

	"github.com/afavendanor/twittor.git/bd"
	"github.com/afavendanor/twittor.git/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
