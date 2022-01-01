package routers

import (
	"github.com/afavendanor/twittor.git/bd"
	"net/http"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Se debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet" + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

