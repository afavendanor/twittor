package routers

import (
	"encoding/json"
	"github.com/afavendanor/twittor.git/bd"
	"github.com/afavendanor/twittor.git/models"
	"net/http"
	"time"
)

/*
	graboTweet permite insertar en tweet en bd
 */
func GraboTweet(w http.ResponseWriter, r *http.Request ) {

	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "Datos incorrectos " + err.Error(), http.StatusBadRequest)
		return
	}

	registro := models.GraboTweet{
		UserID: IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha: time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente " + err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

