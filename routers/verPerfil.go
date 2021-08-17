package routers

import (
	"encoding/json"
	"github.com/afavendanor/twittor.git/bd"
	"net/http"
)

/*	VerPefil permite extraer los valores del perfil */
func VerPerfil(w http.ResponseWriter, r *http.Request)  {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Se debe enviar el paràmetro ID", http.StatusBadRequest)
	}

	perfil, err := bd.BuscarPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurriò un error al intentar buscar el registro " + err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(perfil)

}
