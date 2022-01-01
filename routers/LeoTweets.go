package routers

import (
	"encoding/json"
	"github.com/afavendanor/twittor.git/bd"
	"net/http"
	"strconv"
)

/*
	Funcion para leer tweets de bd
*/
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Se debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Se debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Se debe enviar el parámetro página con un valor mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	respuesta, correcto := bd.LeoTweets(ID, pag)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
