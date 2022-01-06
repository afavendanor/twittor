package routers

import (
	"encoding/json"
	"github.com/afavendanor/twittor.git/bd"
	"net/http"
	"strconv"
)

/*
	Muestra informacion de los tweets de mis seguidores
*/
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parametro página mayor a cero", http.StatusBadRequest)
		return
	}

	result, correcto := bd.LeoTweetSeguidores(IDUsuario, pagTemp)
	if correcto == false {
		http.Error(w, "Error al leer los twwets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
